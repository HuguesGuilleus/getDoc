package until

import (
	"strings"
	"sort"
)

type Item struct {
	Name string
	FileName string
	LineNum int
	Comment []string
	Type string
}

// A list a Item
type Index []*Item

// Read the file of path, axtract and add the doc to the index
func (ind *Index) AddFile(path string) {
	allLines := parseLines(splitFile(path))
	for i,line := range allLines {
		if line.Type == TYPE_FUNCTION {
			comments := allLines.getComment(i)
			*ind = append(*ind, &Item{
				Name: line.Str,
				FileName: path,
				LineNum: i,
				Comment: strings.Split(comments, "\n"),
				Type: "func",
			})
		} else if line.Type == TYPE_TYPEDEF {
			comments := allLines.getComment(i)
			*ind = append(*ind, &Item{
				Name: line.Str,
				FileName: path,
				LineNum: i,
				Comment: strings.Split(comments, "\n"),
				Type: "type",
			})
		}
	}
}

// List all file documented in the Index
func (ind Index) ListFile() (new []string) {
	all := []string{}
	for _, item := range ind {
		all = append(all, item.FileName)
	}
	sort.Strings(all)
	last := ""
	for _, file := range all {
		if last != file {
			new = append(new, file)
			last = file
		}
	}
	return
}
