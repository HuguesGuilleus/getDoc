package doc

import (
	"./data"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Save the index in a a file in html
func (ind *Index) SaveHTML(path string) {
	blobInfo, err := os.Stat(path)
	if err != nil {
		printErr(err)
		return
	}
	if blobInfo.IsDir() {
		ind.saveHTMLinDir(path)
	} else {
		ind.saveHTMLinFile(path)
	}
}

// Save the doc in path(who are a directory)
func (ind *Index) saveHTMLinDir(path string) {
	if path[len(path)-1] == '/' {
		ind.saveHTMLinFile(path + "doc.html")
	} else {
		ind.saveHTMLinFile(path + "/doc.html")
	}
}

// Save the doc in path
func (ind *Index) saveHTMLinFile(path string) {
	if ind == nil {
		fmt.Fprintln(os.Stderr, "The index is null")
		return
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()
	data.Index.Execute(file, *ind)
}

// The index for XML and JSON encoding
type DataIndex struct {
	List Index
	Files []string
	Date string
}

// Cast Index to DataIndex
func (ind *Index) DataIndex() *DataIndex {
	return &DataIndex{
		List : *ind,
		Files: ind.ListFile(),
		Date: ind.Date(),
	}
}

// Save the data in a file in JSON encoding
// path must be a file not a directory
func (ind *DataIndex) Json(path string) (err bool) {
	data, e := json.Marshal(*ind)
	if e != nil {
		printErr(e)
		return true
	}
	e = ioutil.WriteFile(path, data, 0664)
	if e != nil {
		printErr(e)
		return true
	}
	log.Print("SAVED IN JSON: ",path)
	return false
}
