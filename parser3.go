package xmlbench

import (
	"xmlbench/file"

	"github.com/lestrrat/go-libxml2/parser"
)

func Parser3(filePath string) error {
	// Get file reader
	fileReader, err := file.GetFileReader(filePath)
	if err != nil {
		return err
	}
	// Parse XML
	p := parser.New()
	doc, err := p.ParseReader(fileReader)
	if err != nil {
		return err
	}
	defer doc.Free()
	return nil
}
