package colorize

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Property_String(t *testing.T) {
	ppt := Property{true, true, true, true, true, true}
	assertion := assert.New(t)

	assertion.Equal("1;5;4;7;", ppt.String())
}
