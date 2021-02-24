// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"testing"
)

func TestLangBash_parse(t *testing.T) {
	testingFile = "script.bash"
	testParser(t, "functionWithFunction", []string{"function fx() {"}, Element{
		Name:     "fx",
		LineName: "function fx() {",
		Type:     "func",
	})
	testParser(t, "functionParenthesis", []string{"fx() {"}, Element{
		Name:     "fx",
		LineName: "fx() {",
		Type:     "func",
	})
	testParser(t, "var", []string{"a=yolo"}, Element{
		Name:     "a",
		LineName: "a=yolo",
		Type:     "var",
	})
	testParser(t, "environment", []string{"export A=yolo"}, Element{
		Name:     "A",
		LineName: "export A=yolo",
		Type:     "var",
	})
}

func TestLangBash_type(t *testing.T) {
	testType(t, langBash_type, []testingLine{
		{TYPE_CODE, "#!/bin/bash", "#!/bin/bash"},
		{TYPE_COMMENT, "# aaa", "aaa"},
		{TYPE_COMMENT, "#", ""},
		{TYPE_VAR, "a=yolo", "a=yolo"},
		{TYPE_VAR, "export a=yolo", "export a=yolo"},
		{TYPE_FUNCTION, "yolo1() {", "yolo1() {"},
		{TYPE_FUNCTION, "function yolo2() {", "function yolo2() {"},
		{TYPE_FUNCTION, "function yolo3 {", "function yolo3 {"},
		{TYPE_CODE, `echo "a=3"`, `echo "a=3"`},
	})
}
