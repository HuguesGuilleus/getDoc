// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

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
			langGo_type(lines)
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
	testType(t, langGo_type, []testingLine{
		{TYPE_COMMENT, "/// aaa", "aaa"},
		{TYPE_COMMENT, "//", ""},
		{TYPE_CODE, "	a = 4.2", "	a = 4.2"},
		{TYPE_FUNCTION, `func yolo(a, b int) {`, "func yolo(a, b int)"},
		{TYPE_FUNCTION, `func (s *swag) yolo(a, b int) {`, "func (s *swag) yolo(a, b int)"},
		{TYPE_TYPEDEF, `type swag int`, "type swag int"},
		{TYPE_TYPEDEF, `type name interface {`, "type name interface"},
		{TYPE_VAR, `var yolo int`, "var yolo int"},
		{TYPE_VAR, `var yolo int = 5`, "var yolo int = 5"},
		{TYPE_VAR, `var yolo = 5`, "var yolo = 5"},
		{TYPE_CONST, `const yolo int`, "const yolo int"},
		{TYPE_CONST, `const yolo int = 5`, "const yolo int = 5"},
		{TYPE_CONST, `const yolo = 5`, "const yolo = 5"},
		{TYPE_CODE, `var (`, "var ("},
		{TYPE_COMMENT, "/// aaa", "aaa"},
		{TYPE_VAR, `a = 5`, "a = 5"},
		{TYPE_VAR, `a int = 5`, "a int = 5"},
		{TYPE_VAR, `a int`, "a int"},
		{TYPE_CODE, `)`, ")"},
		{TYPE_CODE, "const (", "const ("},
		{TYPE_CONST, `b int = 5`, "b int = 5"},
		{TYPE_CODE, ")", ")"},
	})
}
