package data

import (
	"text/template"
)

var Index = template.Must(template.New("index").Parse(index))
