package until

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLines(t *testing.T) {
	expected := listLine{
		&Line{
			Type: TYPE_CODE,
			Str:  "int a = 3",
		},
		&Line{
			Type: TYPE_CODE,
			Str:  "",
		},
		&Line{
			Type: TYPE_COMMENT,
			Str:  "// trree comment",
		},
		&Line{
			Type: TYPE_COMMENT,
			Str: "	//",
		},
		&Line{
			Type: TYPE_COMMENT,
			Str:  "//",
		},
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
}
