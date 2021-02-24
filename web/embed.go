package webdata

import (
	"bytes"
	"embed"
	"html/template"
)

var (
	//go:embed style
	styleFS embed.FS
	//go:embed js
	jsFS embed.FS
	//go:embed index.gohtml
	index []byte
	// The template to write the list of elements in HTML.
	Index *template.Template
)

func init() {
	index = bytes.Replace(index, []byte("/*STYLE*/"), readFs(styleFS), 1)
	index = bytes.Replace(index, []byte("/*JS*/"), readFs(jsFS), 1)
	Index = template.Must(template.New("index").Parse(string(index)))
}
