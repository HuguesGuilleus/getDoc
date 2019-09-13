package doc

import (
	"regexp"
)

var (
	langBash_comment         = regexp.MustCompile("^\\s*#+\\s*(.*)")
	langBash_shebang         = regexp.MustCompile("^#!")
	langBash_funcFunction    = regexp.MustCompile("^function\\s+\\w+")
	langBash_funcParenthesis = regexp.MustCompile("^\\w+\\(\\)")
	langBash_var             = regexp.MustCompile("^(?:export\\s+)?\\w+=")
	langBash_name            = regexp.MustCompile("^(?:function\\s+|export\\s+)?(\\w+).*")
)

func langBash_parse(index *Index, lines fileLines, fileName string) {
	langBash_type(lines)
	var ty string
	for i, l := range lines {
		switch l.Type {
		case TYPE_FUNCTION:
			ty = "func"
		case TYPE_VAR:
			ty = "var"
		default:
			continue
		}
		index.push(&Element{
			Name:     langBash_name.ReplaceAllString(l.Str, "$1"),
			LineName: l.Str,
			FileName: fileName,
			LineNum:  i + 1,
			Type:     ty,
			Comment:  lines.getComment(i),
			Lang:     "bash",
		})
	}
}

func langBash_type(lines fileLines) {
	for _, l := range lines {
		if langBash_comment.MatchString(l.Str) {
			if langBash_shebang.MatchString(l.Str) {
				l.Type = TYPE_CODE
			} else {
				l.Str = langBash_comment.ReplaceAllString(l.Str, "$1")
				l.Type = TYPE_COMMENT
			}
		} else if langBash_funcFunction.MatchString(l.Str) {
			l.Type = TYPE_FUNCTION
		} else if langBash_funcParenthesis.MatchString(l.Str) {
			l.Type = TYPE_FUNCTION
		} else if langBash_var.MatchString(l.Str) {
			l.Type = TYPE_VAR
		} else {
			l.Type = TYPE_CODE
		}
	}
}
