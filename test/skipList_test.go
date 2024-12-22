package test

import (
	"fmt"
	"github.com/shoggothforever/adsx/model/skipList"
	rand2 "math/rand"
	"strconv"
	"testing"
)

func TestInsert(t *testing.T) {
	type stype int
	sl := skipList.NewSList[stype]()
	var v stype
	var s string
	//rand2.Seed(42)
	for v = 100; v > 0; v-- {

		v1 := stype(rand2.Intn(100))
		s1 := strconv.Itoa(int(v1))
		v2 := stype(rand2.Intn(100))
		s2 := strconv.Itoa(int(v2))
		sl.Insert(v1, string(s))
		sl.Insert(v2, string(s))
		if v1%2 == 0 {
			sl.Delete(s1, v1)

		}
		if v2%2 == 0 {
			sl.Delete(s2, v2)
		}
	}
	//sl.Delete("s", (float64)(v-1))
	fmt.Println("跳表的长度为", sl.Length, "  跳表的层数为", sl.Level)
	var i uint = sl.Level
	for ; i >= 0; i-- {
		head := sl.Head
		for head.Next(i) != nil {
			fmt.Println()
			t1 := *head.Next(i)
			fmt.Printf("第%d层 %+v\n", i, t1)
			head = head.Next(i)
		}
	}
}
