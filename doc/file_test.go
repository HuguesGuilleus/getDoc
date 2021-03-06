// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndexPush(t *testing.T) {
	ind := &Index{}
	el := &Element{
		Name: "yolo",
	}
	ind.push(el)
	assert.Equal(t, 1, len(*ind), "length of the index must be one 1")
	if len(*ind) == 1 {
		assert.Equal(t, el, (*ind)[0], "The first Element")
	}
}

func TestListFile(t *testing.T) {
	ind := Index{
		&Element{FileName: "yolo"},
		&Element{FileName: "swag"},
		&Element{FileName: "swag"},
	}
	assert.Equal(t, []string{"swag", "yolo"}, ind.ListFile(), "")
}

func TestGetExt(t *testing.T) {
	assert.Equal(t, "c", getExt("aaa/a.ed.c"), "Input: aaa/a.ed.c")
	assert.Equal(t, "h", getExt("aaa/aded.h"), "Input: aaa/aded.h")
	assert.Equal(t, "go", getExt("aaa/ad.go"), "Input: aaa/ad.go")
	assert.Equal(t, "js", getExt("aaa/ad.js"), "Input: aaa/ad.js")
	assert.Equal(t, "Makefile", getExt("Makefile"), "Input: Makefile")
	assert.Equal(t, "Makefile", getExt("./Makefile"), "Input: ./Makefile")
	assert.Equal(t, "", getExt(""), "Input: <nothing>")
}

func TestLangKnown(t *testing.T) {
	parserListSave := parserList
	defer func() { parserList = parserListSave }()
	langA := &parserFuncs{}
	parserList = map[string]*parserFuncs{
		"a": langA,
		"b": &parserFuncs{},
	}
	if langKnown("zzz/xxx.a") != langA {
		t.Error("A known lang")
	}
	if langKnown("xxx/yyy.zz") != nil {
		t.Error("langKnown with unknwolang must return nil", langKnown("zz"))
	}
}

func TestSplitFile(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		assert.Equal(t, fileLines{
			&line{Str: "// File for test"},
			&line{Str: "aaa"},
			&line{},
		}, splitFile("../dataTest/split.c"), "")
	})
	t.Run("NoFile", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("The function must panic a err")
			}
		}()
		splitFile("dataTest/_")
		t.Error("The function must panic")
	})
}

func TestGetComment(t *testing.T) {
	t.Run("Get Before", func(t *testing.T) {
		list := fileLines{
			&line{Str: "xxx"},
			&line{TYPE_COMMENT, "aaa"},
			&line{},
		}
		assert.Equal(t, []string{"aaa"}, list.getComment(2))
	})
	t.Run("NoComment", func(t *testing.T) {
		list := fileLines{
			&line{Str: "xxx"},
		}
		assert.Equal(t, []string{}, list.getComment(0))
	})
	t.Run("No Comment", func(t *testing.T) {
		list := fileLines{
			&line{TYPE_CODE, "xxx"},
			&line{},
		}
		assert.Equal(t, []string{}, list.getComment(1))
	})
	t.Run("MultLine", func(t *testing.T) {
		list := fileLines{
			&line{TYPE_COMMENT, "aaa"},
			&line{TYPE_COMMENT, ""},
			&line{TYPE_COMMENT, ""},
			&line{TYPE_COMMENT, "bbb"},
			&line{TYPE_COMMENT, "ccc"},
			&line{},
		}
		assert.Equal(t, []string{"aaa", "bbb ccc"}, list.getComment(5))
	})
	t.Run("Rm Begin,End EmptyLine", func(t *testing.T) {
		list := fileLines{
			&line{TYPE_COMMENT, ""},
			&line{TYPE_COMMENT, "aaa"},
			&line{TYPE_COMMENT, ""},
			&line{TYPE_COMMENT, "bbb"},
			&line{TYPE_COMMENT, "ccc"},
			&line{TYPE_COMMENT, ""},
			&line{},
		}
		assert.Equal(t, []string{"aaa", "bbb ccc"}, list.getComment(6))
	})
}
