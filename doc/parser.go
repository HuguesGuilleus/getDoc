// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"sync"
)

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

type parserFunc func(index *Index, lines fileLines, fileName string)

// A parser for one language
type parserFuncs struct {
	Parse func(index *Index, lines fileLines, fileName string)
	Type  func(fileLines)
}

var parserList = map[string]*parserFuncs{
	// "c":    langC_parse,
	// "h":    langC_parse,
	// "js":   langJs_parse,
	// "go":   langGo_parse,
	// "sh":   langBash_parse,
	// "bash": langBash_parse,
	"bash": &parserFuncs{
		Parse: langBash_parse,
		Type:  langBash_type,
	},
}

// var parserList = map[string]parserFunc{
// 	"c":    langC_parse,
// 	"h":    langC_parse,
// 	"js":   langJs_parse,
// 	"go":   langGo_parse,
// 	"sh":   langBash_parse,
// 	"bash": langBash_parse,
// }

// Simple secure for the list
var parserListMutex sync.Mutex
