package skipList

import (
	"math"
	"skiplist/config"
	"skiplist/utils"
	"sync"
)

type SkipList[T config.Generic] struct {
	Level      uint
	Length     uint
	Head, Tail *SkipNode[T]
	m          sync.Mutex
}

func NewSList[T config.Generic]() *SkipList[T] {
	sl := &SkipList[T]{
		Level:  1,
		Length: 0,
		Head:   &SkipNode[T]{},
		Tail:   &SkipNode[T]{},
		m:      sync.Mutex{},
	}
	level := config.KMaxHeight
	var v T
	sl.Head = NewSNode[T](level, "", v)
	for i := (uint)(0); i < level; i++ {
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
	s.m.Lock()
	defer s.m.Unlock()
	x := s.Head
	for i := s.Level - 1; i >= 0 && i < math.MaxUint; i-- {
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
		if x.Next(i) != nil && utils.Eq(x.Next(i).Value, value) && x.Next(i).Key == key {
			return x.Next(i)
		}
		update[i] = x
	}
	level := utils.RandomLevel()
	//fmt.Println("randomLevel:", level)
	if level > s.Level {
		for i := s.Level; i < level; i++ {
			rank[i] = 0
			update[i] = s.Head
			//update[i].level = level
			update[i].SetSpan(i, s.Length)
		}
		s.Level = level
	}
	x = NewSNode[T](level, key, value)
	for i := (uint)(0); i < level; i++ {
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
func (s *SkipList[T]) DeleteNode(x *SkipNode[T], update [config.KMaxHeight]*SkipNode[T]) *SkipNode[T] {
	var i uint = 0
	for ; i < s.Level; i++ {
		if update[i].Next(i) == x {
			update[i].SetSpan(i, update[i].Span(i)+x.Span(i)-1)
			update[i].SetNext(i, x.Next(i))
		} else {
			update[i].SetSpan(i, update[i].Span(i)-1)
		}
	}
	if x.Next(0) != nil {
		x.Skips[0].Next.back = x.back
	} else {
		s.Tail = x.back
	}
	//检测删除节点之后最高层是否为空，如果是就删除这一层，并且继续向下检测
	for s.Level > 1 && s.Head.Next(s.Level-1) == nil {
		s.Level--
	}
	s.Length--
	return x
}
func (s *SkipList[T]) Delete(key string, value T) bool {

	var update [config.KMaxHeight]*SkipNode[T]
	s.m.Lock()
	defer s.m.Unlock()
	y := s.Head
	for i := s.Level - 1; i >= 0 && i < math.MaxUint; i-- {
		for y.Next(i) != nil && (utils.Lq(y.Next(i).Value, value) || utils.Eq(y.Next(i).Value, value) && y.Next(i).Key < key) {
			y = y.Next(i)
		}
		update[i] = y
	}
	//fmt.Println(update)
	y = y.Next(0)
	if y != nil && utils.Eq(value, y.Value) && y.Key == key {
		//fmt.Println("正在删除节点:", y)
		s.DeleteNode(y, update)
		return true
	}
	return false

}
