package xmlbench

import (
	"bytes"

	"github.com/lestrrat/go-libxml2/parser"
)

// Parser3 decodes XML with libxml2 package
func Parser3(xmlStr []byte) error {
	p := parser.New()
	doc, err := p.ParseReader(bytes.NewReader(xmlStr))
	if err != nil {
		return err
	}
	defer doc.Free()
	return nil
}
