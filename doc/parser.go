package doc

import ()

type parserFunc func(lines fileLines, index *Index)

var parserList = map[string]parserFunc{}
