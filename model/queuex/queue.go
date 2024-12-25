package queuex

import "github.com/shoggothforever/adsx/config"

type Queue[T config.Generic] struct {
	content []T
}

func NewQueue[T config.Generic](content []T) *Queue[T] {
	if content == nil {
		return &Queue[T]{content: make([]T, 0, 10)}
	}
	return &Queue[T]{content: content}
}
func (s *Queue[T]) Push(item T) {
	s.content = append([]T{item}, s.content...)
}
func (s *Queue[T]) Pop() T {
	if len(s.content) > 0 {
		top := s.content[0]
		if len(s.content) > 1 {
			s.content = s.content[1:]
		} else {
			s.content = nil
		}
		return top
	}
	panic("queue is empty")
}
func (s *Queue[T]) Front() T {
	if len(s.content) > 0 {
		return s.content[0]
	}
	panic("queue is empty")
}
func (s *Queue[T]) Size() int {
	return len(s.content)
}
