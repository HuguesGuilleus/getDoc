// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"regexp"
)

var (
	langJs_comment  = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langJs_function = regexp.MustCompile("^(function\\s\\w+\\(.*\\)).*")
	langJs_class    = regexp.MustCompile("^(class\\s+\\w+).*")
	langJs_var      = regexp.MustCompile("^(var|let)\\s+\\w+.*")
	langJs_const    = regexp.MustCompile("^const\\s+\\w+.*")
	langJs_Name     = regexp.MustCompile("^(?:class|var|let|function|const)\\s+(\\w+).*")
)

// The parser function for the JavaScript.
func langJs_parse(index *Index, lines fileLines, fileName string) {
	langJs_type(lines)
	for i, l := range lines {
		ty := ""
		switch l.Type {
		case TYPE_FUNCTION:
			ty = "func"
		case TYPE_CLASS:
			ty = "class"
		case TYPE_VAR:
			ty = "var"
		case TYPE_CONST:
			ty = "const"
		default:
			continue
		}
		index.push(&Element{
			Comment:  lines.getComment(i),
			FileName: fileName,
			Lang:     "js",
			LineName: l.Str,
			LineNum:  i + 1,
			Name:     langJs_Name.ReplaceAllString(l.Str, "$1"),
			Type:     ty,
		})
	}
}

// get the type of each line, and get get info.
// ex: "// aaa" --> "aaa"
func langJs_type(lines fileLines) {
	for _, line := range lines {
		if langJs_comment.MatchString(line.Str) {
			line.Type = TYPE_COMMENT
			line.Str = langJs_comment.ReplaceAllString(line.Str, "$1")
		} else if langJs_function.MatchString(line.Str) {
			line.Type = TYPE_FUNCTION
			line.Str = langJs_function.ReplaceAllString(line.Str, "$1")
		} else if langJs_class.MatchString(line.Str) {
			line.Type = TYPE_CLASS
			line.Str = langJs_class.ReplaceAllString(line.Str, "$1")
		} else if langJs_var.MatchString(line.Str) {
			line.Type = TYPE_VAR
		} else if langJs_const.MatchString(line.Str) {
			line.Type = TYPE_CONST
		} else {
			line.Type = TYPE_CODE
		}
	}
}
