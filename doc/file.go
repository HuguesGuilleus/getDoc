package doc

import (
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	TYPE_NODEF = 0
)

var extRegexp = regexp.MustCompile(".*\\.(\\w+)$")

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
	return extRegexp.ReplaceAllString(path, "$1")
}

// return the parser for the lang. If there are no parser,
// the function return nil
func langKnown(ext string) parserFunc {
	for lang, parser := range parserList {
		if ext == lang {
			return parser
		}
	}
	return nil
}

// Read and split file in a string for each line. If error, panic.
func splitFile(path string) (lines fileLines) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	for _, str := range strings.Split(string(data), "\n") {
		lines = append(lines, &line{
			Type: TYPE_NODEF,
			Str:  str,
		})
	}
	return
}
