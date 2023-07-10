package main

import (
	"fmt"
	"skiplist/model"
)

func main() {
	sl := model.NewSList[float64]()
	fmt.Print(*sl.Head)
	var v float64
	for v = 111; v > 11; v-- {
		sl.Insert(v, "s")
		//sl.Insert(v, "s")
	}

	fmt.Println("跳表的长度为", sl.Length, "  跳表的层数为", sl.Level)
	var i uint = 0
	for ; i < sl.Length; i++ {
		head := sl.Head
		for head.Next(int(i)) != nil {
			fmt.Println()
			t1 := *head.Next(int(i))
			//t2, _ := t1.Back()

			//if t2 != nil {
			//	fmt.Printf("第%d层 %+v\n back:%+v", i, t1, *t2)
			//} else {
			fmt.Printf("第%d层 %+v\n", i, t1)
			//}
			head = head.Next(int(i))
		}
	}

	//sl.Insert(3, "s")
	//fmt.Println(sl)
}
