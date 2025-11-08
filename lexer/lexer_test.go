package lexer_test

import (
	"testing"

	"github.com/ManasRaut/md_lex/ir"
	"github.com/ManasRaut/md_lex/lexer"
)

func TestLexer(t *testing.T) {
	tests := []struct {
		name       string
		source     string
		wantTokens []ir.Token
	}{
		{
			name:   "Heading 1",
			source: "# Heading 1",
			wantTokens: []ir.Token{
				{T: ir.TK_HEADING_1, V: "# "},
				{T: ir.TK_NORMAL_TEXT, V: "Heading 1"},
			},
		},
		{
			name:   "Underline in middle",
			source: "A __Underline__ in a line",
			wantTokens: []ir.Token{
				{T: ir.TK_NORMAL_TEXT, V: "A "},
				{T: ir.TK_UNDERLINE, V: "__"},
				{T: ir.TK_NORMAL_TEXT, V: "Underline"},
				{T: ir.TK_UNDERLINE, V: "__"},
				{T: ir.TK_NORMAL_TEXT, V: " in a line"},
			},
		},
		{
			name:   "Underline at start",
			source: "__Underline__ in a line and a rouge _ underscore",
			wantTokens: []ir.Token{
				{T: ir.TK_UNDERLINE, V: "__"},
				{T: ir.TK_NORMAL_TEXT, V: "Underline"},
				{T: ir.TK_UNDERLINE, V: "__"},
				{T: ir.TK_NORMAL_TEXT, V: " in a line and a rouge _ underscore"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lex := lexer.InitLexer(tt.source)
			gotTokens := lex.Parse()

			for i := range tt.wantTokens {
				got := gotTokens[i]
				want := tt.wantTokens[i]
				if got != want {
					t.Fatalf("Lexer.Parse(): Got token: %v, want token %v", got, want)
				}
			}
		})
	}
}
