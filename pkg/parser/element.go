// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"io"
)

// Create a parser for on file.
type GenParser func(r io.Reader, path string) interface {
	// Return the next documented element, or nil if error or EOF.
	//
	// Element.FileName, Lang will be overwrite by caller.
	Next() *Element
	// Return the error (other than EOF) or nil.
	Error() error
}

// On element: function, var, typedef, class ...
type Element struct {
	// The name of the element
	Name string
	// The header of the element
	LineName string
	// The type: func, var, const, class ...
	Type string
	// The file where are the definition of the element
	FileName string
	// The line of the definition in the file
	LineNum int
	// The comment before the element. Each item is a paragraph.
	Comment []string
	// The language of the file
	Lang string
}
