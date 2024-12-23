package stackx

import "github.com/shoggothforever/adsx/config"

type Stack[T config.Generic] struct {
	content []T
}

func NewStack[T config.Generic](content []T) *Stack[T] {
	if content == nil {
		return &Stack[T]{content: make([]T, 0, 10)}
	}
	return &Stack[T]{content: content}
}
func (s *Stack[T]) push(item T) {
	s.content = append(s.content, item)
}
func (s *Stack[T]) Pop() T {
	if len(s.content) > 0 {
		top := s.content[len(s.content)-1]
		s.content = s.content[:len(s.content)-1]
		return top
	}
	panic("stack is empty")
}
func (s *Stack[T]) Top() T {
	if len(s.content) > 0 {
		return s.content[len(s.content)-1]
	}
	panic("stack is empty")
}
func (s *Stack[T]) Size() int {
	return len(s.content)
}
