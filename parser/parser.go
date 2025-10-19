package parser

import (
	"slices"

	"github.com/ManasRaut/lexe/ir"
)

var SINGLE_LINE_BLOCK_ELEMENTS []ir.TokenType = []ir.TokenType{
	ir.TK_HEADING_1,
	ir.TK_HEADING_2,
	ir.TK_HEADING_3,
	ir.TK_HEADING_4,
	ir.TK_HEADING_5,
	ir.TK_HEADING_6,

	// ir.TK_HORIZONTAL_LINE,
	// ir.TK_BLOCK_QUOTE,
	ir.TK_BULLET_POINT,
	ir.TK_LIST_SEQUENCE,
	// ir.TK_CODE_BLOCK,
	// ir.TK_IMAGE,
	// ir.TK_LINE_BREAK,
	// ir.TK_CHECKED_BOX,
	// ir.TK_UNCHECKED_BOX,
	// ir.TK_BOLD_AND_ITALIC,
	// ir.TK_BOLD,
	// ir.TK_ITALIC,
	// ir.TK_EMPHASIS,
	// ir.TK_ESCAPE_CHARACTER,
	// ir.TK_STRIKETHROUGH,
	ir.TK_NORMAL_TEXT,
}

var SELF_STANDING_ELEMENTS []ir.TokenType = []ir.TokenType{
	ir.TK_LINE_BREAK,
	ir.TK_HORIZONTAL_LINE,
	ir.TK_IMAGE,
}

var MULTI_LINE_ELEMENTS []ir.TokenType = []ir.TokenType{
	ir.TK_CODE_BLOCK,
	ir.TK_BLOCK_QUOTE,
}

// TODO: replace this with ElementIRRepresent containing info as : type, block or inline or multiline, can contain inline, etc
type BlockElementShell struct {
	ir.BlockElement
	V []ir.Token
}

type Parser struct {
	tks  []ir.Token
	elms []ir.BlockElement
}

func (p *Parser) Parse() error {
	blkElems := parseBlockElements(p.tks)
	p.elms = parseInlineElements(blkElems)
	return nil
}

func (p *Parser) GetElements() []ir.BlockElement {
	return p.elms
}

func parseBlockElements(tkns []ir.Token) []BlockElementShell {

	r := make([]BlockElementShell, 0, len(tkns))
	i := 0
	l := len(tkns)

	for {

		if i >= l {
			break
		}

		t := tkns[i]
		// Default is normal text
		el := BlockElementShell{
			BlockElement: ir.BlockElement{
				T: ir.EL_NORMAL_TEXT,
				V: []ir.InlineElement{
					{T: ir.EL_NORMAL_TEXT, V: t.V},
				},
			},
		}

		if slices.Contains(SINGLE_LINE_BLOCK_ELEMENTS, t.T) {
			el, i = parseSingleLinkBlockElements(i, l, tkns)
		} else if slices.Contains(SELF_STANDING_ELEMENTS, t.T) {
			el, i = parseSelfStandingElements(i, tkns)
		} else if slices.Contains(MULTI_LINE_ELEMENTS, t.T) {
			el, i = parseMultiLineBlockElements(i, l, tkns)
		}
		r = append(r, el)
	}

	return r
}

func parseSingleLinkBlockElements(i int, l int, tkns []ir.Token) (BlockElementShell, int) {
	t := tkns[i]
	j := i
	for {
		if j >= l || tkns[j].T == ir.TK_LINE_BREAK {
			j--
			break
		}
		j++
	}
	j++

	b := BlockElementShell{
		BlockElement: ir.BlockElement{
			T: getTokensBlockElement(t.T),
		},
		V: tkns[i+1 : j],
	}
	return b, j
}

func parseSelfStandingElements(i int, tkns []ir.Token) (BlockElementShell, int) {
	t := tkns[i]
	b := BlockElementShell{
		BlockElement: ir.BlockElement{
			T: getTokensBlockElement(t.T),
		},
		V: []ir.Token{},
	}
	return b, i + 1
}

func parseMultiLineBlockElements(i int, l int, tkns []ir.Token) (BlockElementShell, int) {
	/*
		create an multi line block element shell
		iterate till end of the multiline block is reached
			if  end not found treat all of this as NORMAL_TEXT
			if multiline block element con contain INLINE_ELEMENTS
				give inline elements
			else
				all content is NORMAL_TEXT
	*/
	panic("Not implemented yet !!!")
}

func parseInlineElements(blkElems []BlockElementShell) []ir.BlockElement {
	panic("Not implemented yet !!!")
}

func NewParser(tks []ir.Token) *Parser {
	return &Parser{
		tks:  tks,
		elms: make([]ir.BlockElement, len(tks)),
	}
}

func getTokensBlockElement(t ir.TokenType) ir.ElementType {
	r := ir.EL_NORMAL_TEXT

	switch t {
	case ir.TK_HEADING_1:
		return ir.EL_HEADING_1
	case ir.TK_HEADING_2:
		return ir.EL_HEADING_2
	case ir.TK_HEADING_3:
		return ir.EL_HEADING_3
	case ir.TK_HEADING_4:
		return ir.EL_HEADING_4
	case ir.TK_HEADING_5:
		return ir.EL_HEADING_5
	case ir.TK_HEADING_6:
		return ir.EL_HEADING_6
	case ir.TK_HORIZONTAL_LINE:
		return ir.EL_HORIZONTAL_LINE
	case ir.TK_BLOCK_QUOTE:
		return ir.EL_BLOCK_QUOTE
	case ir.TK_BULLET_POINT:
		return ir.EL_BULLET_POINT
	case ir.TK_LIST_SEQUENCE:
		return ir.EL_LIST_SEQUENCE
	case ir.TK_CODE_BLOCK:
		return ir.EL_CODE_BLOCK
	case ir.TK_IMAGE:
		return ir.EL_IMAGE
	case ir.TK_LINE_BREAK:
		return ir.EL_LINE_BREAK
	}

	return r
}
