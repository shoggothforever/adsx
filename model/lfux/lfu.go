package lfux

import (
	"container/list"
	"fmt"
	"github.com/shoggothforever/adsx/config"
	"maps"
)

// CacheEntry represents a single entry in the cache
type CacheEntry[K, T config.Generic] struct {
	key       K
	value     T
	frequency int
	element   *list.Element
}

// LFUCache represents a Least Frequently Used (LFU) cache
type LFUCache[K, T config.Generic] struct {
	capacity    int
	minFreq     int
	entries     map[K]*CacheEntry[K, T]
	freqListMap map[int]*list.List
}

// NewLFUCache creates a new LFU cache with the given capacity
func NewLFUCache[K, T config.Generic](capacity int) *LFUCache[K, T] {
	return &LFUCache[K, T]{
		capacity:    capacity,
		entries:     make(map[K]*CacheEntry[K, T]),
		freqListMap: make(map[int]*list.List),
	}
}

// Get retrieves the value for the given key from the cache
func (cache *LFUCache[K, T]) Get(key K) T {
	if entry, exists := cache.entries[key]; exists {
		cache.incrementFrequency(entry)
		return entry.value
	}
	return T(0)
}

// incrementFrequency increases the frequency of a cache entry
func (cache *LFUCache[K, T]) incrementFrequency(entry *CacheEntry[K, T]) {
	freq := entry.frequency
	cache.freqListMap[freq].Remove(entry.element)

	if cache.freqListMap[freq].Len() == 0 {
		delete(cache.freqListMap, freq)
		if cache.minFreq == freq {
			// 更新 minFreq 为下一个存在的频率
			cache.updateMinFreq()
		}
	}

	entry.frequency++
	if cache.freqListMap[entry.frequency] == nil {
		cache.freqListMap[entry.frequency] = list.New()
	}
	entry.element = cache.freqListMap[entry.frequency].PushFront(entry)
}

// Put adds a key-value pair to the cache
func (cache *LFUCache[K, T]) Put(key K, value T) {
	if cache.capacity == 0 {
		return
	}

	if entry, exists := cache.entries[key]; exists {
		//fmt.Println("exist ", entry)
		entry.value = value
		cache.incrementFrequency(entry)
	} else {
		if len(cache.entries) >= cache.capacity {
			cache.evict()
		}
		newEntry := &CacheEntry[K, T]{key: key, value: value, frequency: 1}
		fmt.Println("push ", newEntry)
		if cache.freqListMap[1] == nil {
			cache.freqListMap[1] = list.New()
		}
		newEntry.element = cache.freqListMap[1].PushFront(newEntry)
		cache.entries[key] = newEntry
		cache.minFreq = 1
	}
}

// updateMinFreq 更新 minFreq 为当前缓存中存在的最小频率
func (cache *LFUCache[K, T]) updateMinFreq() {
	cache.minFreq = 0 // 重置 minFreq
	// 遍历 freqListMap 找到最小的非空频率
	for freq := range maps.Keys(cache.freqListMap) {
		if cache.minFreq == 0 || freq < cache.minFreq {
			cache.minFreq = freq
		}
	}
}

// evict removes the least frequently used cache entry
func (cache *LFUCache[K, T]) evict() {
	if cache.capacity == 0 {
		return
	}
	l := cache.freqListMap[cache.minFreq]
	fmt.Println("evict ", l.Back().Value.(*CacheEntry[K, T]))
	entry := l.Remove(l.Back()).(*CacheEntry[K, T])
	delete(cache.entries, entry.key)

	if l.Len() == 0 {
		delete(cache.freqListMap, cache.minFreq)
		// 更新 minFreq 为下一个存在的频率
		cache.updateMinFreq()
	}
}
