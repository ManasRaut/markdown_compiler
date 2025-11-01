package tests_test

import (
	"testing"

	"github.com/ManasRaut/lexe/ir"
	"github.com/ManasRaut/lexe/lexer"
)

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
