package test

import (
	"skiplist/model/lrux"
	"testing"
)

func TestLru(t *testing.T) {
	lru := lrux.Constructor(5)
	lru.Put(1, 1)

}