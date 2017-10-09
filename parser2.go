package xmlbench

import (
	"bytes"

	"github.com/cvik/xml"
)

// Parser2 decodes XML with cvik/xml package
func Parser2(xmlStr []byte) error {
	// Parse XML
	doc, err := xml.NewParser(bytes.NewReader(xmlStr)).ParseXML()
	if err != nil {
		return err
	}
	_ = doc
	return nil
}
