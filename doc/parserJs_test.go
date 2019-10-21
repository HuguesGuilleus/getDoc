// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"testing"
)

func TestLangJs_parse(t *testing.T) {
	testingFile = "a.js"
	testParser(t, "function", []string{"function fx() {"}, Element{
		Name:     "fx",
		LineName: "function fx()",
		Type:     "func",
	})
	testParser(t, "class", []string{"class Yolo {"}, Element{
		Name:     "Yolo",
		LineName: "class Yolo",
		Type:     "class",
	})
	testParser(t, "var", []string{`var swag = 1559;`}, Element{
		Name:     "swag",
		LineName: "var swag = 1559;",
		Type:     "var",
	})
	testParser(t, "const", []string{"const SWAG = 1559;"}, Element{
		Name:     "SWAG",
		LineName: "const SWAG = 1559;",
		Type:     "const",
	})
}

func TestLangJs_type(t *testing.T) {
	testType(t, langJs_type, []testingLine{
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
