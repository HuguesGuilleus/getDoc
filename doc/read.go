package doc

import (
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
		if info.IsDir() {
			return nil
		} else {
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
	if parser := langKnown(getExt(path)); parser != nil {
		log.Print("READ FILE: ", path)
		parser(ind, splitFile(path), path)
	}
}

// Get the name of the file or directory
func getTitle(path string, wg *sync.WaitGroup)  {
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
		log.SetPrefix("\033[01;31m" + oldPrefix)
		// print
		log.Print("ERROR: ", err, "\033[0m")
	}
}
