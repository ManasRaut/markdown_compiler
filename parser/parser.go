package parser

import (
	"github.com/ManasRaut/lexe/ir"
)

type unFinishedElement struct {
	Def elementDefinition
	V   []ir.Token
}

type Parser struct {
	tks  []ir.Token
	elms []ir.BlockElement
}

func (p *Parser) GetElements() []ir.BlockElement {
	return p.elms
}

func (p *Parser) Parse() error {
	blkElems := parseBlockElements(p.tks)
	p.elms = parseInlineElements(blkElems)
	return nil
}

func parseBlockElements(tkns []ir.Token) []unFinishedElement {

	r := make([]unFinishedElement, 0, len(tkns))
	i := 0
	l := len(tkns)

	for {

		if i >= l {
			break
		}

		t := tkns[i]
		def := elementDefinitions[t.T]
		// Default is normal text
		el := unFinishedElement{
			Def: elementDefinitions[ir.TK_NORMAL_TEXT],
			V: []ir.Token{{
				T: ir.TK_NORMAL_TEXT, V: t.V,
			}},
		}

		switch def.Category {
		case ir.CATEGORY_BLOCK:
			el, i = parseAllBlockElements(i, tkns, def)
		case ir.CATEGORY_SELF_CONTAINED:
			el, i = parseSelfStandingElements(i, tkns, def)
		default:
			i++
		}
		r = append(r, el)
	}

	return r
}

func parseAllBlockElements(i int, tkns []ir.Token, def elementDefinition) (unFinishedElement, int) {
	j := i
	l := len(tkns)
	for {
		if j != i && (j >= l || tkns[j].T == def.EndToken) {
			break
		}
		j++
	}

	nextIndex := j
	if def.StartToken == def.EndToken {
		if j >= l {
			return unFinishedElement{
				Def: elementDefinitions[ir.TK_NORMAL_TEXT],
				V:   []ir.Token{{T: ir.TK_NORMAL_TEXT, V: tkns[i].V}},
			}, i + 1
		}
		nextIndex++
	}

	s := i + 1
	if def.T == ir.EL_NORMAL_TEXT {
		s--
	}
	b := unFinishedElement{
		Def: def,
		V:   tkns[s:j],
	}
	return b, nextIndex
}

func parseSelfStandingElements(i int, tkns []ir.Token, def elementDefinition) (unFinishedElement, int) {
	b := unFinishedElement{
		Def: def,
		V:   []ir.Token{},
	}
	return b, i + 1
}

func parseInlineElements(blkElems []unFinishedElement) []ir.BlockElement {
	panic("Not implemented yet !!!")
}

func NewParser(tks []ir.Token) *Parser {
	return &Parser{
		tks:  tks,
		elms: make([]ir.BlockElement, len(tks)),
	}
}
