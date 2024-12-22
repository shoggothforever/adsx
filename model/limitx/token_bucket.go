package limitx

import (
	"math"
	"sync"
	"time"
)

type Rate float64

const Inf = Rate(math.MaxFloat64)
const InfDuration = time.Duration(math.MaxInt64)

type Limiter struct {
	mu       sync.Mutex
	rate     Rate
	burst    int
	tokens   float64
	lastTime time.Time
}

func NewLimiter(rate float64, burst int) *Limiter {
	return &Limiter{
		rate:  Rate(rate),
		burst: burst,
	}
}
func (l *Limiter) Allow() bool {
	return l.AllowN(time.Now(), 1)
}
func (l *Limiter) AllowN(t time.Time, n int) bool {
	return l.reserveN(t, n, 0).ok
}

type Reservation struct {
	ok        bool
	lim       *Limiter
	tokens    int
	timeToAct time.Time
	// This is the Limit at reservation time, it can change later.
	rate Rate
}

func (l *Limiter) reserveN(t time.Time, n int, reserveTime time.Duration) Reservation {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.rate == Inf {
		return Reservation{
			ok:        true,
			lim:       l,
			tokens:    n,
			timeToAct: t,
		}
	} else if l.rate == 0 {
		var ok bool
		if l.burst >= n {
			l.burst -= n
			ok = true
		}
		return Reservation{
			ok:        ok,
			lim:       l,
			tokens:    l.burst,
			timeToAct: t,
		}
	}
	t, tokens := l.advanced(t)
	tokens -= float64(n)
	var waitDuration time.Duration
	if tokens < 0 {
		if l.rate < 0 {
			waitDuration = InfDuration
		} else {
			waitDuration = time.Duration(-tokens / float64(l.rate) * float64(time.Second))
		}

	}
	ok := n <= l.burst && waitDuration <= reserveTime
	r := Reservation{
		ok:   ok,
		lim:  l,
		rate: l.rate,
	}
	if ok {
		r.tokens = n
		r.timeToAct = t
		l.lastTime = t
		l.tokens = tokens
	}
	return r
}

func (l *Limiter) advanced(t time.Time) (newT time.Time, newTokens float64) {
	last := l.lastTime
	if t.Before(last) {
		last = t
	}
	elapsed := t.Sub(last)
	tokens := l.tokens + elapsed.Seconds()*float64(l.rate)
	if burst := float64(l.burst); burst < tokens {
		tokens = burst
	}
	return t, tokens
}
