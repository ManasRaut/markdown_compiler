package parser

import (
	"testing"

	"github.com/ManasRaut/lexe/ir"
)

func Test_parseBlockElements(t *testing.T) {
	tests := []struct {
		name string

		tkns []ir.Token
		want []unFinishedElement
	}{
		{
			name: "Base Case",
			tkns: []ir.Token{
				{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"},
				{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_LINE_BREAK, V: "\n"},
			},
			want: []unFinishedElement{
				{Def: elementDefinitions[ir.TK_HEADING_1], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}}},
				{Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
				{Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
				{Def: elementDefinitions[ir.TK_CODE_BLOCK], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}}},
				{Def: elementDefinitions[ir.TK_NORMAL_TEXT], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"}}},
				{Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
				{Def: elementDefinitions[ir.TK_NORMAL_TEXT], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "```"}}}, {Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
				{Def: elementDefinitions[ir.TK_NORMAL_TEXT], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 1');"}}}, {Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
				{Def: elementDefinitions[ir.TK_NORMAL_TEXT], V: []ir.Token{{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello Line 2');"}}}, {Def: elementDefinitions[ir.TK_LINE_BREAK], V: []ir.Token{}},
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
				if w.Def != g.Def {
					t.Fatalf("Def at %d, is %v, want %v", i, g.Def, w.Def)
				}
				if len(w.V) != len(g.V) {
					t.Fatalf("at %d, tokens length is %v, want %v", i, len(g.V), len(w.V))
				}
				for j := range len(w.V) {
					wv := w.V[j]
					gv := g.V[j]
					if !wv.Equal(&gv) {
						t.Fatalf("%d::%d, token is %v, want %v", i, j, gv, wv)
					}
				}
			}
		})
	}
}

func Test_parseAllBlockElements(t *testing.T) {
	tests := []struct {
		name      string
		tkns      []ir.Token
		wantElems unFinishedElement
		wantI     int
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_HEADING_1],
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with line break",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_NORMAL_TEXT, V: "This is heading 1"}, {T: ir.TK_LINE_BREAK, V: "\n"}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_HEADING_1],
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with no contents",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_HEADING_1],
				V:   []ir.Token{},
			},
			wantI: 1,
		},
		{
			name: "Multiline Block",
			tkns: []ir.Token{{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello world');"}, {T: ir.TK_CODE_BLOCK, V: "```"}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_CODE_BLOCK],
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "console.log('Hello world');"},
				},
			},
			wantI: 3,
		},
		{
			name: "Multi line line block element with error",
			tkns: []ir.Token{{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_NORMAL_TEXT, V: "console.log('Hello world');"}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_NORMAL_TEXT],
				V: []ir.Token{
					{T: ir.TK_NORMAL_TEXT, V: "```"},
				},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseAllBlockElements(0, tt.tkns, elementDefinitions[tt.tkns[0].T])
			w := tt.wantElems

			if w.Def != g.Def {
				t.Fatalf("Def got %v, want %v", g.Def, w.Def)
			}
			if len(w.V) != len(g.V) {
				t.Fatalf("Tokens length is %v, want %v", len(g.V), len(w.V))
			}
			if i != tt.wantI {
				t.Fatalf("New index at %v, want %v", i, tt.wantI)
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
		wantElems unFinishedElement
		wantI     int
	}{
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_LINE_BREAK, V: "\n"}},
			wantElems: unFinishedElement{
				Def: elementDefinitions[ir.TK_LINE_BREAK],
				V:   []ir.Token{},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseSelfStandingElements(0, tt.tkns, elementDefinitions[tt.tkns[0].T])
			w := tt.wantElems

			if w.Def != g.Def {
				t.Fatalf("got %v, want %v", g.Def, w.Def)
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
