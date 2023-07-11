package main

import (
	"fmt"
	"skiplist/model"
)

func main() {
	type stype int
	sl := model.NewSList[stype]()
	var v stype
	for v = 111; v > 11; v-- {
		sl.Insert(v, "s")
		//sl.Insert(v, "s")
		if v%2 == 0 {
			sl.Delete("s", v)
		}
	}
	//sl.Delete("s", (float64)(v-1))
	fmt.Println("跳表的长度为", sl.Length, "  跳表的层数为", sl.Level)
	var i uint = 0
	for ; i < sl.Length; i++ {
		head := sl.Head
		for head.Next(i) != nil {
			fmt.Println()
			t1 := *head.Next(i)
			fmt.Printf("第%d层 %+v\n", i, t1)
			head = head.Next(i)
		}
	}
}
