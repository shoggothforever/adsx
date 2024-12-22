package test

import (
	"github.com/shoggothforever/adsx/model/lfux"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestLFUCache(t *testing.T) {
	convey.Convey("testLFUCacheWithIntKey", t, func() {

		cache := lfux.NewLFUCache[int, int](3)
		cache.Put(1, 10)
		cache.Put(2, 20)
		cache.Put(3, 30)
		convey.So(cache.Get(1), convey.ShouldEqual, 10)
		cache.Put(4, 40)                                // This should evict key 2
		convey.So(cache.Get(2), convey.ShouldEqual, 0)  // Should return -1, since key 2 was evicted
		convey.So(cache.Get(3), convey.ShouldEqual, 30) // Should return 30
		convey.So(cache.Get(4), convey.ShouldEqual, 40) // Should return 40

		cache.Put(5, 50)                                // This should evict key 1
		convey.So(cache.Get(1), convey.ShouldEqual, 0)  // Should return -1, since key 1 was evicted
		convey.So(cache.Get(5), convey.ShouldEqual, 50) // Should return 50

	})
	convey.Convey("testLFUCacheWithStringKey", t, func() {
		cache := lfux.NewLFUCache[string, int](3)
		cache.Put("1", 10)
		cache.Put("2", 20)
		cache.Put("3", 30)
		convey.So(cache.Get("1"), convey.ShouldEqual, 10)
		cache.Put("4", 40)                                // This should evict key 2
		convey.So(cache.Get("2"), convey.ShouldEqual, 0)  // Should return -1, since key 2 was evicted
		convey.So(cache.Get("3"), convey.ShouldEqual, 30) // Should return 30
		convey.So(cache.Get("4"), convey.ShouldEqual, 40) // Should return 40

		cache.Put("5", 50)                                // This should evict key 1
		convey.So(cache.Get("1"), convey.ShouldEqual, 0)  // Should return -1, since key 1 was evicted
		convey.So(cache.Get("5"), convey.ShouldEqual, 50) // Should return 50

	})
}
