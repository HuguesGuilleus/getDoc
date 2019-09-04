package main

import (
	"./until"
	"./data"
	"os"
)

func main() {
	index := until.Index{}
	index.AddFile("dataTest/main.c")

	// Écriture de index.html
	file,err := os.OpenFile("doc.html", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data.Index.Execute(file, index)

	// ioutil.WriteFile("index.html", )
}
