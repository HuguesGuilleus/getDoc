// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangBash_parse(t *testing.T) {
	eq := func(name string, lines fileLines, el Element) {
		t.Run(name, func(t *testing.T) {
			index := Index{}
			FileName := "yolo.bash"
			el.FileName = FileName
			el.Lang = "bash"
			langBash_type(lines)
			langBash_parse(&index, lines, FileName)
			if len(index) == 0 {
				t.Error("There are many element")
				return
			}
			assert.Equal(t, el, *index[0], "")
		})
	}
	eq("functionWithFunction", fileLines{
		&line{Str: `# A fonction for the do something maby important`},
		&line{Str: `function fx() {`},
	}, Element{
		Name:     "fx",
		LineName: "function fx() {",
		Type:     "func",
		LineNum:  2,
		Comment:  []string{`A fonction for the do something maby important`},
	})
	eq("functionParenthesis", fileLines{
		&line{Str: `# A fonction for the do something maby important`},
		&line{Str: `fx() {`},
	}, Element{
		Name:     "fx",
		LineName: "fx() {",
		Type:     "func",
		LineNum:  2,
		Comment:  []string{`A fonction for the do something maby important`},
	})
	eq("functionVar", fileLines{
		&line{Str: `a=yolo`},
	}, Element{
		Name:     "a",
		LineName: "a=yolo",
		Type:     "var",
		LineNum:  1,
		Comment:  []string{},
	})
	eq("functionEnvironment", fileLines{
		&line{Str: `export A=yolo`},
	}, Element{
		Name:     "A",
		LineName: "export A=yolo",
		Type:     "var",
		LineNum:  1,
		Comment:  []string{},
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
