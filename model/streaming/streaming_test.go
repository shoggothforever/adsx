package streaming

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMid(t *testing.T) {
	convey.Convey("Test Mid", t, func() {
		mids := MakeMedianFinder[int]()
		nums := []int{1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7}
		for _, num := range nums {
			mids.AddNum(num)
		}
		fmt.Println(float64(mids.GetMidDouble()) / 2.0)
		fmt.Println(float64(mids.GetMidDouble()))
	})
}
