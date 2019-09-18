// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var Title string = ""

// Read a file or a directory
func Read(root string) (ind *Index) {
	ind = &Index{}
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)
	go getTitle(root, wg)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		defer rec()
		panicing(err)
		if !info.IsDir() {
			wg.Add(1)
			go ind.readFile(path, wg)
		}
		return nil
	})
	return ind
}

// Read one file, test if the type is know, split the line and call the parser
func (ind *Index) readFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	if parser := langKnown(path); parser != nil {
		log.Print("READ FILE: ", path)
		lines := splitFile(path)
		parser.Type(lines)
		parser.Parse(ind, lines, path)
	}
}

// Read files in debug mode
func ReadDebug(roots []string) {
	for _, root := range roots {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			defer rec()
			panicing(err)
			if !info.IsDir() {
				if parser := langKnown(path); parser != nil {
					fmt.Println("READ FILE: ", path)
					lines := splitFile(path)
					parser.Type(lines)
					for i, l := range lines {
						t := nameType[l.Type]
						fmt.Printf("%3d %6s :: %s\n", i, t, l.Str)
					}
				}
			}
			return nil
		})
	}
}

// Get the name of the file or directory
func getTitle(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	path, err := filepath.Abs(path)
	printErr(err)
	Title = filepath.Base(path)
}

// recover and print the error
func rec() {
	err := recover()
	printErr(err)
}

// Panic if error is not nil
func panicing(err error) {
	if err != nil {
		panic(err)
	}
}

// Print a error in os.Stderr with red
func printErr(err interface{}) {
	if err != nil {
		// Output
		oldWriter := log.Writer()
		defer log.SetOutput(oldWriter)
		log.SetOutput(os.Stderr)
		// Prefix
		oldPrefix := log.Prefix()
		defer log.SetPrefix(oldPrefix)
		log.SetPrefix("\033[01;31m=== ")
		// print
		log.Print("ERROR: ", err, "\033[0m")
	}
}
