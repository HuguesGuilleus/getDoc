// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package parser

import (
	"io"
)

// Parse the file, and save the data to the Index.
type Parser func(r io.Reader, file string, index Index) error

// Used to save fonded elements.
type Index interface {
	Element(*Element)
}

// On element: function, var, typedef, class ...
type Element struct {
	ID
	Position

	// And extern element of it package/class can access in read or
	// write to this element.
	Public bool
	// The type: func, var, const, class ...
	Type string
	// The header of the element
	Code []string

	// The comment before the element. Each item is a paragraph.
	Comments []Comment

	// Associed constructor, field, method, class implementation...
	// Brief, all element with a public and direct link.
	DirectChild []*Element

	// Intern usage of functions, structure ...
	InternUse []*Usage
}

type ID struct {
	// The module, package depend of langage semantic.
	NameSpace string
	// The name of the element
	Name string
}

// The position of an element or a
type Position struct {
	// The file where are the definition of the element
	File string
	// The line of the definition in the file
	Line int
	// The language of the file
	Lang string
}

// The usage of an element to other.
type Usage struct {
	ID       // the usage of the used element
	Position // the position of the usage.
}
