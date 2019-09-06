package doc

import ()

type parserFunc func(index *Index, lines fileLines, fileName string)

var parserList = map[string]parserFunc{
	"c": langC_parse,
}
