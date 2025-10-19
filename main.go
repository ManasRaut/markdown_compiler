package main

import (
	_ "embed"
	"fmt"

	"github.com/ManasRaut/lexe/ir"
	"github.com/ManasRaut/lexe/lexer"
	"github.com/ManasRaut/lexe/parser"
)

//go:embed markdown_elements.md
var exmapleSource string

//go_:embed markdown_elements2.md

func main() {

	lex := lexer.InitLexer(exmapleSource)
	tokens := lex.Parse()

	printTokens(tokens)
	fmt.Println("\n============================\n")

	psr := parser.NewParser(tokens)

	fmt.Println(psr.Parse())
}

func printTokens(tokens []ir.Token) {
	for _, token := range tokens {
		if token.T == ir.TK_LINE_BREAK {
			fmt.Println()
		} else {
			fmt.Printf("(%s:::%s) ", token.String(), token.V)
		}
	}
}
