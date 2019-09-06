package doc

import (
	"regexp"
)

// TODO: Typedef

var (
	langC_comment  = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langC_function = regexp.MustCompile("^([\\w* ]+\\s+\\w+\\(.*\\)).*")
	// lineTypedef  = regexp.MustCompile("\\s*typedef.*")
)

// get the type of each line, and get get info.
// ex: "// aaa" --> "aaa"
func langC_type(lines fileLines) {
	for _, line := range lines {
		if langC_comment.MatchString(line.Str) {
			line.Type = TYPE_COMMENT
			line.Str = langC_comment.ReplaceAllString(line.Str, "$1")
			continue
		} else if langC_function.MatchString(line.Str) {
			line.Type = TYPE_FUNCTION
			line.Str = langC_function.ReplaceAllString(line.Str, "$1")
		} else {
			line.Type = TYPE_CODE
		}
	}
}
