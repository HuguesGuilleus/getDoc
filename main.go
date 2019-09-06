package main

import (
	"./doc"
	"log"
	"time"
)

func main() {
	start := time.Now()
	defer log.Print("TIME: ",time.Since(start))
	log.SetPrefix("==== ")
	log.SetFlags(0)
	ind := doc.Read("dataTest/")
	ind.SaveHTML("./doc.html")
}
