package colorize

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Colors_HasColor(t *testing.T) {
	it := assert.New(t)

	// for ColorNone
	it.False((Color(ColorNone)).HasColor())

	testColors := []Color{
		ColorRed,
		ColorGreen,
		ColorYellow,
		ColorBlue,
		ColorMagenta,
		ColorCyan,
		ColorWhite,
		ColorGray,
		ColorBlack,
	}
	for _, color := range testColors {
		it.True(color.HasColor())
	}
}
