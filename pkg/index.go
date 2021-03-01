// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/HuguesGuilleus/getDoc/web"
	"io"
	"io/fs"
	"log"
	"path"
	"strings"
	"sync"
	"time"
)

type Doc struct {
	// Meta informations
	Title string
	Time  time.Time

	Index Index

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
func (d *Doc) Read(root string, fsys fs.FS) error {
	var wg sync.WaitGroup
	defer wg.Wait()

	return fs.WalkDir(fsys, ".", func(p string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if entry.IsDir() {
			return nil
		}

		parser := parserList[strings.TrimPrefix(path.Ext(p), ".")]
		if parser == nil {
			return nil
		}
		wg.Add(1)
		go d.readFile(fsys, root, p, &wg, parser)

		return nil
	})
}

func (d *Doc) readFile(fsys fs.FS, root, p string, wg *sync.WaitGroup, parser *parserFuncs) {
	defer wg.Done()

	f, err := fs.ReadFile(fsys, p)
	if err != nil {
		d.Log.Printf("Read %q fail: %v\n", p, err)
		return
	}

	splited := strings.Split(string(f), "\n")
	lines := make(fileLines, len(splited), len(splited))
	for i, l := range splited {
		lines[i] = &line{Str: l}
	}

	parser.Type(lines)
	parser.Parse(&d.Index, lines, p)

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
	d.Index.sort()
}
