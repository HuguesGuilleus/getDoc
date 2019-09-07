package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangJs_parse(t *testing.T) {
	eq := func(name string, lines fileLines, el Element) {
		t.Run(name, func(t *testing.T) {
			index := Index{}
			FileName := "swag.js"
			el.FileName = FileName
			el.Lang = "js"
			langJs_parse(&index, lines, FileName)
			if len(index) == 0 {
				t.Error("There are many element")
				return
			}
			assert.Equal(t, el, *index[0], "")
		})
	}
	eq("function", fileLines{
		&line{Str: `// A fonction for the do something maby important`},
		&line{Str: `function fx() {`},
	}, Element{
		Name:     "fx",
		LineName: "function fx()",
		Type:     "func",
		LineNum:  2,
		Comment:  []string{`A fonction for the do something maby important`},
	})
	eq("class", fileLines{
		&line{Str: `// A fonction for the do something maby important`},
		&line{Str: `class Yolo {`},
	}, Element{
		Name:     "Yolo",
		LineName: "class Yolo",
		Type:     "class",
		LineNum:  2,
		Comment:  []string{`A fonction for the do something maby important`},
	})
	eq("var", fileLines{
		&line{Str: `var swag = 1559;`},
	}, Element{
		Name:     "swag",
		LineName: "var swag = 1559;",
		Type:     "var",
		LineNum:  1,
		Comment:  []string{},
	})
	eq("const", fileLines{
		&line{Str: `const SWAG = 1559;`},
	}, Element{
		Name:     "SWAG",
		LineName: "const SWAG = 1559;",
		Type:     "const",
		LineNum:  1,
		Comment:  []string{},
	})
}

func TestLangJs_type(t *testing.T) {
	input := fileLines{
		&line{Str: "/// aaa"},
		&line{Str: "//"},
		&line{Str: "	a = 4.2 ;"},
		&line{Str: `function fx() {`},
		&line{Str: `xxxfunction fx() {`},
		&line{Str: `class ClassName {`},
		&line{Str: `const YOLO = "Yolo !!!!!!!!!!!!!!!" ;`},
		&line{Str: `var swag1 = "Swag   !d!bfg !Igh sdg,sfbku qef" ;`},
		&line{Str: `let swag2 = "Swag!!"`},
	}
	langJs_type(input)
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
			Str:  `function fx()`,
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  `xxxfunction fx() {`,
			Type: TYPE_CODE,
		},
		&line{
			Str:  `class ClassName`,
			Type: TYPE_CLASS,
		},
		&line{
			Str:  `const YOLO = "Yolo !!!!!!!!!!!!!!!" ;`,
			Type: TYPE_CONST,
		},
		&line{
			Str:  `var swag1 = "Swag   !d!bfg !Igh sdg,sfbku qef" ;`,
			Type: TYPE_VAR,
		},
		&line{
			Str:  `let swag2 = "Swag!!"`,
			Type: TYPE_VAR,
		},
	}, input, "")
}
