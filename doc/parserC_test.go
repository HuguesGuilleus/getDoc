package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangCType(t *testing.T) {
	input := fileLines{
		&line{Str: "/// aaa"},
		&line{Str: "//"},
		&line{Str: "	a = 4.2 ;"},
		&line{Str: "int yolo(f float) {"},
		&line{Str: "* int yolo(f float)"},
	}
	langC_type(input)
	assert.Equal(t, fileLines{
		&line{
			Str:  "aaa",
			Type: TYPE_COMMENT,
		},
		&line{
			Str:  "",
			Type: TYPE_COMMENT,
		},
		&line{
			Str: "	a = 4.2 ;",
			Type: TYPE_CODE,
		},
		&line{
			Str:  "int yolo(f float)",
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  "* int yolo(f float)",
			Type: TYPE_FUNCTION,
		},
	}, input, "")
}
