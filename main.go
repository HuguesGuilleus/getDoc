package main

import (
	"./doc"
	"log"
	"os"
)

func main() {
	log.SetPrefix("==== ")
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	ind := doc.Read("./")
	ind.SaveHTML("./doc.html")
	ind.DataIndex().Json("doc.json")
	ind.DataIndex().Xml("doc.xml")
}
