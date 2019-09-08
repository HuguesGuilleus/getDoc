package doc

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var Title string = ""

// Read a file or a directory
func Read(path string) (ind *Index) {
	ind = &Index{}
	blobInfo, err := os.Stat(path)
	if err != nil {
		printErr(err)
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	if blobInfo.IsDir() {
		go ind.readDir(path, wg)
		t, err := filepath.Abs(filepath.Dir(path))
		printErr(err)
		Title = filepath.Base(t)
	} else {
		go ind.readFile(path, wg)
		Title = blobInfo.Name()
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
			go ind.readDir(path+file.Name(), wg)
		} else {
			go ind.readFile(path+file.Name(), wg)
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
