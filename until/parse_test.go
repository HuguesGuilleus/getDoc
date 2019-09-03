package until

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLines(t *testing.T) {
	t.Run("Code", func(t *testing.T) {
		exp := listLine{
			&Line{
				Type: TYPE_CODE,
				Str:  "int a = 3",
			},
			&Line{
				Type: TYPE_CODE,
				Str:  "",
			},
		}
		assert.Equal(t, parseLines([]string{"int a = 3", ""}), exp, "")
	})
	t.Run("Comment", func(t *testing.T) {
		input := []string{
			"// three comment",
			"///// big comment",
			"	//",
			"//",
		}
		exp := listLine{
			&Line{
				Type: TYPE_COMMENT,
				Str:  "three comment",
			},
			&Line{
				Type: TYPE_COMMENT,
				Str:  "big comment",
			},
			&Line{
				Type: TYPE_COMMENT,
				Str:  "",
			},
			&Line{
				Type: TYPE_COMMENT,
				Str:  "",
			},
		}
		assert.Equal(t, parseLines(input), exp, "Code")
	})
	t.Run("Fonction&Typedef", func(t *testing.T) {
		expected := listLine{
			&Line{
				Type: TYPE_FUNCTION,
				Str:  "*int yolo(int argc, char* []argv) {",
			},
			&Line{
				Type: TYPE_TYPEDEF,
				Str:  "typedef int struct {",
			},
		}
		input := make([]string, len(expected))
		for i, e := range expected {
			input[i] = e.Str
		}
		assert.Equal(t, parseLines(input), expected, "")
	})
}

func TestGetComment(t *testing.T) {
	t.Run("Get Before", func(t *testing.T) {
		before := []int{
			TYPE_CODE,
			TYPE_FUNCTION,
			TYPE_TYPEDEF,
		}
		for _, ty := range before {
			list := listLine{
				&Line{
					Type: ty,
					Str:  "xxx",
				},
				&Line{
					Type: TYPE_COMMENT,
					Str:  "aaa",
				},
				&Line{
					Type: TYPE_COMMENT,
					Str:  "bbb",
				},
				&Line{},
			}
			assert.Equal(t, "aaa bbb ", list.getComment(3), "")
		}
	})
	t.Run("Mult Line", func(t *testing.T) {
		list := listLine{
			&Line{
				Type: TYPE_COMMENT,
				Str:  "aaa",
			},
			&Line{
				Type: TYPE_COMMENT,
				Str:  "",
			},
			&Line{
				Type: TYPE_COMMENT,
				Str:  "bbb",
			},
			&Line{},
		}
		assert.Equal(t, "aaa \nbbb ", list.getComment(3), "")
	})
}
