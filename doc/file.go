package doc

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

const (
	TYPE_NODEF    = 0
	TYPE_COMMENT  = 1
	TYPE_CODE     = 2
	TYPE_FUNCTION = 3
	TYPE_TYPEDEF  = 4
)

var (
	getExtDot   = regexp.MustCompile(".*\\.(\\w+)$")
	getExtSlash = regexp.MustCompile(".*/(\\w+)$")
)

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

// List all file who have an element of the list.
// The return list are sorted and all file are uniq
func (ind Index) ListFile() (files []string) {
	all := []string{}
	for _, el := range ind {
		all = append(all, el.FileName)
	}
	sort.Strings(all)
	last := ""
	for _, file := range all {
		if last != file {
			files = append(files, file)
			last = file
		}
	}
	return
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

// return the parser for the lang. If there are no parser,
// the function return nil
func langKnown(ext string) parserFunc {
	parserListMutex.Lock()
	defer parserListMutex.Unlock()
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

// Get all the commentary before a num line.
// A empty commentary procuce a new line ('\n').
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
	builder := strings.Builder{}
	for begin++; begin < num; begin++ {
		str := list[begin].Str
		if len(str) == 0 {
			comment = append(comment, builder.String())
			builder.Reset()
		} else {
			if builder.Len() != 0 {
				builder.WriteRune(' ')
			}
			builder.WriteString(list[begin].Str)
		}
	}
	return append(comment, builder.String())
}
