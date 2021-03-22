// getDoc
// 2019, 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"regexp"
)

var (
	langGo_comment        = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langGo_func           = regexp.MustCompile("^(func\\s+(?:\\(.*\\)\\s+)?\\w+\\(.*\\))\\s+{")
	langGo_funcName       = regexp.MustCompile("^func\\s+(?:\\(.*\\)\\s+)?(\\w+)\\(.*\\)")
	langGo_typedef        = regexp.MustCompile("^(type\\s+\\w+\\s+[^{ ]*).*")
	langGo_varSimple      = regexp.MustCompile("^(var\\s+\\w+\\s+(?:[^=]*)?(?:=.*)?)")
	langGo_constSimple    = regexp.MustCompile("^(const\\s+\\w+\\s+(?:[^=]*)?(?:=.*)?)")
	langGo_varBegin       = regexp.MustCompile("^var\\s+\\(\\s*")
	langGo_constBegin     = regexp.MustCompile("^const\\s+\\(\\s*")
	langGo_varConstMiddle = regexp.MustCompile("\\s*\\w+.*")
	langGo_varConstEnd    = regexp.MustCompile("^\\s*\\)\\s*$")
	langGo_Name           = regexp.MustCompile("\\s*(?:type|var|const)?\\s*(\\w+).*")
	langGo_Space          = regexp.MustCompile("^\\s*(.*)\\s*$")
)

func langGo_parse(index *Index, lines fileLines, fileName string) {
	for i, l := range lines {
		var ty, name string
		switch l.Type {
		case TYPE_FUNCTION:
			name = langGo_funcName.ReplaceAllString(l.Str, "$1")
			ty = "func"
		case TYPE_TYPEDEF:
			name = langGo_Name.ReplaceAllString(l.Str, "$1")
			ty = "type"
		case TYPE_VAR:
			name = langGo_Name.ReplaceAllString(l.Str, "$1")
			ty = "var"
		case TYPE_CONST:
			name = langGo_Name.ReplaceAllString(l.Str, "$1")
			ty = "const"
		default:
			continue
		}
		index.push(&Element{
			Name:     name,
			LineName: l.Str,
			FileName: fileName,
			Comment:  lines.getComment(i),
			Type:     ty,
			LineNum:  i + 1,
			Lang:     "go",
		})
	}
}

func langGo_type(lines fileLines) {
	var mode int = TYPE_NODEF
	for _, l := range lines {
		switch mode {
		case TYPE_VAR, TYPE_CONST:
			if langGo_comment.MatchString(l.Str) {
				l.Type = TYPE_COMMENT
				l.Str = langGo_comment.ReplaceAllString(l.Str, "$1")
			} else if langGo_varConstEnd.MatchString(l.Str) {
				l.Type = TYPE_CODE
				mode = TYPE_NODEF
			} else if langGo_varConstMiddle.MatchString(l.Str) {
				l.Str = langGo_Space.ReplaceAllString(l.Str, "$1")
				l.Type = mode
			} else {
				l.Type = TYPE_CODE
			}
		default:
			if langGo_comment.MatchString(l.Str) {
				l.Type = TYPE_COMMENT
				l.Str = langGo_comment.ReplaceAllString(l.Str, "$1")
			} else if langGo_func.MatchString(l.Str) {
				l.Type = TYPE_FUNCTION
				l.Str = langGo_func.ReplaceAllString(l.Str, "$1")
			} else if langGo_typedef.MatchString(l.Str) {
				l.Type = TYPE_TYPEDEF
				l.Str = langGo_typedef.ReplaceAllString(l.Str, "$1")
			} else if langGo_varSimple.MatchString(l.Str) {
				l.Type = TYPE_VAR
			} else if langGo_constSimple.MatchString(l.Str) {
				l.Type = TYPE_CONST
			} else if langGo_varBegin.MatchString(l.Str) {
				l.Type = TYPE_CODE
				mode = TYPE_VAR
			} else if langGo_constBegin.MatchString(l.Str) {
				l.Type = TYPE_CODE
				mode = TYPE_CONST
			} else {
				l.Type = TYPE_CODE
			}
		}
	}
}
