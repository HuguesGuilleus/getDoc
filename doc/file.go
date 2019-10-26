// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
	"time"
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

// Sort the Index of element by the Name
func (ind Index) sort() {
	sort.Slice(ind, func(i, j int) bool {
		return ind[i].Name < ind[j].Name
	})
}

// List all file who have an element of the list.
// The return list are sorted and all file are uniq
func (ind Index) ListFile() (files []string) {
	all := []string{}
	for _, el := range ind {
		all = append(all, el.FileName)
	}
	return uniq(all)
}

// List all file who have an element of the list.
// The return list are sorted and all file are uniq
func (ind Index) ListType() (files []string) {
	all := []string{}
	for _, el := range ind {
		all = append(all, el.Type)
	}
	return uniq(all)
}

// List all file who have an element of the list.
// The return list are sorted and all file are uniq
func (ind Index) ListLang() (out []string) {
	all := []string{}
	for _, el := range ind {
		all = append(all, el.Lang)
	}
	return uniq(all)
}

// Sort and remove double string
func uniq(in []string) (out []string) {
	sort.Strings(in)
	last := ""
	for _, item := range in {
		if item != last {
			out = append(out, item)
			last = item
		}
	}
	return out
}

// Get the time of the parse for use with JSON
func (ind *Index) Date() string {
	dataTime, err := time.Now().MarshalJSON()
	if err != nil {
		printErr(err)
		return ""
	}
	dataTime = dataTime[1 : len(dataTime)-1]
	return string(dataTime)
}

// Get the date for humain
func (_ *Index) HumainDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Get the Title for <h1> in template
func (_ *Index) Title() string {
	return Title
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
func langKnown(path string) *parserFuncs {
	parserListMutex.Lock()
	defer parserListMutex.Unlock()
	return parserList[getExt(path)]
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
