package bitmap_trie

import (
	"fmt"
	"unsafe"
)

// BitmapTrie 是使用位图实现的 Trie 树,目前仅支持插入纯数字id(string)
type BitmapTrie struct {
	// 联通性位图 [][10]，每个单元是10个bit，表示当前数字与下一层数字的联通性
	connBitmap [][64]uint64
}

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const grow = 10

// 传入最大数字位数的长度，大于0
func NewBirdMap() *BitmapTrie {
	return &BitmapTrie{
		connBitmap: make([][64]uint64, grow),
	}
}

// Insert 在Trie中插入一个ID
func (t *BitmapTrie) Insert(id string) {
	// 遍历ID的每一位
	for i := 0; i < len(id)-1; i++ {
		// 获取当前位的数字
		currDigit := id[i] - '0'
		nextDigit := id[i+1] - '0'
		if len(t.connBitmap) <= i {
			t.connBitmap = append(t.connBitmap, make([][64]uint64, grow)...)
		}
		// 标记联通性位图，表示当前层的数字currDigit可以连接到下一层的nextDigit
		t.connBitmap[i][currDigit] |= 1 << nextDigit
	}
	fmt.Println(unsafe.Sizeof(t.connBitmap[0]))
	fmt.Println(len(t.connBitmap))
}

// Search 验证一个ID是否存在
func (t *BitmapTrie) Search(id string) bool {
	// 遍历ID的每一位
	for i := 0; i < len(id)-1; i++ {
		// 获取当前位的数字
		currDigit := id[i] - '0'
		nextDigit := id[i+1] - '0'
		// 检查联通性位图，判断当前数字是否能连到下一位的数字
		if (t.connBitmap[i][currDigit] & (1 << nextDigit)) == 0 {
			// 当前数字无法连接到下一数字
			return false
		}
	}
	// 遍历完表明全联通
	return true
}
