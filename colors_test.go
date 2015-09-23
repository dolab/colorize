package colorize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Colors_HasColor(t *testing.T) {
	assertion := assert.New(t)

	// for ColorNone
	assertion.False((Color(ColorNone)).HasColor())

	colors := []Color{
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
	for _, color := range colors {
		assertion.True(color.HasColor())
	}
}
