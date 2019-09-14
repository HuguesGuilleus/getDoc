package main

import (
	"github.com/HuguesGuilleus/parseOpt/check"
	"testing"
)

func TestSpec(t *testing.T) {
	check.Check(t, spec)
}
