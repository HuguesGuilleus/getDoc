// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

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
			t.Log("   Expected:", l.A)
			t.Log("   Received:", input[i].Str)
		}
	}
}
