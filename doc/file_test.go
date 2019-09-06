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
	assert.Equal(t, 1, len(*ind), "length must 1")
	if len(*ind) == 1 {
		assert.Equal(t, el, (*ind)[0], "The first Element")
	}
}

func TestGetExt(t *testing.T) {
	assert.Equal(t, "c", getExt("aaa/a.ed.c"), "Input: aaa/a.ed.c")
	assert.Equal(t, "h", getExt("aaa/aded.h"), "Input: aaa/aded.h")
	assert.Equal(t, "go", getExt("aaa/ad.go"), "Input: aaa/ad.go")
	assert.Equal(t, "js", getExt("aaa/ad.js"), "Input: aaa/ad.js")
	assert.Equal(t, "", getExt(""), "Input: <nothing>")
}

func TestLangKnown(t *testing.T) {
	parserListSave := parserList
	defer func() { parserList = parserListSave }()
	f1 := func(lines fileLines, ind *Index) {}
	parserList = map[string]parserFunc{
		"a": f1,
		"b": func(lines fileLines, ind *Index) {},
	}
	// todo: verify by pointer...
	if langKnown("a") == nil {
		t.Error("A known lang")
	}
	if langKnown("zz") != nil {
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
