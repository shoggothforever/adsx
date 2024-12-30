package heapx

import (
	"github.com/shoggothforever/adsx/config"
	"github.com/shoggothforever/adsx/utils"
)

func cmpDefault[T config.Generic](a T, b T) bool {
	return utils.Gt(a, b)
}

type Heap[T config.Generic] struct {
	content []T
	size    int
	cmp     func(a, b T) bool
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
		size:    len(content),
		content: content,
		cmp:     cmp,
	}
	h.BuildHeap()
	return h
}
func (h *Heap[T]) BuildHeap() {
	for i := h.size/2 - 1; i >= 0; i-- {
		h.HeapModify(i)
	}
}
func (h *Heap[T]) HeapModify(parent int) {
	l := parent*2 + 1
	r := parent*2 + 2
	large := parent
	if l < h.size && h.cmp(h.content[l], h.content[large]) {
		large = l
	}
	if r < h.size && h.cmp(h.content[r], h.content[large]) {
		large = r
	}
	if parent != large {
		h.content[parent], h.content[large] = h.content[large], h.content[parent]
		h.HeapModify(large)
	}
}
func (h *Heap[T]) Push(item T) {
	// 添加新元素到堆尾
	h.content = append(h.content, item)
	h.size++
	// 向上调整
	child := h.size - 1
	parent := (child - 1) / 2
	for child > 0 && h.cmp(h.content[child], h.content[parent]) {
		h.content[child], h.content[parent] = h.content[parent], h.content[child]
		child = parent
		parent = (child - 1) / 2
	}
}
func (h *Heap[T]) Pop() {
	h.content[0], h.content[h.size-1] = h.content[h.size-1], h.content[0]
	h.content = h.content[:h.size-1]
	h.size--
	h.HeapModify(0)
}
func (h *Heap[T]) Top() T {
	// 添加新元素到堆尾
	if h.size == 0 {
		panic("zero")
	}
	return h.content[0]
}
func (h *Heap[T]) Empty() bool {
	// 添加新元素到堆尾
	return h.size == 0
}
func (h *Heap[T]) Size() int {
	return h.size
}
