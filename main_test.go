// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"github.com/HuguesGuilleus/parseOpt/check"
	"testing"
)

func TestSpec(t *testing.T) {
	check.Check(t, spec)
}
