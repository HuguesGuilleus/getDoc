// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangC_parse(t *testing.T) {
	fileName := "aaa.c"
	t.Run("Function", func(t *testing.T) {
		index := Index{}
		lines := fileLines{
			&line{
				Str:  "// My function for Say Hello World.",
				Type: TYPE_COMMENT,
			},
			&line{
				Str:  "int hello() {",
				Type: TYPE_FUNCTION,
			},
		}
		elementFunc := Element{
			Name:     "hello",
			LineName: "int hello()",
			Type:     "func",
			FileName: fileName,
			LineNum:  2,
			Comment:  []string{"My function for Say Hello World."},
			Lang:     "c",
		}
		langC_type(lines)
		langC_parse(&index, lines, fileName)
		assert.Equal(t, elementFunc, *index[0], "")
	})
	t.Run("MacroConst", func(t *testing.T) {
		index := Index{}
		lines := fileLines{
			&line{
				Str: "	#define YOLO 14",
				Type: TYPE_COMMENT,
			},
		}
		element := Element{
			Name:     "YOLO",
			LineName: "#define YOLO 14",
			Type:     "macroConst",
			LineNum:  1,
			Comment:  []string{},
			Lang:     "c",
			FileName: fileName,
		}
		langC_type(lines)
		langC_parse(&index, lines, fileName)
		assert.Equal(t, element, *index[0], "")
	})
	t.Run("MacroFunc", func(t *testing.T) {
		index := Index{}
		lines := fileLines{
			&line{
				Str: "	#define ERR(xxx ...)	fprintf(stderr, xxx)",
				Type: TYPE_COMMENT,
			},
		}
		element := Element{
			Name: "ERR",
			LineName: "#define ERR(xxx ...)	fprintf(stderr, xxx)",
			Type:     "macroFunc",
			LineNum:  1,
			Comment:  []string{},
			Lang:     "c",
			FileName: fileName,
		}
		langC_type(lines)
		langC_parse(&index, lines, fileName)
		assert.Equal(t, element, *index[0], "")
	})
	t.Run("TypedefSimple", func(t *testing.T) {
		index := Index{}
		lines := fileLines{
			&line{
				Str:  "typedef int bool ;",
				Type: TYPE_TYPEDEF,
			},
		}
		elementFunc := Element{
			Name:     "bool",
			LineName: "typedef int bool ;",
			Type:     "type",
			LineNum:  1,
			Comment:  []string{},
			Lang:     "c",
			FileName: fileName,
		}
		langC_type(lines)
		langC_parse(&index, lines, fileName)
		assert.Equal(t, elementFunc, *index[0], "")
	})
	t.Run("TypedefMultlines", func(t *testing.T) {
		index := Index{}
		lines := fileLines{
			&line{
				Str:  "typedef struct {",
				Type: TYPE_TYPEDEF,
			},
			&line{
				Str: "	int swag ;",
				Type: TYPE_CODE,
			},
			&line{
				Str:  "} yolo ;",
				Type: TYPE_CODE,
			},
		}
		elementFunc := Element{
			Name:     "yolo",
			LineName: "typedef struct",
			Type:     "type",
			LineNum:  1,
			Comment:  []string{},
			Lang:     "c",
			FileName: fileName,
		}
		langC_type(lines)
		langC_parse(&index, lines, fileName)
		assert.Equal(t, elementFunc, *index[0], "")
	})
	t.Run("Var", func(t *testing.T) {
		lines := fileLines{
			&line{Str: "	int yolo = 42 ;"},
		}
		langC_type(lines)
		t.Run("golbal in .h", func(t *testing.T) {
			index := Index{}
			fileName := "aaa.h"
			elementFunc := Element{
				Name:     "yolo",
				LineName: "int yolo = 42 ;",
				Type:     "var",
				LineNum:  1,
				Comment:  []string{},
				Lang:     "c",
				FileName: fileName,
			}
			langC_parse(&index, lines, fileName)
			if assert.Equal(t, 1, len(index), "One element expected") {
				assert.Equal(t, elementFunc, *index[0])
			}
		})
		t.Run("golbal in .c", func(t *testing.T) {
			index := Index{}
			fileName := "aaa.c"
			lines := fileLines{
				&line{Str: "int yolo = 42 ;"},
			}
			elementFunc := Element{
				Name:     "yolo",
				LineName: "int yolo = 42 ;",
				Type:     "var",
				LineNum:  1,
				Comment:  []string{},
				Lang:     "c",
				FileName: fileName,
			}
			langC_type(lines)
			langC_parse(&index, lines, fileName)
			if assert.Equal(t, 1, len(index), "One element expected") {
				assert.Equal(t, elementFunc, *index[0])
			}
		})
		t.Run("local var in .c", func(t *testing.T) {
			index := Index{}
			langC_parse(&index, lines, "aaa.c")
			if len(index) != 0 {
				t.Fail()
			}
		})
	})
}

func TestLangC_Type(t *testing.T) {
	testType(t, langC_type, []testingLine{
		{TYPE_COMMENT, "/// aaa", "aaa"},
		{TYPE_COMMENT, "//", ""},
		{TYPE_CODE, "	a = 4.2 ;", "	a = 4.2 ;"},
		{TYPE_FUNCTION, "int yolo(f float) {", "int yolo(f float)"},
		{TYPE_FUNCTION, "* int yolo(f float)", "* int yolo(f float)"},
		{TYPE_TYPEDEF, "typedef int bool ;", "typedef int bool ;"},
		{TYPE_MACROCONST, "#define YOLO 42", "#define YOLO 42"},
		{TYPE_MACROFUNC, "#define ERR(xxx ...)	fprintf(stderr, xxx)", "#define ERR(xxx ...)	fprintf(stderr, xxx)"},
		{TYPE_VAR, "int yolo = 14 ;", "int yolo = 14 ;"},
		{TYPE_VAR, "int * yolo = 14 ;", "int * yolo = 14 ;"},
	})
}
