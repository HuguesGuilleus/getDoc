// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"testing"
)

// A line for the test
type testingLine struct {
	// The type expected
	T int
	// B is the string Before typing
	B string
	// A is the string After typing
	A string
}

// Type test the typing of a group of line
func testType(t *testing.T, fx func(fileLines), lines []testingLine) {
	input := make(fileLines, len(lines), len(lines))
	for i, l := range lines {
		input[i] = &line{
			Str: l.B,
		}
	}
	fx(input)
	for i, l := range lines {
		if l.T != input[i].Type {
			t.Errorf("Type error (line %d)", i)
			t.Log("   Input line: ", l.B)
			t.Logf("   Type: (expected: %d) %d", l.T, input[i].Type)
		}
		if l.A != input[i].Str {
			t.Errorf("String error (line %d)", i)
			t.Log("   Input line: ", l.B)
			t.Logf("   Expected: '%s'", l.A)
			t.Logf("   Received: '%s'", input[i].Str)
		}
	}
}

var (
	testingFile string
	testingLang string
)

// testParser test a parser with linesSimple. You must set testingFile
func testParser(t *testing.T, name string, linesSimple []string, mod Element) {
	t.Run(name, func(t *testing.T) {
		ret := testParserGetElement(t, linesSimple, mod.LineNum)
		// Config mod
		mod.FileName = testingFile
		mod.Comment = []string{"Comment"}
		if len(mod.Lang) == 0 {
			mod.Lang = testingLang
		}
		if mod.LineNum == 0 {
			mod.LineNum = 2
		}
		// Zone of test the value
		if ret.Name != mod.Name {
			testParserLog(t, "Name", mod.Name, ret.Name)
		}
		if mod.LineName != ret.LineName {
			testParserLog(t, "LineName", mod.LineName, ret.LineName)
		}
		if mod.Type != ret.Type {
			testParserLog(t, "Type", mod.Type, ret.Type)
		}
		if mod.LineNum != ret.LineNum {
			testParserLog(t, "LineNum", mod.LineNum, ret.LineNum)
		}
		if len(ret.Comment) != 1 || mod.Comment[0] != ret.Comment[0] {
			testParserLog(t, "Comment", mod.Comment, ret.Comment)
		}
		if mod.FileName != ret.FileName {
			testParserLog(t, "FileName", mod.FileName, ret.FileName)
		}
		if mod.Lang != ret.Lang {
			testParserLog(t, "Lang", mod.Lang, ret.Lang)
		}
	})
}

func testParserGetElement(t *testing.T, linesSimple []string, lineNum int) Element {
	// Get the parser
	testingLang = getExt(testingFile)
	parser, ok := parserList[testingLang]
	if ok != true {
		t.Error("There are not parser for:", testingFile)
		t.SkipNow()
	}
	// Config lines
	var lines fileLines
	if lineNum == 0 {
		lines = make(fileLines, len(linesSimple)+1, len(linesSimple)+1)
		lines[0] = &line{}
		for i, l := range linesSimple {
			lines[i+1] = &line{Str: l}
		}
		parser.Type(lines)
		lines[0] = &line{
			Type: TYPE_COMMENT,
			Str:  "Comment",
		}
	} else {
		lines = make(fileLines, len(linesSimple), len(linesSimple))
		for i, l := range linesSimple {
			lines[i] = &line{Str: l}
		}
		parser.Type(lines)
	}
	// Create the Element
	index := make(Index, 0, 1)
	parser.Parse(&index, lines, testingFile)
	if len(index) != 1 {
		t.Error("index passing by testingParser must have a len of 1: ", len(index))
		t.SkipNow()
	}
	return *index[0]
}

// Fail the test and Print the name, the model value and the returned value
func testParserLog(t *testing.T, name string, mod, ret interface{}) {
	t.Errorf("Error %s:", name)
	t.Logf("--- Model   : %+v", mod)
	t.Logf("+++ Returned: %+v", ret)
	t.Log()
}
