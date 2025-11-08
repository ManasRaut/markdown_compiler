package parser

import (
	"testing"

	"github.com/ManasRaut/md_lex/ir"
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
				{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_PLAIN_TEXT, V: "This is heading 1"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"},
				{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"},
				{T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_LINE_BREAK, V: "\n"},
				{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_LINE_BREAK, V: "\n"},
			},
			want: []unFinishedElement{
				{Def: ir.HEADING_1_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "This is heading 1"}}},
				{Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
				{Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
				{Def: ir.CODE_BLOCK_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}}},
				{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}, {T: ir.TK_CODE_BLOCK, V: "```"}}},
				{Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
				{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "```"}}}, {Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
				{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 1');"}}}, {Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
				{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello Line 2');"}}}, {Def: ir.LINE_BREAK_DEFINITION, V: []ir.Token{}},
			},
			// TODO: add inline elements test cases
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
		// TODO: add inline elements test cases
		{
			name: "Base case",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_PLAIN_TEXT, V: "This is heading 1"}},
			wantElems: unFinishedElement{
				Def: ir.HEADING_1_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "List sequence",
			tkns: []ir.Token{{T: ir.TK_LIST_SEQUENCE, V: "1. "}, {T: ir.TK_PLAIN_TEXT, V: "This is a ordered list item"}},
			wantElems: unFinishedElement{
				Def:      ir.LIST_SEQUENCE_DEFINITION,
				Metadata: "1. ",
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This is a ordered list item"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with line break",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}, {T: ir.TK_PLAIN_TEXT, V: "This is heading 1"}, {T: ir.TK_LINE_BREAK, V: "\n"}},
			wantElems: unFinishedElement{
				Def: ir.HEADING_1_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This is heading 1"},
				},
			},
			wantI: 2,
		},
		{
			name: "Heading 1 with no contents",
			tkns: []ir.Token{{T: ir.TK_HEADING_1, V: "# "}},
			wantElems: unFinishedElement{
				Def: ir.HEADING_1_DEFINITION,
				V:   []ir.Token{},
			},
			wantI: 1,
		},
		{
			name: "Multiline Block",
			tkns: []ir.Token{{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello world');"}, {T: ir.TK_CODE_BLOCK, V: "```"}},
			wantElems: unFinishedElement{
				Def: ir.CODE_BLOCK_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "console.log('Hello world');"},
				},
			},
			wantI: 3,
		},
		{
			name: "Multi line line block element with error",
			tkns: []ir.Token{{T: ir.TK_CODE_BLOCK, V: "```"}, {T: ir.TK_PLAIN_TEXT, V: "console.log('Hello world');"}},
			wantElems: unFinishedElement{
				Def: ir.PLAIN_TEXT_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "```"},
				},
			},
			wantI: 1,
		},
		{
			name: "Normal text with mutiple and nested inline elements",
			tkns: []ir.Token{
				{T: ir.TK_PLAIN_TEXT, V: "This paragraph contains an "},
				{T: ir.TK_ITALIC, V: "*"},
				{T: ir.TK_PLAIN_TEXT, V: "italic "},
				{T: ir.TK_ITALIC, V: "*"},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_PLAIN_TEXT, V: "bold"},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_PLAIN_TEXT, V: " and "},
				{T: ir.TK_STRIKETHROUGH, V: "~~"},
				{T: ir.TK_PLAIN_TEXT, V: " strikethourgh text."},
				{T: ir.TK_STRIKETHROUGH, V: "~~"},
				{T: ir.TK_PLAIN_TEXT, V: "Also has "},
				{T: ir.TK_STRIKETHROUGH, V: "~~"},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_PLAIN_TEXT, V: " nested strikethought and bold "},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_STRIKETHROUGH, V: "~~"},
				{T: ir.TK_PLAIN_TEXT, V: "."},
			},
			wantElems: unFinishedElement{
				Def: ir.PLAIN_TEXT_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This paragraph contains an "},
					{T: ir.TK_ITALIC, V: "*"},
					{T: ir.TK_PLAIN_TEXT, V: "italic "},
					{T: ir.TK_ITALIC, V: "*"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: "bold"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: " and "},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: " strikethourgh text."},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: "Also has "},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: " nested strikethought and bold "},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: "."},
				},
			},
			wantI: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseAllBlockElements(0, tt.tkns, ir.ElementDefinitions[tt.tkns[0].T])
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
				Def: ir.LINE_BREAK_DEFINITION,
				V:   []ir.Token{},
			},
			wantI: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			g, i := parseSelfStandingElements(0, tt.tkns, ir.ElementDefinitions[tt.tkns[0].T])
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

