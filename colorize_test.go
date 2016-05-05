package colorize

import (
	"fmt"
	"testing"

	"github.com/golib/assert"
)

func Test_New(t *testing.T) {
	colorize := New("yellow+bBuih:black+h")

	expected := "\x1b[0;1;5;4;7;93m\x1b[100mColorful line!\x1b[0m"
	assert.Equal(t, expected, colorize.Paint("Colorful line!"))
}

func Test_Paint(t *testing.T) {
	var colorize Colorize
	testCases := map[Color]string{
		ColorRed:     "\x1b[0;31mColorful text!\x1b[0m",
		ColorGreen:   "\x1b[0;32mColorful text!\x1b[0m",
		ColorYellow:  "\x1b[0;33mColorful text!\x1b[0m",
		ColorBlue:    "\x1b[0;34mColorful text!\x1b[0m",
		ColorMagenta: "\x1b[0;35mColorful text!\x1b[0m",
		ColorCyan:    "\x1b[0;36mColorful text!\x1b[0m",
		ColorWhite:   "\x1b[0;37mColorful text!\x1b[0m",
		ColorGray:    "\x1b[0;38mColorful text!\x1b[0m",
		ColorBlack:   "\x1b[0;30mColorful text!\x1b[0m",
	}
	for color, expected := range testCases {
		colorize.SetFgColor(color)

		assert.Equal(t, expected, colorize.Paint("Colorful text!"))
	}
}

func Test_SetPlain(t *testing.T) {
	plain := Colorize{Fg: ColorMagenta}

	plain.SetPlain(true)
	assert.Equal(t, "Plain text", plain.Paint("Plain text"))

	plain.SetPlain(false)
	assert.Equal(t, "\x1b[0;35mColorful text\x1b[0m", plain.Paint("Colorful text"))
}

func Test_SetFgColor(t *testing.T) {
	red := Colorize{Fg: ColorRed}
	text := "Foreground text"

	expected := "\x1b[0;31mForeground text\x1b[0m"
	assert.Equal(t, expected, red.Paint(text))

	red.SetFgColor(ColorBlue)
	assert.NotEqual(t, expected, red.Paint(text))
}

func Test_SetBgColor(t *testing.T) {
	yellow := Colorize{Fg: ColorRed, Bg: ColorYellow}
	text := "Background text"

	expected := "\x1b[0;31m\x1b[43mBackground text\x1b[0m"
	assert.Equal(t, expected, yellow.Paint(text))

	yellow.SetBgColor(ColorBlack)
	assert.NotEqual(t, expected, yellow.Paint(text))
}

func Test_PaintWithMultipleInterface(t *testing.T) {
	blue := Colorize{Fg: ColorBlue}

	expected := "\x1b[0;34mMultiple types: 1 1.24 kg\x1b[0m"
	assert.Equal(t, expected, blue.Paint("Multiple types:", 1, 1.24, "kg"))
}

func Test_PaintWithComplexType(t *testing.T) {
	green := Colorize{Bg: ColorGreen}
	data := struct {
		int
		string
	}{1, "colorize"}

	expected := fmt.Sprintf("\x1b[42mComplex types: %v\x1b[0m", data)
	assert.Equal(t, expected, green.Paint("Complex types:", data))
}

func Test_NormalHightIntensity(t *testing.T) {
	ncolorize := Colorize{Fg: ColorBlack}
	hcolorize := Colorize{Fg: ColorBlack, Prop: Property{Fgi: true}}

	assert.NotEqual(t, ncolorize.Paint("intensity colorize"), hcolorize.Paint("intensity colorize"))
}

func Test_Toggle(t *testing.T) {
	colorize := Colorize{Fg: ColorYellow, Bg: ColorBlack}
	colorize.ToggleFgIntensity()
	colorize.ToggleBgIntensity()
	colorize.ToggleBold()
	colorize.ToggleBlink()
	colorize.ToggleUnderline()
	colorize.ToggleInverse()

	expected := "\x1b[0;1;5;4;7;93m\x1b[100mColorful\x1b[0m"
	assert.Equal(t, expected, colorize.Paint("Colorful"))
}
