package xmlbench

import (
	"testing"
	"xmlbench/file"
)

func BenchmarkParsers(b *testing.B) {
	benchmarks := []struct {
		xmlFilePath string
	}{
		{"hotels.xml"},
		{"flights.xml"},
	}
	for _, bm := range benchmarks {
		// Read file contents
		fileContents, err := file.GetFileContents(bm.xmlFilePath)
		if err != nil {
			b.Fatal(err)
		}

		b.Run(bm.xmlFilePath+"-encoding/xml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser1(fileContents); err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(bm.xmlFilePath+"-cvik/xml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser2(fileContents); err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(bm.xmlFilePath+"-libxml2", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser3(fileContents); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
