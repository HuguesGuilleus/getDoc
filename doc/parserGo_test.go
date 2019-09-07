package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangGo_type(t *testing.T) {
	input := fileLines{
		&line{Str: "/// aaa"},
		&line{Str: "//"},
		&line{Str: "	a = 4.2 ;"},
		&line{Str: `func yolo(a, b int) {`},
		&line{Str: `func (s *swag) yolo(a, b int) {`},
		&line{Str: `type swag int`},
		&line{Str: `type name interface {`},
		&line{Str: `var yolo int`},
		&line{Str: `var yolo int = 5`},
		&line{Str: `var yolo = 5`},
		&line{Str: `const yolo int`},
		&line{Str: `const yolo int = 5`},
		&line{Str: `const yolo = 5`},
		&line{Str: `var (`},
		&line{Str: `a = 5`},
		&line{Str: "/// aaa"},
		&line{Str: `a int = 5`},
		&line{Str: `a int`},
		&line{Str: `)`},
		&line{Str: `const (`},
		&line{Str: `b int = 5`},
		&line{Str: `)`},
	}
	langGo_type(input)
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
			Str:  "func yolo(a, b int)",
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  "func (s *swag) yolo(a, b int)",
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  "type swag int",
			Type: TYPE_TYPEDEF,
		},
		&line{
			Str:  "type name interface",
			Type: TYPE_TYPEDEF,
		},
		&line{
			Str:  "var yolo int",
			Type: TYPE_VAR,
		},
		&line{
			Str:  "var yolo int = 5",
			Type: TYPE_VAR,
		},
		&line{
			Str:  "var yolo = 5",
			Type: TYPE_VAR,
		},
		&line{
			Str:  "const yolo int",
			Type: TYPE_CONST,
		},
		&line{
			Str:  "const yolo int = 5",
			Type: TYPE_CONST,
		},
		&line{
			Str:  "const yolo = 5",
			Type: TYPE_CONST,
		},
		&line{
			Str:  "var (",
			Type: TYPE_CODE,
		},
		&line{
			Str:  "a = 5",
			Type: TYPE_VAR,
		},
		&line{
			Str:  "aaa",
			Type: TYPE_COMMENT,
		},
		&line{
			Str:  "a int = 5",
			Type: TYPE_VAR,
		},
		&line{
			Str:  "a int",
			Type: TYPE_VAR,
		},
		&line{
			Str:  ")",
			Type: TYPE_CODE,
		},
		&line{
			Str:  "const (",
			Type: TYPE_CODE,
		},
		&line{
			Str:  "b int = 5",
			Type: TYPE_CONST,
		},
		&line{
			Str:  ")",
			Type: TYPE_CODE,
		},
	}, input, "")
}
