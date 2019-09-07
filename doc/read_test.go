package doc

import (
	"testing"
)

func BenchmarkReadAndSave(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ind := Read("../")
		ind.SaveHTML("../doc.html")
		ind.DataIndex().Json("doc.json")
		ind.DataIndex().Xml("doc.xml")
	}
}
