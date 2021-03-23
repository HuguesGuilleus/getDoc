// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"testing"
)

func TestLangJs(t *testing.T) {
	TestParserOne(t, "function", LangJs, Element{
		Name:     "fx",
		LineName: `function fx()`,
		Type:     "func",
		LineNum:  2,
		Comment:  []string{"Comment"},
	}, `// Comment
function fx() {
}`)
	TestParserOne(t, "class", LangJs, Element{
		Name:     "Yolo",
		LineName: `class Yolo`,
		Type:     "class",
		LineNum:  1,
	}, `class Yolo {`)
	TestParserOne(t, "var", LangJs, Element{
		Name:     "swag",
		LineName: `var swag = 1559;`,
		Type:     "var",
		LineNum:  1,
	}, `var swag = 1559;`)
	TestParserOne(t, "const", LangJs, Element{
		Name:     "SWAG",
		LineName: `const SWAG = 1559;`,
		Type:     "const",
		LineNum:  1,
	}, `const SWAG = 1559;`)
}

func TestLangJs_type(t *testing.T) {
	testType(t, langJs_type, []testingLine{
		{TYPE_COMMENT, "/// aaa", "aaa"},
		{TYPE_COMMENT, "//", ""},
		{TYPE_COMMENT, "/**aaa", "aaa"},
		{TYPE_COMMENT, "bbb*/", "bbb"},
		{TYPE_COMMENT, "/*aaa*/", "aaa"},
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
