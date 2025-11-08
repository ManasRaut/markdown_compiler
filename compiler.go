package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/ManasRaut/md_lex/converters"
	"github.com/ManasRaut/md_lex/lexer"
	"github.com/ManasRaut/md_lex/parser"
)

type MDLexCompiler[R any] struct {
	converter converters.Converter[R]
}

func (c *MDLexCompiler[R]) Compile(r io.Reader) (*R, error) {
	sourceCode, err := readSourceCode(r)
	if err != nil {
		return nil, err
	}

	lexer := lexer.InitLexer(sourceCode)
	tokens := lexer.Parse()

	parser := parser.NewParser(tokens)
	err = parser.Parse()
	if err != nil {
		return nil, err
	}
	mdElements := parser.GetElements()

	uiElements, err := c.converter.Convert(mdElements)

	return uiElements, nil
}

func readSourceCode(r io.Reader) (string, error) {
	sourceCode := strings.Builder{}

	for {
		buf := make([]byte, 1024)
		n, err := r.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}
			return "", fmt.Errorf("failed to read source code, caused by: %e", err)
		}

		for i := 0; i < n; i++ {
			sourceCode.WriteByte(buf[i])
		}
	}

	return sourceCode.String(), nil
}

func NewMDLexCompiler[R any](converter converters.Converter[R]) (*MDLexCompiler[R], error) {
	return &MDLexCompiler[R]{
		converter: converter,
	}, nil
}
