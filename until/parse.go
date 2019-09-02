package until

import (
	// "strings"
	"regexp"
)

var (
	lineComment  = regexp.MustCompile("\\s*//.*")
	lineFunction = regexp.MustCompile("\\s*[\\w\\[\\]*]+\\s+\\w+\\(.*\\).*")
	lineTypedef  = regexp.MustCompile("\\s*typedef.*")
)

const (
	TYPE_CODE     = 0
	TYPE_COMMENT  = 1
	TYPE_FUNCTION = 2
	TYPE_TYPEDEF  = 3
)

type Line struct {
	Type int
	Str  string
}

type listLine []*Line

// Choose the type of each line
func parseLines(rawLines []string) (parsedLines listLine) {
	for _, line := range rawLines {
		t := 0
		if lineComment.MatchString(line) {
			t = TYPE_COMMENT
		} else if lineFunction.MatchString(line) {
			t = TYPE_FUNCTION
		} else if lineTypedef.MatchString(line) {
			t = TYPE_TYPEDEF
		} else {
			t = TYPE_CODE
		}
		parsedLines = append(parsedLines, &Line{
			Type: t,
			Str:  line,
		})
	}
	return
}
