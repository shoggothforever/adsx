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
	level uint
	Skips []Skip[T]
	back  *SkipNode[T]
}

func (s *SkipNode[T]) Next(level uint) *SkipNode[T] {
	if level < 0 || level >= s.level {
		return nil
	}
	return s.Skips[level].Next
}
func (s *SkipNode[T]) SetNext(level uint, next *SkipNode[T]) {
	if level < 0 || level >= s.level {
		return
	}
	s.Skips[level].Next = next
}
func (s *SkipNode[T]) Span(level uint) uint {
	if level < 0 || level >= s.level {
		return 0
	}
	return s.Skips[level].Span
}
func (s *SkipNode[T]) SetSpan(level uint, span uint) {
	if level < 0 || level >= s.level {
		return
	}
	s.Skips[level].Span = span
}
func (s *SkipNode[T]) Back() *SkipNode[T] {
	return s.back
}
func (s *SkipNode[T]) GetLevel() uint {
	return s.level
}

func (s *SkipNode[T]) SetLevel(level uint) {
	if level > 0 && level <= config.KMaxHeight {
		s.level = level
	} else if level <= 0 {
		level = 1
	} else {
		level = config.KMaxHeight
	}

}
func NewSNode[T config.Generic](level uint, key string, value T) *SkipNode[T] {
	return &SkipNode[T]{
		Key:   key,
		Value: value,
		level: level,
		Skips: make([]Skip[T], level),
	}
}
