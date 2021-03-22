// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"github.com/HuguesGuilleus/getDoc/parser"
)

var ParserList map[string]parser.Parser = parser.GetParserList()
