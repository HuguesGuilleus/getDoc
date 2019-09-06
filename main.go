package main

import (
	"./doc"
)

func main() {
	ind := doc.Read("dataTest/main.c")
	ind.SaveHTML("./doc.html")
}
