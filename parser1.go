package xmlbench

import (
	"ag/hotels"
	"encoding/xml"
	"xmlbench/file"
)

// Parser1 decodes XML with encoding/xml package
func Parser1(filePath string) error {
	// Get file reader
	fileReader, err := file.GetFileReader(filePath)
	if err != nil {
		return err
	}
	// Parse XML
	searchResponse := &hotels.SearchResponse{}
	decoder := xml.NewDecoder(fileReader)
	if err := decoder.Decode(searchResponse); err != nil {
		return err
	}
	return nil
}
