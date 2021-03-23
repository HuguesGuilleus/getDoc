// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const testFileName = "filetest.x"

// TestParserOne call the p Parser with the input and test that the return
// is equal to Expected. If name is not empty, call a sub test with name.
func TestParserOne(t *testing.T, name string, p Parser, expected Element, input string) {
	TestParser(t, name, p, Index{&expected}, input)
}

func TestParser(t *testing.T, name string, p Parser, expected Index, input string) {
	if name != "" {
		t.Run(name, func(t *testing.T) {
			TestParser(t, "", p, expected, input)
		})
		return
	}

	var index Index
	err := p(testFileName, strings.NewReader(input), &index)
	if err != nil {
		t.Error("Unexpected error:", err)
		return
	}

	for _, e := range index {
		e.Lang = ""
		if len(e.Comment) == 0 {
			e.Comment = nil
		}
	}

	for _, e := range expected {
		e.FileName = testFileName
	}

	assert.Equal(t, expected, index)
}
