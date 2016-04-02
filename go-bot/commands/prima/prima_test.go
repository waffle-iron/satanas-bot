package prima

import (
	"github.com/xdigu/go-bot"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func Testprima(t *testing.T) {
	Convey("Given a text", t, func() {
		cmd := &bot.PassiveCmd{}

		Convey("When the text does not match prima", func() {
			cmd.Raw = "My name is go-bot, I am awesome."
			s, err := prima(cmd)

			So(err, ShouldBeNil)
			So(s, ShouldEqual, "")
		})

		Convey("When the text match prima", func() {
			cmd.Raw = "eu não votei na prima!"

			s, err := prima(cmd)

			So(err, ShouldBeNil)
			So(s, ShouldNotEqual, "")
			So(strings.HasPrefix(s, ":prima: "), ShouldBeTrue)
		})
	})
}