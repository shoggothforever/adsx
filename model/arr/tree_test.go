package arr

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTreeArr(t *testing.T) {
	convey.Convey("test TreeArr", t, func() {
		nums := []int{1, 2, 3, 4, 5, 6}
		arr := NewTreeArr(nums)
		fmt.Println(arr.SumRange(1, 5))
		fmt.Println(arr.SumRange(1, 4))
		fmt.Println(arr.SumRange(0, 5))
		fmt.Println(arr.SumRange(0, 4))
		//fmt.Println(arr.getSum(5))
	})
}
