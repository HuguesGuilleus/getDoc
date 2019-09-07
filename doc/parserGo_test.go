package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangGo_parse(t *testing.T) {
	eqMult := func(name string, lines fileLines, el Element) {
		t.Run(name, func(t *testing.T) {
			index := Index{}
			FileName := "swag.go"
			el.FileName = FileName
			el.Lang = "go"
			if el.LineNum == 0 {
				el.LineNum = 2
			}
			if el.Comment == nil {
				el.Comment = []string{}
			}
			langGo_parse(&index, lines, FileName)
			if len(index) == 0 {
				t.Error("There are any element.")
				return
			}
			assert.Equal(t, el, *index[0], "")
		})
	}
	eq := func(name string, ln string, el Element) {
		lines := fileLines{
			&line{Str: `// A blob for the do something maby very important.`},
			&line{Str: ln},
		}
		el.Comment = []string{`A blob for the do something maby very important.`}
		eqMult(name, lines, el)
	}
	eq("function", "func fx() {", Element{
		Name:     "fx",
		LineName: "func fx()",
		Type:     "func",
	})
	eq("type", "type yolo interface {", Element{
		Name:     "yolo",
		LineName: "type yolo interface",
		Type:     "type",
	})
	eq("type", "type yolo int", Element{
		Name:     "yolo",
		LineName: "type yolo int",
		Type:     "type",
	})
	eq("varSimple", "var yolo int = 5", Element{
		Name:     "yolo",
		LineName: "var yolo int = 5",
		Type:     "var",
	})
	eq("constSimple", "const yolo int = 5", Element{
		Name:     "yolo",
		LineName: "const yolo int = 5",
		Type:     "const",
	})
	eqMult("varMultLine", fileLines{
		&line{Str: "var ("},
		&line{Str: "	yolo = 5"},
		&line{Str: ")"},
	}, Element{
		Name:     "yolo",
		LineName: "yolo = 5",
		Type:     "var",
	})
	eqMult("ConstMultLine", fileLines{
		&line{Str: "const ("},
		&line{Str: "// It's work!"},
		&line{Str: "	yolo = 5"},
		&line{Str: ")"},
	}, Element{
		Name:     "yolo",
		LineName: "yolo = 5",
		Comment:  []string{"It's work!"},
		Type:     "const",
		LineNum:  3,
	})
}

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
