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
	"github.com/HuguesGuilleus/go-workerglobalscope/console"
	"github.com/HuguesGuilleus/go-workerglobalscope/message"
	"github.com/HuguesGuilleus/go-workerglobalscope/reflectjs/uint8array"
	"io"
	"sync"
	"syscall/js"
)

func main() {
	message.Post(struct{ Type string }{Type: "logreset"})

	var d doc.Doc
	d.Log.SetOutput(L{})

	var wg sync.WaitGroup

	for m := range message.Listen() {
		switch t := m.Get("type").String(); t {
		case "title":
			d.Title = m.Get("title").String()
			log("Reset title: " + d.Title)
		case "blob":
			wg.Add(1)
			go func(name string, b js.Value) {
				defer wg.Done()
				r := Body{Value: b}.Reader()
				d.ReadOne(name, r)
			}(m.Get("name").String(), m.Get("blob"))

		case "ask":
			wg.Wait()
			var buff bytes.Buffer
			d.SaveHTML(&buff, false)
			a := uint8array.New(buff.Bytes())

			opt := js.Global().Get("Object").New()
			opt.Set("type", "text/html; charset=utf-8")
			blob := js.Global().Get("Blob").New(js.Global().Get("Array").Call("of", a), opt)
			message.Post(struct {
				Type string
				Blob js.Value
			}{
				Type: "gen",
				Blob: blob,
			})

		case "resize", "reset":
			console.Info("Un dev message type:", t)
		default:
			console.Error("Unknown message type:", t)
		}
	}
}

/* Logger section */

// Send a line message to the js main worker, the message will be printed into the main page.
func log(line string) {
	if line == "" {
		return
	} else if line[len(line)-1] != '\n' {
		line += "\n"
	}
	message.Post(struct {
		Type string
		Line string
	}{
		Type: "log",
		Line: line,
	})
}

type L struct{}

func (_ L) Write(ms []byte) (int, error) {
	log(string(ms))
	return len(ms), nil
}

/* js BODY */

// You can use it with a response, a Blob for example.
//
// https://fetch.spec.whatwg.org/#body
type Body struct {
	js.Value
}

// TODO: from a js.value(blob) --> io.Reader

// Get an array buffer of the Body, and to each call of Read copy the bytes
// into the destination []byte.
func (b Body) Reader() io.Reader {
	buff, err := Await(b.Call("arrayBuffer"))
	if err.Truthy() {
		console.Error(err)
		return io.MultiReader()
	}
	a := uint8array.Uint8Array.New(buff)
	// a := uint8array.Uint8Array.New(b.Value)
	return &bodyReader{
		array: a,
		size:  a.Get("byteLength").Int(),
		pos:   0,
	}
}

type bodyReader struct {
	array js.Value
	size  int
	pos   int
}

func (b *bodyReader) Read(dst []byte) (int, error) {
	if b.pos >= b.size {
		return 0, io.EOF
	}

	var readed int
	if b.pos == 0 {
		readed = js.CopyBytesToGo(dst, b.array)
	} else {
		readed = js.CopyBytesToGo(dst, b.array.Call("subarray", b.pos))
	}

	b.pos += readed
	return readed, nil
}

func Await(promise js.Value) (resolve, reject js.Value) {
	c := make(chan struct{})
	defer close(c)

	then := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resolve = args[0]
		c <- struct{}{}
		return nil
	})
	defer then.Release()
	promise.Call("then", then)

	catch := js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		reject = args[0]
		c <- struct{}{}
		return nil
	})
	defer catch.Release()
	promise.Call("catch", catch)

	<-c
	return
}
