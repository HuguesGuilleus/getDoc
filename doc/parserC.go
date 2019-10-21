// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"regexp"
)

var (
	langC_comment          = regexp.MustCompile("^\\s*/{2,}\\s*(.*)")
	langC_function         = regexp.MustCompile("^([\\w* ]+\\s+\\w+\\(.*\\))[\\s{]*$")
	langC_functionName     = regexp.MustCompile("^[\\w* ]+\\s+(\\w+)\\(.*\\).*")
	langC_Typedef          = regexp.MustCompile("\\s*typedef")
	langC_TypedefSimple    = regexp.MustCompile("\\s*typedef\\s+.*\\s+([\\w]+)\\s*;\\s*$")
	langC_TypedefMultBegin = regexp.MustCompile("(typedef\\s*\\w+)\\s*{")
	langC_TypedefMultEnd   = regexp.MustCompile("\\s*}\\s*(\\w+)\\s*;")
	langC_MacroConst       = regexp.MustCompile("^\\s*(#define\\s+\\w+\\s+.+)$")
	langC_MacroName        = regexp.MustCompile("^\\s*#define\\s+(\\w+).*")
	langC_MacroFunc        = regexp.MustCompile("^\\s*(#define\\s+\\w+\\(.*\\)\\s+.+)$")
	langC_var              = regexp.MustCompile("^(\\s*)(\\w+\\s*\\*\\s*|\\w+\\s+)(\\w+)([^()]+)$")
	langC_keyWord          = regexp.MustCompile("return|typdef|if|else|do|while|for|switch|case|struct|enum")
)

func langC_parse(index *Index, lines fileLines, fileName string) {
	for i, l := range lines {
		switch l.Type {
		case TYPE_FUNCTION:
			index.push(&Element{
				Name:     langC_functionName.ReplaceAllString(l.Str, "$1"),
				LineName: l.Str,
				Type:     "func",
				FileName: fileName,
				LineNum:  i + 1,
				Comment:  lines.getComment(i),
				Lang:     "c",
			})
		case TYPE_TYPEDEF:
			if langC_TypedefSimple.MatchString(l.Str) {
				index.push(&Element{
					Name:     langC_TypedefSimple.ReplaceAllString(l.Str, "$1"),
					LineName: l.Str,
					Type:     "type",
					FileName: fileName,
					LineNum:  i + 1,
					Comment:  lines.getComment(i),
					Lang:     "c",
				})
			} else if langC_TypedefMultBegin.MatchString(l.Str) {
				name := ""
				for j := i + 1; j < len(lines); j++ {
					if langC_TypedefMultEnd.MatchString(lines[j].Str) {
						name = langC_TypedefMultEnd.ReplaceAllString(lines[j].Str, "$1")
						break
					}
				}
				index.push(&Element{
					Name:     name,
					LineName: langC_TypedefMultBegin.ReplaceAllString(l.Str, "$1"),
					Type:     "type",
					FileName: fileName,
					LineNum:  i + 1,
					Comment:  lines.getComment(i),
					Lang:     "c",
				})
			}
		case TYPE_MACROCONST:
			index.push(&Element{
				Name:     langC_MacroName.ReplaceAllString(l.Str, "$1"),
				LineName: langC_MacroConst.ReplaceAllString(l.Str, "$1"),
				Type:     "macroConst",
				FileName: fileName,
				LineNum:  i + 1,
				Comment:  lines.getComment(i),
				Lang:     "c",
			})
		case TYPE_MACROFUNC:
			index.push(&Element{
				Name:     langC_MacroName.ReplaceAllString(l.Str, "$1"),
				LineName: langC_MacroFunc.ReplaceAllString(l.Str, "$1"),
				Type:     "macroFunc",
				FileName: fileName,
				LineNum:  i + 1,
				Comment:  lines.getComment(i),
				Lang:     "c",
			})
		case TYPE_VAR:
			if fileName[len(fileName)-1] == 'c' && langC_var.ReplaceAllString(l.Str, "$1") != "" {
				break
			}
			index.push(&Element{
				Name:     langC_var.ReplaceAllString(l.Str, "$3"),
				LineName: langC_var.ReplaceAllString(l.Str, "$2$3$4"),
				Type:     "var",
				FileName: fileName,
				LineNum:  i + 1,
				Comment:  lines.getComment(i),
				Lang:     "c",
			})
		}
	}
}

// get the type of each line, and get get info.
// ex: "// aaa" --> "aaa"
func langC_type(lines fileLines) {
	skipLines := 0
	for i, line := range lines {
		if skipLines > 0 {
			line.Type = TYPE_CODE
			skipLines--
			continue
		}
		if langC_comment.MatchString(line.Str) {
			line.Type = TYPE_COMMENT
			line.Str = langC_comment.ReplaceAllString(line.Str, "$1")
		} else if langC_function.MatchString(line.Str) {
			line.Type = TYPE_FUNCTION
			line.Str = langC_function.ReplaceAllString(line.Str, "$1")
		} else if langC_Typedef.MatchString(line.Str) {
			line.Type = TYPE_TYPEDEF
			for j, l := range lines[i:] {
				if langC_TypedefMultEnd.MatchString(l.Str) {
					skipLines = j
				}
			}
		} else if langC_var.MatchString(line.Str) {
			if !langC_keyWord.MatchString(line.Str) {
				line.Type = TYPE_VAR
			} else {
				line.Type = TYPE_CODE
			}
		} else if langC_MacroConst.MatchString(line.Str) {
			line.Type = TYPE_MACROCONST
		} else if langC_MacroFunc.MatchString(line.Str) {
			line.Type = TYPE_MACROFUNC
		} else {
			line.Type = TYPE_CODE
		}
	}
}
