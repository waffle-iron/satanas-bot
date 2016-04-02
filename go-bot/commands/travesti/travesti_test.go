package travesti

import (
	"github.com/xdigu/go-bot"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func Testtravesti(t *testing.T) {
	Convey("Given a text", t, func() {
		cmd := &bot.PassiveCmd{}

		Convey("When the text does not match travesti", func() {
			cmd.Raw = "My name is go-bot, I am awesome."
			s, err := travesti(cmd)

			So(err, ShouldBeNil)
			So(s, ShouldEqual, "")
		})

		Convey("When the text match travesti", func() {
			cmd.Raw = "eu não votei na travesti!"

			s, err := travesti(cmd)

			So(err, ShouldBeNil)
			So(s, ShouldNotEqual, "")
			So(strings.HasPrefix(s, ":travesti: "), ShouldBeTrue)
		})
	})
}
