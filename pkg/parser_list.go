// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"github.com/HuguesGuilleus/getDoc/parser"
)

var ParserList map[string]parser.Parser = parser.GetParserList()

func ParserListExt() (list []string) {
	list = make([]string, len(ParserList))
	i := 0
	for k := range ParserList {
		list[i] = k
		i++
	}
	return
}

// Get the extention of the p (file path) then get the Parse with ParserList.
func getParser(p string) parser.Parser {
	for i := len(p); i > 0; i-- {
		switch p[i-1] {
		case '/', '.':
			return ParserList[p[i:]]
		}
	}
	return ParserList[p]
}
