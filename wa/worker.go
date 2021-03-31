// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

// +build js,wasm

// To build this file:
// GOOS=js GOARCH=wasm go build -o worker.wasm worker.go

package main

import (
	"bytes"
	"github.com/HuguesGuilleus/getDoc/pkg"
	"github.com/HuguesGuilleus/go-workerglobalscope"
	"github.com/HuguesGuilleus/go-workerglobalscope/console"
	"github.com/HuguesGuilleus/go-workerglobalscope/message"
	"strings"
	"syscall/js"
)

func main() {
	logReset("[INIT]\n[EXTENTION] " +
		strings.Join(doc.ParserListExt(), ", ") +
		"\n\n")
	message.Post(struct {
		Type string   `js:"type"`
		Ext  []string `js:"ext"`
	}{Type: "ext", Ext: doc.ParserListExt()})

	var d doc.Doc
	d.Log.SetOutput(L{})

	for m := range message.Listen() {
		switch t := m.Get("type").String(); t {
		case "title":
			d.Title = m.Get("title").String()
			log("[TITLE] " + d.Title)
		case "blob":
			go func(m js.Value) {
				r, err := ws.ReadBody(m.Get("blob"))
				if err != nil {
					log("[ERROR] get blob fail: " + err.Error() + "\n")
					return
				}
				d.ReadOne(m.Get("name").String(), r)
			}(m)
		case "ask":
			var buff bytes.Buffer
			var t string
			switch m.Get("format").String() {
			case "html":
				d.SaveHTML(&buff, false)
				t = "text/html; charset=utf-8"
			case "json":
				d.SaveJSON(&buff, false)
				t = "application/json"
			case "xml":
				d.SaveXML(&buff, false)
				t = "application/xml"
			default:
				console.Error("Unknwon saver format:", m.Get("format").String())
				continue
			}
			message.Post(struct {
				Type string   `js:"type"`
				Blob js.Value `js:"blob"`
			}{
				Type: "doc",
				Blob: ws.NewBlob(t, buff.Bytes()),
			})
		case "reset":
			d = doc.Doc{}
			d.Log.SetOutput(L{})
			logReset("[RESET]\n")
		default:
			console.Error("Unknown message type:", t)
		}
	}
}

/* Logger section */

func logReset(ms string) {
	message.Post(struct {
		Type string `js:"type"`
		Text string `js:"text"`
	}{Type: "logReset", Text: ms})
}

// Send a line message to the js main worker, the message will be printed into the main page.
func log(line string) {
	if line == "" {
		return
	} else if line[len(line)-1] != '\n' {
		line += "\n"
	}
	message.Post(struct {
		Type string `js:"type"`
		Line string `js:"line"`
	}{Type: "logLine", Line: line})
}

// A fake writer used in Doc.Log.SetOutput
type L struct{}

func (_ L) Write(ms []byte) (int, error) {
	log(string(ms))
	return len(ms), nil
}
