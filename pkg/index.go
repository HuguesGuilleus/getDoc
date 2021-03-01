// getDoc
// 2021 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
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
	Log log.Logger
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
