// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

const (
	TYPE_NODEF      = 0
	TYPE_COMMENT    = 1
	TYPE_CODE       = 2
	TYPE_FUNCTION   = 3
	TYPE_TYPEDEF    = 4
	TYPE_MACROCONST = 5
	TYPE_MACROFUNC  = 6
	TYPE_CONST      = 7
	TYPE_CLASS      = 8
	TYPE_VAR        = 9
)

var nameType = map[int]string{
	0: "?",
	1: "COMM",
	2: "|",
	3: "FUNC",
	4: "TYPE",
	5: "MACROF",
	6: "MACROC",
	7: "CONST",
	8: "CLASS",
	9: "VAR",
}

type parserFunc func(index *Index, lines fileLines, fileName string)

// A parser for one language
type parserFuncs struct {
	Parse func(index *Index, lines fileLines, fileName string)
	Type  func(fileLines)
}

var parserList = map[string]*parserFuncs{
	"bash": &parserFuncs{
		Parse: langBash_parse,
		Type:  langBash_type,
	},
	"sh": &parserFuncs{
		Parse: langBash_parse,
		Type:  langBash_type,
	},
	"c": &parserFuncs{
		Parse: langC_parse,
		Type:  langC_type,
	},
	"h": &parserFuncs{
		Parse: langC_parse,
		Type:  langC_type,
	},
	"js": &parserFuncs{
		Parse: langJs_parse,
		Type:  langJs_type,
	},
	"go": &parserFuncs{
		Parse: langGo_parse,
		Type:  langGo_type,
	},
}
