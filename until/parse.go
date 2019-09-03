package until

import (
	"regexp"
	"strings"
)

var (
	lineComment  = regexp.MustCompile("\\s*/{2,}\\s*")
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
	t := 0
	for _, line := range rawLines {
		if lineComment.MatchString(line) {
			parsedLines = append(parsedLines, &Line{
				Type: TYPE_COMMENT,
				Str:  lineComment.ReplaceAllString(line, ""),
			})
			continue
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

// Get all the commentary before a num line.
// A empty commentary procuce a new line ('\n').
func (list listLine) getComment(num int) string {
	begin := num - 1
	for ; begin >= 0; begin-- {
		if list[begin].Type != TYPE_COMMENT {
			break
		}
	}
	builder := strings.Builder{}
	begin++
	for ; begin < num; begin++ {
		str := list[begin].Str
		if len(str) == 0 {
			builder.WriteRune('\n')
		} else {
			builder.WriteString(list[begin].Str)
			builder.WriteRune(' ')
		}
	}
	return builder.String()
}
