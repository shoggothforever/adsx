package heapx

import (
	"github.com/shoggothforever/adsx/config"
	"github.com/shoggothforever/adsx/utils"
)

func cmpDefault[T config.Generic](a T, b T) bool {
	return utils.Gt(a, b)
}

type Heap[T config.Generic] struct {
	Content []T
	Size    int
	Cmp     func(a, b T) bool
}

// NewHeap 默认大顶堆
func NewHeap[T config.Generic](content []T, cmp func(a T, b T) bool) *Heap[T] {
	if cmp == nil {
		cmp = cmpDefault
	}
	if content == nil {
		content = make([]T, 0)
	}
	h := &Heap[T]{
		Size:    len(content),
		Content: content,
		Cmp:     cmp,
	}
	h.BuildHeap()
	return h
}
func (h *Heap[T]) BuildHeap() {
	for i := h.Size/2 - 1; i >= 0; i-- {
		h.HeapModify(i)
	}
}
func (h *Heap[T]) HeapModify(parent int) {
	l := parent*2 + 1
	r := parent*2 + 2
	large := parent
	if l < h.Size && h.Cmp(h.Content[l], h.Content[large]) {
		large = l
	}
	if r < h.Size && h.Cmp(h.Content[r], h.Content[large]) {
		large = r
	}
	if parent != large {
		h.Content[parent], h.Content[large] = h.Content[large], h.Content[parent]
		h.HeapModify(large)
	}
}
func (h *Heap[T]) Push(item T) {
	// 添加新元素到堆尾
	h.Content = append(h.Content, item)
	h.Size++
	// 向上调整
	child := h.Size - 1
	parent := (child - 1) / 2
	for child > 0 && h.Cmp(h.Content[child], h.Content[parent]) {
		h.Content[child], h.Content[parent] = h.Content[parent], h.Content[child]
		child = parent
		parent = (child - 1) / 2
	}
}
func (h *Heap[T]) Pop() {
	h.Content[0], h.Content[h.Size-1] = h.Content[h.Size-1], h.Content[0]
	h.Content = h.Content[:h.Size-1]
	h.Size--
	h.HeapModify(0)
}
func (h *Heap[T]) Top() T {
	// 添加新元素到堆尾
	if h.Size == 0 {
		panic("zero")
	}
	return h.Content[0]
}
