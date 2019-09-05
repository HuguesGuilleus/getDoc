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

// One line with her type and the content
type line struct {
	Type int
	Str  string
}

// All the line of a file
type fileLines []*line

// Get the extention of a file
func getExt(path string) string {
	return extRegexp.ReplaceAllString(path, "$1")
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
