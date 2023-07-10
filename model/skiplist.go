package model

import (
	"skiplist/config"
	"skiplist/utils"
)

type SkipList[T config.Generic] struct {
	Level      int
	Length     uint
	Head, Tail *SkipNode[T]
}

func NewSList[T config.Generic]() *SkipList[T] {
	sl := &SkipList[T]{
		Level:  1,
		Length: 0,
		Head:   &SkipNode[T]{},
		Tail:   &SkipNode[T]{},
	}
	level := config.KMaxHeight
	var v T
	sl.Head = NewSNode[T](level, "", v)
	for i := 0; i < level; i++ {
		sl.Head.Skips[i].Next = nil
		sl.Head.Skips[i].Span = 0
	}
	sl.Head.back = nil
	sl.Tail = nil
	return sl
}
func (s *SkipList[T]) Insert(value T, key string) *SkipNode[T] {
	var update [config.KMaxHeight]*SkipNode[T]
	var rank [config.KMaxHeight]uint
	x := s.Head
	for i := s.Level - 1; i >= 0; i-- {
		if i == s.Level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.Next(i) != nil && (utils.Gt(value, x.Next(i).Value) ||
			utils.Eq(x.Next(i).Value, value) && x.Next(i).Key < key) {
			rank[i] += x.Span(i)
			x = x.Next(i)
		}
		if x.Next(i) != nil && utils.Eq(x.Next(i).Value, value) {
			return nil
		}
		update[i] = x
	}
	level := utils.RandomLevel()
	//fmt.Println("randomLevel:", level)
	if level > s.Level {
		for i := s.Level; i < level; i++ {
			rank[i] = 0
			update[i] = s.Head
			update[i].level = level
			update[i].SetSpan(i, s.Length)
		}
		s.Level = level
	}
	x = NewSNode[T](level, key, value)
	for i := 0; i < level; i++ {
		//x.Skips[i].Next = update[i].Next(i)
		x.SetNext(i, update[i].Next(i))
		//update[i].Skips[i].Next = x
		update[i].SetNext(i, x)
		//fmt.Println(i, "   ", update[i].Span(i), "    ", rank[0], "     ", rank[i], "    ", update[i].Span(i)-(rank[0]-rank[i]))
		x.SetSpan(i, update[i].Span(i)-(rank[0]-rank[i]))

		update[i].SetSpan(i, 1+rank[0]-rank[i])
	}
	for i := level; i < s.Level; i++ {
		update[i].SetSpan(i, update[i].Span(i)+1)
	}
	if update[0] == s.Head {
		x.back = nil
	} else {
		x.back = update[0]
	}
	if x.Next(0) != nil {
		x.Skips[0].Next.back = x
	} else {
		s.Tail = x
	}
	s.Length++
	return x
}
