package xmlbench

import (
	"testing"
)

func BenchmarkParser1(b *testing.B) {
	benchmarks := []struct {
		xmlFilePath string
	}{
		{"1.xml"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.xmlFilePath+"-encoding/xml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser1(bm.xmlFilePath); err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(bm.xmlFilePath+"-cvik/xml", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser2(bm.xmlFilePath); err != nil {
					b.Fatal(err)
				}
			}
		})
		b.Run(bm.xmlFilePath+"-libxml2", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if err := Parser3(bm.xmlFilePath); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
