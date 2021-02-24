// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"testing"
)

func TestLangXxx_parse(t *testing.T) {
	testingFile = "prog.go"
	testParser(t, "function", []string{"func fx() {"}, Element{
		Name:     "fx",
		LineName: "func fx()",
		Type:     "func",
	})
	testParser(t, "type", []string{"type yolo interface {"}, Element{
		Name:     "yolo",
		LineName: "type yolo interface",
		Type:     "type",
	})
	testParser(t, "type", []string{"type yolo int"}, Element{
		Name:     "yolo",
		LineName: "type yolo int",
		Type:     "type",
	})
	testParser(t, "varSimple", []string{"var yolo int = 5"}, Element{
		Name:     "yolo",
		LineName: "var yolo int = 5",
		Type:     "var",
	})
	testParser(t, "constSimple", []string{"const yolo int = 5"}, Element{
		Name:     "yolo",
		LineName: "const yolo int = 5",
		Type:     "const",
	})
	testParser(t, "varMultLine", []string{
		"var (",
		"// Comment",
		"	yolo = 5",
		")",
	}, Element{
		Name:     "yolo",
		LineName: "yolo = 5",
		Type:     "var",
		LineNum:  3,
	})
	testParser(t, "varMultLine", []string{
		"const (",
		"// Comment",
		"	yolo = 5",
		")",
	}, Element{
		Name:     "yolo",
		LineName: "yolo = 5",
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
