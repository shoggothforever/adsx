package test

import (
	"fmt"
	"github.com/shoggothforever/adsx/model/heapx"
	"github.com/shoggothforever/adsx/utils"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHeap(t *testing.T) {
	convey.Convey("TestHeap", t, func() {
		h := heapx.NewHeap[string]([]string{"tas", "cel", "row", "abc", "bpf"}, utils.Gt[string])
		fmt.Println(h.Content)
		h.Push("adas")
		fmt.Println(h.Content)
		h.Push("akl")
		fmt.Println(h.Content)
		h.Push("bye")
		fmt.Println(h.Content)
		for h.Size != 0 {
			fmt.Println(h.Top())
			h.Pop()
		}
	})
}
