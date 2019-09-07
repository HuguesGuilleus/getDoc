package doc

import (
	"sync"
)

type parserFunc func(index *Index, lines fileLines, fileName string)

var parserList = map[string]parserFunc{
	"c": langC_parse,
	"h": langC_parse,
}

// Simple secure fo r the list
var parserListMutex sync.Mutex
