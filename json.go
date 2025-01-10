package libjson

import (
	"io"
)

func NewReader(r io.Reader) (JSON, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return JSON{}, err
	}
	p := parser{l: lexer{data: data}}
	obj, err := p.parse(data)
	if err != nil {
		return JSON{}, err
	}
	return JSON{obj}, nil
}

func New(data []byte) (JSON, error) {
	p := parser{l: lexer{data: data}}
	obj, err := p.parse(data)
	if err != nil {
		return JSON{}, err
	}
	return JSON{obj}, nil
}
