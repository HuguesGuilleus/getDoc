package doc

import (
	"regexp"
)

// TODO: Typedef

var (
	langC_comment      = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langC_function     = regexp.MustCompile("^([\\w* ]+\\s+\\w+\\(.*\\)).*")
	langC_functionName = regexp.MustCompile("^[\\w* ]+\\s+(\\w+)\\(.*\\).*")
	// lineTypedef  = regexp.MustCompile("\\s*typedef.*")
)

func langC_parse(index *Index, lines fileLines, fileName string) {
	langC_type(lines)
	for i, l := range lines {
		if l.Type == TYPE_FUNCTION {
			comment := lines.getComment(i)
			index.push(&Element{
				Name:     langC_functionName.ReplaceAllString(l.Str, "$1"),
				LineName: l.Str,
				Type:     "func",
				FileName: fileName,
				LineNum:  i + 1,
				Comment:  comment,
				Lang: "c",
			})
		}
	}
}

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
