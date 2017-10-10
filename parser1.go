package xmlbench

import (
	"bytes"
	"encoding/xml"
)

// Parser1 decodes XML with encoding/xml package
func Parser1(xmlStr []byte) error {
	v := SearchResponse{}
	decoder := xml.NewDecoder(bytes.NewReader(xmlStr))
	if err := decoder.Decode(&v); err != nil {
		return err
	}
	return nil
}
