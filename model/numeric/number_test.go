package numeric

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNumeric(t *testing.T) {
	convey.Convey("testNumeric", t, func() {
		a := "123456789123456789"
		b := "987654321987654321"

		// 测试加法
		fmt.Println("Add:", BigAdd(a, b)) // 1111111111111111110
		convey.So("1111111111111111110", convey.ShouldEqual, BigAdd(a, b))
		//// 测试减法
		fmt.Println("Sub:", BigSub(b, a)) // 864197532864197532
		convey.So("864197532864197532", convey.ShouldEqual, BigSub(b, a))

		// 测试乘法
		fmt.Println("Mul:", BigMul(a, b)) // 121932631356500531347203169112635269
		convey.So("121932631356500531347203169112635269", convey.ShouldEqual, BigMul(a, b))
		// 测试除法
		quotient, remainder := BigDiv(b, a)

		fmt.Println("Div:", quotient, "Remainder:", remainder) // 8 Remainder: 9000000009
		convey.So(quotient, convey.ShouldEqual, "8")
		convey.So(remainder, convey.ShouldEqual, "9000000009")

	})
}
