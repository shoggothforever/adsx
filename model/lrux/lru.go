package lrux

import (
	"sync"
	"time"
)

type LRUCache struct {
	Head, Tail *Node
	m          sync.Mutex
	loc        map[int]*Node
	Cap        int
}
type Node struct {
	K          int
	V          int
	TimeStamp  time.Time
	Prev, Next *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Head: new(Node),
		Tail: new(Node),
		m:    sync.Mutex{},
		loc:  make(map[int]*Node),
		Cap:  capacity,
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head

	return cache
}

func (c *LRUCache) Get(key int) int {
	c.m.Lock()
	defer c.m.Unlock()
	if node, ok := c.loc[key]; ok {
		c.PushToHeadLock(c.loc[key], ok)
		return node.V
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	c.m.Lock()
	defer c.m.Unlock()
	if _, ok := c.loc[key]; ok {
		c.loc[key].V = value
		c.PushToHeadLock(c.loc[key], ok)
		return
	}
	c.PushToHeadLock(&Node{K: key, V: value}, false)
}

func (c *LRUCache) PushToHeadLock(node *Node, exist bool) {
	if node == c.Head.Next {
		return
	}
	if exist {
		c.Head.Next, c.Head.Next.Prev, node.Next, node.Next.Prev, node.Prev, node.Prev.Next = node, node, c.Head.Next, node.Prev, c.Head, node.Next
	} else {
		c.Head.Next, c.Head.Next.Prev, node.Next, node.Prev = node, node, c.Head.Next, c.Head
	}
	c.loc[node.K] = node
	if len(c.loc) > c.Cap {
		c.PopBackLock()
	}
}
func (c *LRUCache) PopBackLock() {
	node := c.Tail.Prev
	node.Prev.Next, c.Tail.Prev = c.Tail, node.Prev
	delete(c.loc, node.K)
}
