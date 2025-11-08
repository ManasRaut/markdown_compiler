package tests_test

import (
	"testing"

	"github.com/ManasRaut/md_lex/ir"
	"github.com/ManasRaut/md_lex/parser"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name         string
		tokens       []ir.Token
		wantErr      error
		wantElements []*ir.MarkdownElement
	}{
		{
			name:         "sample_markdown",
			tokens:       loadTokens(readOutputFile("./sample_markdown.tokens")),
			wantElements: loadElements(readOutputFile("./sample_markdown.elements")),
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			psr := parser.NewParser(tt.tokens)
			gotErr := psr.Parse()
			if gotErr != tt.wantErr {
				t.Fatalf("parser.Parser(): Got err %v, want %v", gotErr, tt.wantErr)
			}

			gotElems := psr.GetElements()
			wantElems := tt.wantElements

			for i := range wantElems {
				got := gotElems[i]
				want := wantElems[i]
				if !deepCompare(got, want) {
					t.Fatalf("parser.Parse(): \nGot Elements:%v\nWant Elements:%v\n\nWhere got:\n\t%v\nwant:\n\t%v\n", got, want, gotElems, wantElems)
				}
			}
			if len(gotElems) != len(wantElems) {
				t.Fatalf("parser.Parse(): Got %d elements at root, want %d elements at root", len(gotElems), len(wantElems))
			}
		})
	}
}
