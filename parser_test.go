package libjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserAtoms(t *testing.T) {
	input := []string{
		"1",
		"1.0",
		"true",
		"false",
		"null",
		"12345",
		`"isastring"`,
	}
	wanted := []any{
		// i forgor 💀
		// (to test the lexing of single digit numbers, the eof handling was wrong)
		1.0,
		1.0,
		true,
		false,
		nil,
		12345.0,
		"isastring",
	}
	for i, in := range input {
		t.Run(in, func(t *testing.T) {
			in := []byte(in)
			p := &parser{l: lexer{data: in}}
			out, err := p.parse(in)
			assert.NoError(t, err)
			assert.EqualValues(t, wanted[i], out)
		})
	}
}

func TestParserArray(t *testing.T) {
	input := []string{
		"[]",
		"[1]",
		"[1, 2,3]",
		`["ayo", true, false, null, 12e12]`,
	}
	wanted := [][]any{
		{},
		{1.0},
		{1.0, 2.0, 3.0},
		{"ayo", true, false, nil, 12e12},
	}
	for i, in := range input {
		t.Run(in, func(t *testing.T) {
			in := []byte(in)
			p := &parser{l: lexer{data: in}}
			out, err := p.parse(in)
			assert.NoError(t, err)
			assert.EqualValues(t, wanted[i], out)
		})
	}
}

func TestParserObject(t *testing.T) {
	input := []string{
		"{}",
		`{ "key": "value" }`,
		`{ "key": 1 }`,
		`{ "key": { "key": 1 } }`,
		`{ "key": { "key": { "key": [1,2,3] } } }`,
		`{ "key1": "value1", "key2": "value2" }`,
	}
	wanted := []any{
		map[string]any{},
		map[string]any{"key": "value"},
		map[string]any{"key": 1.0},
		map[string]any{"key": map[string]any{"key": 1.0}},
		map[string]any{"key": map[string]any{"key": map[string]any{"key": []any{1.0, 2.0, 3.0}}}},
		map[string]any{"key1": "value1", "key2": "value2"},
	}
	for i, in := range input {
		t.Run(in, func(t *testing.T) {
			in := []byte(in)
			p := &parser{l: lexer{data: in}}
			out, err := p.parse(in)
			assert.NoError(t, err)
			assert.EqualValues(t, wanted[i], out)
		})
	}
}

func TestParserEdge(t *testing.T) {
	input := []string{
		`[]`,
		`{}`,
		`""`,
		"true",
		"null",
		`[{ "key": "value" }, {"key": "value"}, [1,2,3], null]`,
	}
	wanted := []any{
		[]any{},
		map[string]any{},
		"",
		true,
		nil,
		[]any{map[string]any{"key": "value"}, map[string]any{"key": "value"}, []any{1.0, 2.0, 3.0}, nil},
	}
	for i, in := range input {
		t.Run(in, func(t *testing.T) {
			in := []byte(in)
			p := &parser{l: lexer{data: in}}
			out, err := p.parse(in)
			assert.NoError(t, err)
			assert.EqualValues(t, wanted[i], out)
		})
	}
}

func TestParserFail(t *testing.T) {
	input := []string{
		"{",
		"]",
		"{ 1: 5 }",
		"{ ,,, }",
		"[,1]",
		`["": 1]`,
		"[1,\n1\n,1",
		"[{",
		"5 1 2 3",
		"true false null",
		"{} {}",
		`"str1" "str2"`,
		"[1,]",
		`{ "obj": {}, }`,
		`{ "obj": [, }`,
		// `{ "key": 1, "key": 2 }`, errors for duplicate keys was disabled, see parser.object()
		`{:"b"}`,
		`{"x"::"b"}`,
		"1.0e+",
		"0E",
		"1eE2",
	}
	for _, in := range input {
		t.Run(in, func(t *testing.T) {
			in := []byte(in)
			p := &parser{l: lexer{data: in}}
			out, err := p.parse(in)
			assert.Error(t, err)
			assert.Nil(t, out)
		})
	}
}