func Test_parseInlineElements(t *testing.T) {
	tests := []struct {
		name string
		e    unFinishedElement
		want []*ir.MarkdownElement
	}{
		{
			name: "Self standing element",
			e:    unFinishedElement{Def: ir.HORIZONTAL_LINE_DEFINITION, V: []ir.Token{}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.HORIZONTAL_LINE_DEFINITION, "", nil)},
		},
		{
			name: "Image",
			e:    unFinishedElement{Def: ir.IMAGE_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: `![The San Juan Mountains are beautiful!](https://user-images.githubusercontent.com/9877795/143689169-e3386847-46ad-4747-9934-2293f3d39abd.png")`}}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.IMAGE_DEFINITION, `![The San Juan Mountains are beautiful!](https://user-images.githubusercontent.com/9877795/143689169-e3386847-46ad-4747-9934-2293f3d39abd.png")`, nil)},
		},
		{
			name: "Underline",
			e:    unFinishedElement{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{{T: ir.TK_UNDERLINE, V: "__"}, {T: ir.TK_PLAIN_TEXT, V: "Underline"}, {T: ir.TK_UNDERLINE, V: "__"}, {T: ir.TK_PLAIN_TEXT, V: " element"}}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "", []*ir.MarkdownElement{
					ir.NewMarkDownElement(ir.UNDERLINE_DEFINITION, "Underline", nil),
					ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, " element", nil),
				})},
		},
		{
			name: "CodeBlock",
			e:    unFinishedElement{Def: ir.CODE_BLOCK_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "This is a heading 1"}}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.CODE_BLOCK_DEFINITION, "This is a heading 1", nil)},
		},
		{
			name: "Heading 1 with normal text",
			e:    unFinishedElement{Def: ir.HEADING_1_DEFINITION, V: []ir.Token{{T: ir.TK_PLAIN_TEXT, V: "This is a heading 1"}}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.HEADING_1_DEFINITION, "", []*ir.MarkdownElement{ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This is a heading 1", nil)})},
		},
		{
			name: "Bold in normal text",
			e: unFinishedElement{Def: ir.PLAIN_TEXT_DEFINITION, V: []ir.Token{
				{T: ir.TK_PLAIN_TEXT, V: "This is a "},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_PLAIN_TEXT, V: "BOLD"},
				{T: ir.TK_BOLD, V: "**"},
				{T: ir.TK_PLAIN_TEXT, V: " text."}}},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "", []*ir.MarkdownElement{
					ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This is a ", nil),
					ir.NewMarkDownElement(ir.BOLD_DEFINITION, "BOLD", nil),
					ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, " text.", nil),
				})},
		},
		{
			name: "Bold and italic with no ending and combined and consecutive normal text",
			e: unFinishedElement{
				Def: ir.PLAIN_TEXT_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This text "},
					{T: ir.TK_PLAIN_TEXT, V: "contains a "},
					{T: ir.TK_BOLD_AND_ITALIC, V: "***"},
					{T: ir.TK_PLAIN_TEXT, V: "Italic and Bold "},
					{T: ir.TK_BOLD_AND_ITALIC, V: "***"},
					{T: ir.TK_PLAIN_TEXT, V: "text with "},
					{T: ir.TK_BOLD_AND_ITALIC, V: "***"},
					{T: ir.TK_PLAIN_TEXT, V: " no ending."},
				},
			},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(
					ir.PLAIN_TEXT_DEFINITION,
					"",
					[]*ir.MarkdownElement{
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This text contains a ", nil),
						ir.NewMarkDownElement(ir.BOLD_AND_ITALIC_DEFINITION, "Italic and Bold ", nil),
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "text with *** no ending.", nil),
					},
				),
			},
		},
		{
			name: "Unfinished emphasis text",
			e: unFinishedElement{
				Def: ir.HEADING_1_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This text "},
					{T: ir.TK_EMPHASIS, V: "`"},
					{T: ir.TK_PLAIN_TEXT, V: " incomplete emphasis text."},
				},
			},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(
					ir.HEADING_1_DEFINITION,
					"",
					[]*ir.MarkdownElement{
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This text ` incomplete emphasis text.", nil),
					},
				),
			},
		},
		{
			name: "Multiple nested inline elements combined",
			e: unFinishedElement{
				Def: ir.PLAIN_TEXT_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This paragraph contains an "},
					{T: ir.TK_ITALIC, V: "*"},
					{T: ir.TK_PLAIN_TEXT, V: "italic "},
					{T: ir.TK_ITALIC, V: "*"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: "bold"},
					{T: ir.TK_PLAIN_TEXT, V: " and bolder"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: " and "},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: " strikethourgh text."},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: "Also has "},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: "a "},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: "double "},
					{T: ir.TK_PLAIN_TEXT, V: "nested strikethought and bold "},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: "and again bold "},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: "."},
				},
			},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(
					ir.PLAIN_TEXT_DEFINITION,
					"",
					[]*ir.MarkdownElement{
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This paragraph contains an ", nil),
						ir.NewMarkDownElement(ir.ITALIC_DEFINITION, "italic ", nil),
						ir.NewMarkDownElement(ir.BOLD_DEFINITION, "bold and bolder", nil),
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, " and ", nil),
						ir.NewMarkDownElement(ir.STRIKETHROUGH_DEFINITION, " strikethourgh text.", nil),
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "Also has ", nil),
						ir.NewMarkDownElement(ir.STRIKETHROUGH_DEFINITION, "", []*ir.MarkdownElement{
							ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "a ", nil),
							ir.NewMarkDownElement(ir.BOLD_DEFINITION, "double nested strikethought and bold ", nil),
							ir.NewMarkDownElement(ir.BOLD_DEFINITION, "and again bold ", nil),
						}),
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, ".", nil),
					},
				),
			},
		},
		{
			name: "Striketgrough with no end but proper bold after it",
			e: unFinishedElement{
				Def: ir.BLOCK_QUOTE_DEFINITION,
				V: []ir.Token{
					{T: ir.TK_PLAIN_TEXT, V: "This paragraph contains a "},
					{T: ir.TK_STRIKETHROUGH, V: "~~"},
					{T: ir.TK_PLAIN_TEXT, V: " incomplete strikethourgh "},
					{T: ir.TK_PLAIN_TEXT, V: "but have a complete "},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: "bold"},
					{T: ir.TK_BOLD, V: "**"},
					{T: ir.TK_PLAIN_TEXT, V: " after it."},
				},
			},
			want: []*ir.MarkdownElement{
				ir.NewMarkDownElement(
					ir.BLOCK_QUOTE_DEFINITION,
					"",
					[]*ir.MarkdownElement{
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, "This paragraph contains a ~~ incomplete strikethourgh but have a complete ", nil),
						ir.NewMarkDownElement(ir.BOLD_DEFINITION, "bold", nil),
						ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, " after it.", nil),
					},
				),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseInlineElements(&tt.e)
			if !deepCompare(tt.want[0], got) {
				t.Fatalf("\nWant \n\t%v,\ngot \n\t%v", tt.want[0], got)
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
