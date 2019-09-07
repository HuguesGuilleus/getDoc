package doc

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Read a file or a directory
func Read(path string) (ind *Index) {
	if ind == nil {
		ind = &Index{}
	}
	blobInfo, err := os.Stat(path)
	if err != nil {
		printErr(err)
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	if blobInfo.IsDir() {
		ind.readDir(path, wg)
	} else {
		ind.readFile(path, wg)
	}
	wg.Wait()
	return
}

// read a dir and for each file call readDir or readFile
func (ind *Index) readDir(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	if path[len(path)-1] != '/' {
		path += "/"
	}
	log.Print("READ DIRECTORY: ", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Print("Error", err)
		return
	}
	for _, file := range files {
		wg.Add(1)
		if file.IsDir() {
			ind.readDir(path+file.Name(), wg)
		} else {
			ind.readFile(path+file.Name(), wg)
		}
	}
}

// Read one file, test if the type is know, split the line and call the parser
func (ind *Index) readFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer rec()
	if parser := langKnown(getExt(path)); parser != nil {
		log.Print("READ FILE: ", path)
		lines := splitFile(path)
		parser(ind, lines, path)
	}
}

// recover and print the error
func rec() {
	err := recover()
	printErr(err)
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
