// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/HuguesGuilleus/getDoc/pkg/parser"
	"github.com/HuguesGuilleus/getDoc/web"
	"io"
	"io/fs"
	"log"
	"path"
	"sync"
	"time"
)

type Doc struct {
	// Meta informations
	Title string
	Time  time.Time

	Index []*parser.Element

	// The logger. To print nothing: SetOutput(io.Discard)
	Log log.Logger `json:"-" xml:"-"`
}

// Init the doc information for safe use.
func (d *Doc) init() {
	if d.Log.Writer() == nil {
		d.Log.SetOutput(io.Discard)
	}
	if d.Time.IsZero() {
		d.Time = time.Now().UTC().Truncate(time.Second)
	}
}

// Get the documentation from files.
func (d *Doc) Read(fsys fs.FS) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	return fs.WalkDir(fsys, ".", func(p string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if entry.IsDir() {
			return nil
		}

		parser := Parsers[path.Ext(p)]
		if parser == nil {
			return nil
		}
		wg.Add(1)
		go d.readFile(fsys, p, &wg, parser)

		return nil
	})
}

func (d *Doc) readFile(fsys fs.FS, p string, wg *sync.WaitGroup, getParser parser.GenParser) {
	defer wg.Done()

	f, err := fsys.Open(p)
	if err != nil {
		d.Log.Printf("Read %q fail: %v\n", p, err)
		return
	}
	defer f.Close()

	par := getParser(f, p)
	for {
		if e := par.Next(); e != nil {
			d.Index = append(d.Index, e)
		} else {

			break
		}
	}
	if err := par.Error(); err != nil {
		d.Log.Printf("Error from %q: %v\n", p, err)
	}

	d.Log.Println("Read", p)
}

func (d *Doc) SaveHTML(w io.Writer) error {
	d.save(w, "HTML")
	return webdata.Index.Execute(w, &d.Index)
}
func (d *Doc) SaveJSON(w io.Writer) error {
	d.save(w, "JSON")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc.Encode(d)
}
func (d *Doc) SaveXML(w io.Writer) error {
	d.save(w, "XML")
	enc := xml.NewEncoder(w)
	enc.Indent("", "\t")
	return enc.Encode(d)
}

// Log the output save and sort the index.
func (d *Doc) save(w interface{}, format string) {
	if n, ok := w.(interface{ Name() string }); ok {
		d.Log.Printf("Save in %s in %q", format, n.Name())
	} else {
		d.Log.Println("Save in", format)
	}
	// d.Index.sort()
	// TODO: sort Element
}
