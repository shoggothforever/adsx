package test

import (
	"github.com/shoggothforever/adsx/utils"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCmp(t *testing.T) {

	convey.Convey("TestCmp", t, func() {
		s1 := "123"
		s2 := "234"
		convey.So(utils.Lt(s1, s2), convey.ShouldBeTrue)
	})
}
