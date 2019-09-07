package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLangC_parse(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		index := Index{}
		fileName := "aaa.c"
		lines := fileLines{
			&line{
				Str:  "// My function por Say Hello World.",
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
			Comment:  []string{"My function por Say Hello World."},
			Lang:     "c",
		}
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
		}
		langC_parse(&index, lines, "")
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
			Type:    "macroFunc",
			LineNum: 1,
			Comment: []string{},
			Lang:    "c",
		}
		langC_parse(&index, lines, "")
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
		}
		langC_parse(&index, lines, "")
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
		}
		langC_parse(&index, lines, "")
		assert.Equal(t, elementFunc, *index[0], "")
	})
}

func TestLangCType(t *testing.T) {
	input := fileLines{
		&line{Str: "/// aaa"},
		&line{Str: "//"},
		&line{Str: "	a = 4.2 ;"},
		&line{Str: "int yolo(f float) {"},
		&line{Str: "* int yolo(f float)"},
		&line{Str: "typedef int bool ;"},
		&line{Str: "#define YOLO 42"},
		&line{Str: "#define ERR(xxx ...)	fprintf(stderr, xxx)"},
	}
	langC_type(input)
	assert.Equal(t, fileLines{
		&line{
			Str:  "aaa",
			Type: TYPE_COMMENT,
		},
		&line{
			Str:  "",
			Type: TYPE_COMMENT,
		},
		&line{
			Str: "	a = 4.2 ;",
			Type: TYPE_CODE,
		},
		&line{
			Str:  "int yolo(f float)",
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  "* int yolo(f float)",
			Type: TYPE_FUNCTION,
		},
		&line{
			Str:  "typedef int bool ;",
			Type: TYPE_TYPEDEF,
		},
		&line{
			Str:  "#define YOLO 42",
			Type: TYPE_MACROCONST,
		},
		&line{
			Str: "#define ERR(xxx ...)	fprintf(stderr, xxx)",
			Type: TYPE_MACROFUNC,
		},
	}, input, "")
}
