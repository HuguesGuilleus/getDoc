package until

import (
	"strings"
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
