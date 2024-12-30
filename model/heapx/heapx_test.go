package heapx

import (
	"fmt"
	"github.com/shoggothforever/adsx/utils"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHeap(t *testing.T) {
	convey.Convey("TestHeap", t, func() {
		h := NewHeap[string]([]string{"tas", "cel", "row", "abc", "bpf"}, utils.Gt[string])
		fmt.Println(h.content)
		h.Push("adas")
		fmt.Println(h.content)
		h.Push("akl")
		fmt.Println(h.content)
		h.Push("bye")
		fmt.Println(h.content)
		for h.Size() != 0 {
			fmt.Println(h.Top())
			h.Pop()
		}
	})
}
