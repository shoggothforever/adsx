package alg

import (
	"fmt"
	"github.com/smartystreets/goconvey/convey"
	"maps"
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
	convey.Convey("TestMyers", t, func() {
		s := make(map[int]int)
		s1 := make(map[int]map[int]int)
		s[0] = 1
		s1[0] = make(map[int]int)
		maps.Copy(s1[0], s)
		s[1] = 2
		s1[1] = make(map[int]int)
		maps.Copy(s1[1], s)
		for key := range maps.Keys(s1) {
			fmt.Printf("s1[%d]:\n", key)
			for innerKey := range maps.Keys(s1[key]) {
				fmt.Printf("s[%d]:\n", innerKey)
			}
		}
		fmt.Println(Myers("abcdebef", "acfdbef"))
	})
}
