// getDoc
// 2019 GUILLEUS Hugues <ghugues@netc.fr>
// BSD 3-Clause "New" or "Revised" License

package main

import (
	"log"
)

// A io.Writer who d'ont have effect
type emptyWriter struct{}

func (*emptyWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func init() {
	log.SetPrefix("==== ")
	log.SetFlags(0)
	log.SetOutput(&emptyWriter{})
}
