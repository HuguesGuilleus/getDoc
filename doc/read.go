package doc

import (
	"log"
	"os"
)

func Read(path string) (ind *Index) {
	ind = &Index{}
	blobInfo, _ := os.Stat(path)
	if blobInfo.IsDir() {
		// ind.readDir()
	} else {
		ind.readFile(path)
	}
	return
}

func (ind *Index) readFile(path string) {
	if parser := langKnown(getExt(path)); parser != nil {
		log.Print("READ FILE:",path)
		lines := splitFile(path)
		parser(ind, lines, path)
	}
}
