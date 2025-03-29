package arr

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSegmentTree(t *testing.T) {
	convey.Convey("Testing SegmentTree", t, func() {
		v1 := []int{1, 8, 3, 4, 7, 1, 6, 2}
		n1 := len(v1)
		t := NewSegMentTree(100, v1)
		convey.So(t.query(1, 3, 1, n1, 1), convey.ShouldEqual, 12)
		convey.So(t.query(3, 8, 1, n1, 1), convey.ShouldEqual, 23)
		t.update(3, 10, 1, n1, 1) // [1, 8, 10, 4, 7, 1, 6, 2]
		convey.So(t.query(1, 3, 1, n1, 1), convey.ShouldEqual, 19)
	})
}
