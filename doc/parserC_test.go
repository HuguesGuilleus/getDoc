package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangC_parse(t *testing.T) {
	index := Index{}
	fileName := "aaa.c"
	lines := fileLines{
		&line{
			Str:  "// My function por Say Hello World.",
			Type: TYPE_COMMENT,
		},
		&line{
			Str:  "int hello() {",
			Type: TYPE_FUNCTION,
		},
	}
	elementFunc := Element{
		Name:     "hello",
		LineName: "int hello()",
		Type:     "func",
		FileName: fileName,
		LineNum:  2,
		Comment:  []string{"My function por Say Hello World."},
	}
	langC_parse(&index, lines, fileName)
	assert.Equal(t, elementFunc, *index[0], "Function")
}

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
