package bitmap_trie

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	convey.Convey("TestBirdMap", t, func() {
		// 检测0-40亿范围的ID是否出现
		// 创建一个BitmapTrie
		trie := NewBirdMap()

		// 插入一些ID
		trie.Insert("12345678901")
		trie.Insert("98765432109")
		trie.Insert("11111111111")
		trie.Insert("11111111")
		ls := strings.Repeat("123", 112)
		trie.Insert(ls)

		// 查询是否存在某个ID
		convey.So(trie.Search("12345678901a"), convey.ShouldEqual, false)
		convey.So(trie.Search("98765432109"), convey.ShouldEqual, true)
		convey.So(trie.Search("11111111111"), convey.ShouldEqual, true)
		convey.So(trie.Search("12345678900"), convey.ShouldEqual, false)
		convey.So(trie.Search("22222222222"), convey.ShouldEqual, false)
		convey.So(trie.Search("11111111"), convey.ShouldEqual, true)
		convey.So(trie.Search(ls), convey.ShouldEqual, true)
		fmt.Println(uint64(1 << 63))
	})
}
