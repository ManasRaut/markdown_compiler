package tests_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/ManasRaut/lexe/ir"
	"github.com/ManasRaut/lexe/lexer"
)

const TOKEN_SEPERATOR = "<<%>>"

func readOutputFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := strings.Builder{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text.WriteString(scanner.Text())
		text.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return text.String()
}

func loadTokens(str string) []ir.Token {

	tokens := make([]ir.Token, 0)

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			tokens = append(tokens, ir.Token{T: ir.TK_LINE_BREAK, V: "\n"})
			continue
		}
		unparsedTokens := strings.Split(line, TOKEN_SEPERATOR)
		for i, ununparsedToken := range unparsedTokens {
			name, value, found := strings.Cut(ununparsedToken, " ")
			if !found {
				panic(fmt.Sprintf("At line %d, Space seperator expected in unparsedToken %s", i, ununparsedToken))
			}
			tokenType := ir.GetTokentype(name)
			if tokenType == ir.TK_UNKNOWN {
				panic(fmt.Sprintf("At line %d, Unknown token type found in expected output %s", i, name))
			}
			tokens = append(tokens, ir.Token{T: tokenType, V: value})
		}
		tokens = append(tokens, ir.Token{T: ir.TK_LINE_BREAK, V: "\n"})
	}

	return tokens
}

func TestLexer(t *testing.T) {
	tests := []struct {
		name       string
		sourceCode string
		wantTokens []ir.Token
	}{
		{
			name:       "sample_markdown",
			sourceCode: readOutputFile("./sample_markdown.md"),
			wantTokens: loadTokens(readOutputFile("./sample_markdown.tokens")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.InitLexer(tt.sourceCode)
			gotTokens := lex.Parse()
			if len(gotTokens) != len(tt.wantTokens) {
				t.Fatalf("lexer.Parse(): Got %d tokens, want %d tokens\ngot:\n\t%s\nwant:\n\t%s", len(gotTokens), len(tt.wantTokens), gotTokens, tt.wantTokens)
			}
			for i := range tt.wantTokens {
				got := gotTokens[i]
				want := tt.wantTokens[i]
				if got != want {
					t.Fatalf("lexer.Parse(): \nGot Token:%v\nWant Token:%v\n\nWhere got:\n\t%v\nwant:\n\t%v\n", got, want, gotTokens, tt.wantTokens)
				}
			}
		})
	}
}

func deepCompare(w *ir.MarkdownElement, g *ir.MarkdownElement) bool {

	if w == nil && g == nil {
		return true
	}
	if (w == nil && g != nil) || (w != nil && g == nil) || (w.V != g.V) || len(w.C) != len(g.C) {
		return false
	}
	if w.C == nil && g.C == nil {
		return true
	}
	if (w.C == nil && g.C != nil) || (w.C != nil && g.C == nil) {
		return false
	}
	for i := range w.C {
		a := w.C[i]
		b := g.C[i]
		if !deepCompare(a, b) {
			return false
		}
	}
	return true
}
