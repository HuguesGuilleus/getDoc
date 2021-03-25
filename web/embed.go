package webdata

import (
	"bytes"
	"embed"
	"github.com/HuguesGuilleus/static.v3"
	"html/template"
	"io/fs"
)

var (
	//go:embed style
	styleFS embed.FS
	//go:embed js
	jsFS embed.FS
	//go:embed index.gohtml
	index []byte
	// The template to write the list of elements in HTML.
	IndexMin *template.Template
	IndexStd *template.Template
)

func init() {
	index = bytes.Replace(index, []byte("/*STYLE*/"), readFs(styleFS, static.CssMinify), 1)
	index = bytes.Replace(index, []byte("/*JS*/"), readFs(jsFS, static.JsMinify), 1)

	IndexStd = template.Must(template.New("index").Parse(string(index)))
	index = static.HtmlMinify(index)
	IndexMin = template.Must(template.New("index").Parse(string(index)))
}

func readFs(s fs.FS, m static.Minifier) []byte {
	var buff bytes.Buffer
	fs.WalkDir(s, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		f, err := fs.ReadFile(s, path)
		if err != nil {
			panic(err)
		}
		buff.Write(m(f))
		return nil
	})
	return buff.Bytes()
}
