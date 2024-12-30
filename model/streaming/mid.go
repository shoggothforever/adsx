package streaming

import (
	"github.com/shoggothforever/adsx/config"
	"github.com/shoggothforever/adsx/model/heapx"
	"github.com/shoggothforever/adsx/utils"
)

type MedianFinder[T config.Generic] struct {
	gtHeap *heapx.Heap[T]
	ltHeap *heapx.Heap[T]
	size   int
}

func MakeMedianFinder[T config.Generic]() *MedianFinder[T] {
	return &MedianFinder[T]{
		gtHeap: heapx.NewHeap[T](nil, utils.Gt),
		ltHeap: heapx.NewHeap[T](nil, utils.Lt),
		size:   0,
	}
}
func (m *MedianFinder[T]) AddNum(x T) {
	m.size++
	if m.gtHeap.Empty() || utils.Leq(x, m.gtHeap.Top()) {
		m.gtHeap.Push(x)
		if m.gtHeap.Size() > m.ltHeap.Size()+1 {
			m.ltHeap.Push(m.gtHeap.Top())
			m.gtHeap.Pop()
		}
	} else {
		m.ltHeap.Push(x)
		if m.ltHeap.Size() > m.gtHeap.Size() {
			m.gtHeap.Push(m.ltHeap.Top())
			m.ltHeap.Pop()
		}
	}
}
func (m *MedianFinder[T]) GetMidDouble() T {
	if m.size%2 == 0 {
		return m.gtHeap.Top() + m.ltHeap.Top()
	} else {
		return m.gtHeap.Top() + m.gtHeap.Top()
	}
}
