package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

var (
	logOutput    *os.File = nil
	commentShort          = regexp.MustCompile("//.*?\\n")
	commentLong           = regexp.MustCompile("/\\*[^\\0]*?\\*/")
	spaceSimple           = regexp.MustCompile("\\s+")
	spaceJs               = regexp.MustCompile(" ?(\\W) ?")
	spaceCss              = regexp.MustCompile(" ?([:;{}]) ?")

	targetCss   = regexp.MustCompile("{CSS}")
	targetJs    = regexp.MustCompile("{JS}")
	targetIndex = regexp.MustCompile("{{INDEX}}")
)

func init() {
	// var err error
	// logOutput, err = os.OpenFile("build.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(logOutput)
}

func main() {
	defer logOutput.Close()

	var index, css, js []byte

	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		index = read("doc/data/index.gohtml")
	}()
	go walk("style/", ".css", &css, wg)
	go walk("search/", ".js", &js, wg)
	wg.Wait()

	index = targetCss.ReplaceAll(index, css)
	index = targetJs.ReplaceAll(index, js)

	file := []byte("package data\nvar index = `{{INDEX}}`")
	file = targetIndex.ReplaceAll(file, index)
	ioutil.WriteFile("doc/data/index.go", file, 0664)
}

func walk(dir, ext string, data *[]byte, wg *sync.WaitGroup) {
	defer wg.Done()
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Error:", err)
		}
		if info.IsDir() || filepath.Ext(info.Name()) != ext {
			return nil
		}
		file := read(path)
		switch ext {
		case ".js":
			file = simpleJs(file)
		case ".css":
			file = simpleCSS(file)
		}
		*data = append(*data, file...)
		return nil
	})
}

// Read a file, if err, log.Fatal
func read(path string) (file []byte) {
	log.Println("Read file:", path)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func simpleJs(data []byte) []byte {
	data = commentShort.ReplaceAll(data, []byte(" "))
	data = commentLong.ReplaceAll(data, []byte(" "))
	data = spaceSimple.ReplaceAll(data, []byte(" "))
	data = spaceJs.ReplaceAll(data, []byte("$1"))
	return data
}

func simpleCSS(data []byte) []byte {
	data = commentLong.ReplaceAll(data, []byte(" "))
	data = spaceSimple.ReplaceAll(data, []byte(" "))
	data = spaceCss.ReplaceAll(data, []byte("$1"))
	return data
}
