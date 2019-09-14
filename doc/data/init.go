// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package data

import (
	"text/template"
)

var Index = template.Must(template.New("index").Parse(index))
