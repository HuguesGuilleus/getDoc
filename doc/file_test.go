package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExt(t *testing.T) {
	assert.Equal(t, "c", getExt("aaa/a.ed.c"), "Input: aaa/a.ed.c")
	assert.Equal(t, "h", getExt("aaa/aded.h"), "Input: aaa/aded.h")
	assert.Equal(t, "go", getExt("aaa/ad.go"), "Input: aaa/ad.go")
	assert.Equal(t, "js", getExt("aaa/ad.js"), "Input: aaa/ad.js")
	assert.Equal(t, "", getExt(""), "Input: <nothing>")
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
