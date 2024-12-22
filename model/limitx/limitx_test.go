package limitx

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func Test(t *testing.T) {
	convey.Convey("test tokenBuckets", t, func() {
		l := NewLimiter(100, 100)
		numProcs := 200
		limitCnt := 0
		for i := 0; i < numProcs/2; i++ {
			go func() {
				ok := l.AllowN(time.Now(), 2)
				if !ok {
					limitCnt++
				}

			}()
		}
		time.Sleep(time.Second)
		for i := numProcs / 2; i < numProcs; i++ {
			go func() {
				ok := l.AllowN(time.Now(), 2)
				if !ok {
					limitCnt++
				}

			}()
		}
		time.Sleep(2 * time.Second)
		fmt.Println(limitCnt)
		convey.So(limitCnt, convey.ShouldEqual, 100)
	})
}
