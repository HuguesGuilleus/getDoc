package main

import (
	"./doc"
	"log"
)

func main() {
	log.SetPrefix("==== ")
	log.SetFlags(0)
	ind := doc.Read("dataTest/")
	ind.SaveHTML("./doc.html")
}
