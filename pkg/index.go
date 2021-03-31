// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/HuguesGuilleus/getDoc/parser"
	"github.com/HuguesGuilleus/getDoc/web"
	"io"
	"io/fs"
	"log"
	"path"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
)

type Doc struct {
	// Meta informations
	Title string
	Time  time.Time

	Index    parser.Index
	ListFile []string
	ListLang []string
	ListType []string

	// Reset in each Doc.init (to each reading), then Save function use it
	// to generate ListFile, ListLang, ListType with the function Doc.genList
	onceListGen sync.Once

	// The logger. To print nothing: SetOutput(io.Discard)
	Log log.Logger `json:"-" xml:"-"`

	wg sync.WaitGroup
}

// Init the doc information for safe use before reading
func (d *Doc) init() {
	if d.Log.Writer() == nil {
		d.Log.SetOutput(io.Discard)
	}
	if d.Time.IsZero() {
		d.Time = time.Now().UTC().Truncate(time.Second)
	}
	d.onceListGen = sync.Once{}
}

// Get the documentation from files.
func (d *Doc) Read(root string, fsys fs.FS) error {
	d.init()
	return fs.WalkDir(fsys, ".", func(p string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		} else if entry.IsDir() {
			return nil
		}

		parser := getParser(p)
		if parser == nil {
			return nil
		}
		d.wg.Add(1)
		go d.readFile(fsys, root, p, parser)
		return nil
	})
}

func (d *Doc) readFile(fsys fs.FS, root, p string, parser parser.Parser) {
	f, err := fsys.Open(p)
	if err != nil {
		d.Log.Printf("[ERROR] %q fail: %v\n", p, err)
		d.wg.Done()
		return
	}
	defer f.Close()
	d.readOneReader(path.Join(root, p), f, parser)
}

// ReadOne gets a parser for the file root and use it one the reader r
// and add the Element to the doc Index. The parsing is in an other goroutine.
func (d *Doc) ReadOne(root string, r io.Reader) {
	d.init()
	parser := getParser(root)
	if parser == nil {
		return
	}
	d.wg.Add(1)
	go d.readOneReader(root, r, parser)
}

func (d *Doc) readOneReader(root string, r io.Reader, parser parser.Parser) {
	defer d.wg.Done()
	d.Log.Println("[READ]", root)
	if err := parser(root, r, &d.Index); err != nil {
		d.Log.Printf("[ERROR] parse %q fail: %v\n", root, err)
		return
	}
}

func (d *Doc) SaveHTML(w io.Writer, indent bool) error {
	d.save(w, "HTML")
	if indent {
		return webdata.IndexStd.Execute(w, d)
	} else {
		return webdata.IndexMin.Execute(w, d)
	}
}
func (d *Doc) SaveJSON(w io.Writer, indent bool) error {
	d.save(w, "JSON")
	enc := json.NewEncoder(w)
	if indent {
		enc.SetIndent("", "\t")
	}
	return enc.Encode(d)
}
func (d *Doc) SaveXML(w io.Writer, indent bool) error {
	d.save(w, "XML")
	enc := xml.NewEncoder(w)
	if indent {
		enc.Indent("", "\t")
	}
	return enc.Encode(d)
}

// Wait all file parsing, log the save event and sort the index.
func (d *Doc) save(w interface{}, format string) {
	d.wg.Wait()
	d.onceListGen.Do(d.genList)

	if n, ok := w.(interface{ Name() string }); ok {
		d.Log.Printf("[SAVE:%s] %s\n", format, n.Name())
	} else {
		d.Log.Printf("[SAVE:%s]\n", format)
	}
}

// Generate all list
func (d *Doc) genList() {
	// Sort Doc.Index
	// lower concats into buff (name + '\0' + file) in lower case.
	lower := func(buff *strings.Builder, name, file string) {
		buff.Reset()
		for _, r := range name {
			buff.WriteRune(unicode.ToLower(r))
		}
		buff.WriteByte(0)
		for _, r := range file {
			buff.WriteRune(unicode.ToLower(r))
		}
	}
	var s1, s2 strings.Builder
	sort.Slice(d.Index, func(i int, j int) bool {
		lower(&s1, d.Index[i].Name, d.Index[i].FileName)
		lower(&s2, d.Index[j].Name, d.Index[j].FileName)
		return s1.String() < s2.String()
	})

	// Generate ListFile, ListLang, ListType
	m := make(map[string]bool, len(d.Index))
	cat := func(m map[string]bool) (l []string) {
		l = make([]string, len(m))
		i := 0
		for k := range m {
			l[i] = k
			i++
			delete(m, k)
		}
		sort.Strings(l)
		return l
	}

	for _, e := range d.Index {
		m[e.FileName] = true
	}
	d.ListFile = cat(m)

	for _, e := range d.Index {
		m[e.Lang] = true
	}
	d.ListLang = cat(m)

	for _, e := range d.Index {
		m[e.Type] = true
	}
	d.ListType = cat(m)
}
