package xmlbench

import (
	"xmlbench/file"

	"github.com/cvik/xml"
)

// Parser1 decodes XML with encoding/xml package
func Parser2(filePath string) error {
	// Get file reader
	fileReader, err := file.GetFileReader(filePath)
	if err != nil {
		return err
	}
	// Parse XML
	doc, err := xml.NewParser(fileReader).ParseXML()
	if err != nil {
		return err
	}
	_ = doc
	return nil
}
