// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package doc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/HuguesGuilleus/getDoc/web"
	"log"
	"os"
	"path/filepath"
)

// Save the index in a a file in html
func (ind *Index) SaveHTML(path string) {
	defer rec()
	ind.sort()
	file := writeFile(path, "doc.html")
	defer file.Close()
	panicing(webdata.Index.Execute(file, ind))
	log.Print("SAVE IN HTML: ", file.Name())
}

// The index for XML and JSON encoding
type DataIndex struct {
	List  Index
	Files []string
	Date  string
}

// Cast Index to DataIndex
func (ind *Index) DataIndex() *DataIndex {
	ind.sort()
	return &DataIndex{
		List:  *ind,
		Files: ind.ListFile(),
		Date:  ind.Date(),
	}
}

// Save the data in a file in JSON encoding
func (ind *DataIndex) Json(path string) {
	defer rec()
	data, e := json.MarshalIndent(*ind, "", "\t")
	panicing(e)
	file := writeFile(path, "doc.json")
	defer file.Close()
	_, e = file.Write(data)
	panicing(e)
	log.Print("SAVE IN JSON: ", file.Name())
}

// Save the data in a file in JSON encoding
func (ind *DataIndex) Xml(path string) {
	defer rec()
	data, e := xml.MarshalIndent(*ind, "", "\t")
	panicing(e)
	file := writeFile(path, "doc.xml")
	defer file.Close()
	_, e = file.Write(data)
	panicing(e)
	log.Print("SAVE IN XML: ", file.Name())
}

// Open a file for writing
// If path is a directory, the file is path+name else if just path
// It panic with error
func writeFile(path, name string) *os.File {
	if len(name) == 0 {
		panic("writeFile: you must give non empty string for name")
	}
	if len(path) == 0 {
		path = name
	} else {
		blobInfo, err := os.Stat(path)
		if err == nil && blobInfo.IsDir() {
			if path[len(path)-1] == '/' {
				path += name
			} else {
				path += "/" + name
			}
		}
	}
	path = filepath.Clean(path)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	panicing(err)
	return file
}
