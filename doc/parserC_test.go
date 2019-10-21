// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"testing"
)

func TestLangC_parse(t *testing.T) {
	testingFile = "prog.c"
	testParser(t, "Function", []string{"int hello() {"}, Element{
		Name:     "hello",
		LineName: "int hello()",
		Type:     "func",
	})
	testParser(t, "MacroConst", []string{"#define YOLO 14"}, Element{
		Name:     "YOLO",
		LineName: "#define YOLO 14",
		Type:     "macroConst",
	})
	testParser(t, "MacroFunc", []string{"	#define ERR(xxx ...)	fprintf(stderr, xxx)"}, Element{
		Name: "ERR",
		LineName: "#define ERR(xxx ...)	fprintf(stderr, xxx)",
		Type: "macroFunc",
	})
	testParser(t, "TypedefSimple", []string{"typedef int bool ;"}, Element{
		Name:     "bool",
		LineName: "typedef int bool ;",
		Type:     "type",
	})
	testParser(t, "TypedefMultlines", []string{
		"typedef struct {",
		"	int swag ;",
		"} yolo ;",
	}, Element{
		Name:     "yolo",
		LineName: "typedef struct",
		Type:     "type",
	})
	testParser(t, "VarGlobal.c", []string{"int yolo = 42 ;"}, Element{
		Name:     "yolo",
		LineName: "int yolo = 42 ;",
		Type:     "var",
	})
	testingFile = "prog.h"
	testParser(t, "VarGlobal.h", []string{"	int yolo = 42 ;"}, Element{
		Name:     "yolo",
		LineName: "int yolo = 42 ;",
		Type:     "var",
		Lang:     "c",
	})
	t.Run("LocalVar in .c", func(t *testing.T) {
		index := Index{}
		langC_parse(&index, fileLines{
			&line{
				Str: "	int yolo = 42 ;",
				Type: TYPE_VAR,
			},
		}, "aaa.c")
		if len(index) != 0 {
			t.Fail()
		}
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
