package parser

import (
	"testing"

	"github.com/ManasRaut/lexe/ir"
)

func Test_parseBlockElements(t *testing.T) {
	tests := []struct {
		name string

		tkns []ir.Token
		want []BlockElementShell
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}},
			want: []BlockElementShell{
				{
					BlockElement: ir.BlockElement{T: ir.EL_HEADING_1},
					V: []ir.Token{
						{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
					},
				},
			},
		},
		{
			name: "Heading 1 with line break",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}, {T: ir.TK_LINE_BREAK, V: "\n"}},
			want: []BlockElementShell{
				{
					BlockElement: ir.BlockElement{T: ir.EL_HEADING_1},
					V: []ir.Token{
						{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
					},
				},
				{
					BlockElement: ir.BlockElement{T: ir.EL_LINE_BREAK},
					V:            []ir.Token{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got := parseBlockElements(tt.tkns)

			if len(tt.want) != len(got) {
				t.Fatalf("Has %d elements, want %d elements", len(got), len(tt.want))
			}

			for i := range len(tt.want) {
				w := tt.want[i]
				g := got[i]
				if w.T != g.T {
					t.Fatalf("at %d, %v, want %v", i, got[i].T, tt.want[i].T)
				}
				if len(w.V) != len(g.V) {
					t.Fatalf("at %d, tokens length is %v, want %v", i, len(g.V), len(w.V))
				}
				for j := range len(w.V) {
					wv := w.V[j]
					gv := g.V[j]
					if !wv.Equal(&gv) {
						t.Fatalf("at %d, token is %v, want %v", i, gv, wv)
					}
				}
			}
		})
	}
}

func Test_parseSingleLinkBlockElements(t *testing.T) {
	tests := []struct {
		name      string
		tkns      []ir.Token
		wantElems BlockElementShell
		wantI     int
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}},
			wantElems: BlockElementShell{
				BlockElement: ir.BlockElement{T: ir.EL_HEADING_1},
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with line break",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}, {T: ir.TK_LINE_BREAK, V: "\n"}},
			wantElems: BlockElementShell{
				BlockElement: ir.BlockElement{T: ir.EL_HEADING_1},
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with no contents",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}},
			wantElems: BlockElementShell{
				BlockElement: ir.BlockElement{T: ir.EL_HEADING_1},
				V:            []ir.Token{},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseSingleLinkBlockElements(0, len(tt.tkns), tt.tkns)
			w := tt.wantElems

			if w.T != g.T {
				t.Fatalf("got %v, want %v", g.T, w.T)
			}
			if len(w.V) != len(g.V) {
				t.Fatalf("Tokens length is %v, want %v", len(g.V), len(w.V))
			}
			if i != tt.wantI {
				t.Fatalf("Consumed tokens are %v, want %v", i, tt.wantI)
			}
			for j := range len(w.V) {
				wv := w.V[j]
				gv := g.V[j]
				if !wv.Equal(&gv) {
					t.Fatalf("at %d, token is %v, want %v", j, gv, wv)
				}
			}
		})
	}
}

func Test_parseSelfStandingElements(t *testing.T) {
	tests := []struct {
		name      string
		tkns      []ir.Token
		wantElems BlockElementShell
		wantI     int
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_LINE_BREAK, V: "\n"}},
			wantElems: BlockElementShell{
				BlockElement: ir.BlockElement{T: ir.EL_LINE_BREAK},
				V:            []ir.Token{},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseSelfStandingElements(0, tt.tkns)
			w := tt.wantElems

			if w.T != g.T {
				t.Fatalf("got %v, want %v", g.T, w.T)
			}
			if len(w.V) != len(g.V) {
				t.Fatalf("Tokens length is %v, want %v", len(g.V), len(w.V))
			}
			if i != tt.wantI {
				t.Fatalf("Consumed tokens are %v, want %v", i, tt.wantI)
			}
			for j := range len(w.V) {
				wv := w.V[j]
				gv := g.V[j]
				if !wv.Equal(&gv) {
					t.Fatalf("at %d, token is %v, want %v", j, gv, wv)
				}
			}
		})
	}
}

func Test_parseMultiLineBlockElements(t *testing.T) {
	tests := []struct {
		name      string
		tkns      []ir.Token
		wantElems BlockElementShell
		wantI     int
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello world');"}, {T: ir.TK_CODE_BLOCK, V: "```"}},
			wantElems: BlockElementShell{
				BlockElement: ir.BlockElement{T: ir.EL_CODE_BLOCK},
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello world');"},
				},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseMultiLineBlockElements(0, len(tt.tkns), tt.tkns)
			w := tt.wantElems

			if w.T != g.T {
				t.Fatalf("got %v, want %v", g.T, w.T)
			}
			if len(w.V) != len(g.V) {
				t.Fatalf("Tokens length is %v, want %v", len(g.V), len(w.V))
			}
			if i != tt.wantI {
				t.Fatalf("Consumed tokens are %v, want %v", i, tt.wantI)
			}
			for j := range len(w.V) {
				wv := w.V[j]
				gv := g.V[j]
				if !wv.Equal(&gv) {
					t.Fatalf("at %d, token is %v, want %v", j, gv, wv)
				}
			}
		})
	}
}
