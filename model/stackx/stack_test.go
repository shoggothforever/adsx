package stackx

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestStack(t *testing.T) {
	convey.Convey("trap water", t, func() {
		convey.So(6, convey.ShouldEqual, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
		convey.So(9, convey.ShouldEqual, trap([]int{4, 2, 0, 3, 2, 5}))
	})
}
func trap(height []int) (ans int) {
	st := NewStack[int](nil)

	for i := 0; i < len(height); i++ {
		bottom := 0
		for st.Size() > 0 && height[st.Top()] < height[i] {
			ans += (height[st.Top()] - bottom) * (i - st.Top() - 1)
			bottom = height[st.Top()]
			st.Pop()
		}
		if st.Size() > 0 {
			ans += (height[i] - bottom) * (i - st.Top() - 1)
		}

		st.push(i)
	}
	return
}
