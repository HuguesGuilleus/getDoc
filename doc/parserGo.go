package doc

import (
	"regexp"
)

// TODO: search many variable def in the same line
// TODO: (langGO) push the elemnt to the index

var (
	langGo_comment        = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langGo_func           = regexp.MustCompile("^(func\\s+(?:\\(.*\\)\\s+)?\\w+\\(.*\\))\\s+{")
	langGo_typedef        = regexp.MustCompile("^(type\\s+\\w+\\s+[^{ ]*).*")
	langGo_varSimple      = regexp.MustCompile("^(var\\s+\\w+\\s+(?:[^=]*)?(?:=.*)?)")
	langGo_constSimple    = regexp.MustCompile("^(const\\s+\\w+\\s+(?:[^=]*)?(?:=.*)?)")
	langGo_varBegin       = regexp.MustCompile("^var\\s+\\(\\s*")
	langGo_constBegin     = regexp.MustCompile("^const\\s+\\(\\s*")
	langGo_varConstMiddle = regexp.MustCompile("\\s*\\w+.*")
	langGo_varConstEnd    = regexp.MustCompile("\\s*\\)\\s*")
)

func langGo_parse(index *Index, lines fileLines, fileName string) {
	langGo_type(lines)
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
