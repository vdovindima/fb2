// Package fb2 represent .fb2 format parser
package fb2

import (
	"encoding/xml"
)

// Parser struct
type Parser struct {
	book []byte
}

// New creates new Parser
func New(data []byte) *Parser {
	return &Parser{
		book: data,
	}
}

// Unmarshal parse data to FB2 type
func (p *Parser) Unmarshal() (result FB2, err error) {
	if err = xml.Unmarshal(p.book, &result); err != nil {
		return
	}

	result.UnmarshalCoverpage(p.book)

	return
}
