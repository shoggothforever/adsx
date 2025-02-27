package mapx

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHackMapx(t *testing.T) {
	convey.Convey("testHackMapx", t, func() {
		m := make(map[int]int)
		_, hm := mapTypeAndValue(m)

		fmt.Printf("Elements | h.B | Buckets\n\n")

		var prevB uint8
		for i := 0; i < 8; i++ {
			m[i] = i
			if hm.B != prevB {
				fmt.Printf("%8d | %3d | %8d\n", hm.count, hm.B, 1<<hm.B)
				prevB = hm.B
			}
		}
		showSpread(m)
	})
}
