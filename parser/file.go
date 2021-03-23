// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"io"
	"regexp"
	"strings"
)

var (
	getExtDot   = regexp.MustCompile(".*\\.(\\w+)$")
	getExtSlash = regexp.MustCompile(".*/(\\w+)$")
)

type Parser func(path string, r io.Reader, index *Index) error

// On element: function, var, typedef, class ...
type Element struct {
	// The name of the element
	Name string
	// The header of the element
	LineName string
	// The type: func, var, const, class ...
	Type string
	// The file where are the definition of the element
	FileName string
	// The line of the definition in the file
	LineNum int
	// The comment before the element. Each item is a paragraph.
	Comment []string
	// The language of the file
	Lang string
}

// All the element of a project
type Index []*Element

// push a element to an Index
func (ind *Index) push(el *Element) {
	*ind = append(*ind, el)
}

// One line with her type and the content
type line struct {
	Type int
	Str  string
}

// All the lines of a file
type fileLines []*line

// Get the extention of a file
func getExt(path string) string {
	if getExtDot.MatchString(path) {
		return getExtDot.ReplaceAllString(path, "$1")
	} else {
		return getExtSlash.ReplaceAllString(path, "$1")
	}
}

// Get all the commentary before a num line.
// A empty commentary procuce a new paragraph.
func (list fileLines) getComment(num int) (comment []string) {
	begin := num - 1
	for ; begin > -1; begin-- {
		if list[begin].Type != TYPE_COMMENT {
			break
		}
	}
	if begin == num-1 {
		return []string{}
	}
	begin++
	builder := strings.Builder{}
	for beginPara := true; begin < num; begin++ {
		if s := list[begin].Str; len(s) == 0 {
			beginPara = true
		} else {
			if beginPara == false {
				builder.WriteRune(' ')
			} else if builder.Len() != 0 {
				comment = append(comment, builder.String())
				builder.Reset()
			}
			builder.WriteString(s)
			beginPara = false
		}
	}
	return append(comment, builder.String())
}
