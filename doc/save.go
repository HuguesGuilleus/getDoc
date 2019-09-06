package doc

import (
	"./data"
	"fmt"
	"os"
)

// Save the index in a a file in html
func (ind *Index) SaveHTML(path string) {
	blobInfo, _ := os.Stat(path)
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
