package binary_search

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBs(t *testing.T) {
	convey.Convey("testBs", t, func() {
		nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		fmt.Println("nums:", LowerBound(nums, 5))
		fmt.Println("nums:", UpperBound(nums, 5))
	})
}
