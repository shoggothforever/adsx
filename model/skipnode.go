package model

import (
	"skiplist/config"
)

type Skip[T config.Generic] struct {
	Next *SkipNode[T]
	Span uint
}

type SkipNode[T config.Generic] struct {
	Key   string
	Value T
	level int
	Skips []Skip[T]
	back  *SkipNode[T]
}

func (s *SkipNode[T]) Next(level int) *SkipNode[T] {
	if level < 0 || level >= s.level {
		return nil
	}
	return s.Skips[level].Next
}
func (s *SkipNode[T]) SetNext(level int, next *SkipNode[T]) {
	if level < 0 || level >= s.level {
		return
	}
	s.Skips[level].Next = next
}
func (s *SkipNode[T]) Span(level int) uint {
	if level < 0 || level >= s.level {
		return 0
	}
	return s.Skips[level].Span
}
func (s *SkipNode[T]) SetSpan(level int, span uint) {
	if level < 0 || level >= s.level {
		return
	}
	s.Skips[level].Span = span
}
func (s *SkipNode[T]) Back() (*SkipNode[T], error) {
	return s.back, nil
}
func (s *SkipNode[T]) GetLevel() int {
	return s.level
}

func (s *SkipNode[T]) SetLevel(level int) {
	if level > 0 && level <= config.KMaxHeight {
		s.level = level
	} else if level <= 0 {
		level = 1
	} else {
		level = config.KMaxHeight
	}

}
func NewSNode[T config.Generic](level int, key string, value T) *SkipNode[T] {
	return &SkipNode[T]{
		Key:   key,
		Value: value,
		level: level,
		Skips: make([]Skip[T], level),
	}
}
