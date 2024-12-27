package alg

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestString(t *testing.T) {
	convey.Convey("TestReverseString", t, func() {
		s := "你好hello world世界"
		convey.So(ReverseString(s), convey.ShouldEqual, "界世dlrow olleh好你")
	})
	convey.Convey("TestLCS", t, func() {
		s1 := "abcde"
		s2 := "ace"
		convey.So(Lcs(s1, s2), convey.ShouldEqual, 3)
	})
	convey.Convey("TestLIS", t, func() {
		s := "132485"
		convey.So(Lis(s), convey.ShouldEqual, 4)
	})
}
