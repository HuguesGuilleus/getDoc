// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

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
			langJs_type(lines)
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

// A line for the test
type TestingLine struct {
	// The type expected
	T int
	// B is the string Before typing
	B string
	// A is the string After typing
	A string
}

// Type test the typing of a group of line
func Type(t *testing.T, fx func(fileLines), lines []TestingLine) {
	input := make(fileLines, len(lines), len(lines))
	for i, l := range lines {
		input[i] = &line{
			Str: l.B,
		}
	}
	fx(input)
	for i, l := range lines {
		if l.T != input[i].Type {
			t.Errorf("Type error (line %d)", i)
			t.Log("   Input line: ", l.B)
			t.Logf("   Type: (expected: %d) %d", l.T, input[i].Type)
		}
		if l.A != input[i].Str {
			t.Errorf("String error (line %d)", i)
			t.Log("   Input line: ", l.B)
			t.Log("   Expected:", l.A)
			t.Log("   Received:", input[i].Str)
		}
	}
}

func TestLangJs_type(t *testing.T) {
	Type(t, langJs_type, []TestingLine{
		{TYPE_COMMENT, "/// aaa", "aaa"},
		{TYPE_COMMENT, "//", ""},
		{TYPE_CODE, "	a = 4.2 ;", "	a = 4.2 ;"},
		{TYPE_FUNCTION, "function fx() {", "function fx()"},
		{TYPE_CODE, `xxxfunction fx() {`, `xxxfunction fx() {`},
		{TYPE_CLASS, `class ClassName {`, `class ClassName`},
		{TYPE_CONST, `const YOLO = "Yolo !!!!!!!!!!!!!!!" ;`,
			`const YOLO = "Yolo !!!!!!!!!!!!!!!" ;`},
		{TYPE_VAR, `var swag1 = "Swag   !d!bfg !Igh sdg,sfbku qef" ;`, `var swag1 = "Swag   !d!bfg !Igh sdg,sfbku qef" ;`},
		{TYPE_VAR, `let swag2 = "Swag!!"`, `let swag2 = "Swag!!"`},
	})
}
