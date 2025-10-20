package parser

import (
	"strings"

	"github.com/ManasRaut/lexe/ir"
)

type unFinishedElement struct {
	Def ir.ElementDefinition
	V   []ir.Token
}

type Parser struct {
	tks  []ir.Token
	elms []*ir.MarkdownElement
}

func (p *Parser) GetElements() []*ir.MarkdownElement {
	return p.elms
}

func (p *Parser) Parse() error {
	blkElems := parseBlockElements(p.tks)
	for _, el := range blkElems {
		p.elms = append(p.elms, parseInlineElements(&el))
	}
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
		def := ir.ElementDefinitions[t.T]
		// Default is normal text
		el := unFinishedElement{
			Def: ir.ElementDefinitions[ir.TK_NORMAL_TEXT],
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

func parseAllBlockElements(i int, tkns []ir.Token, def ir.ElementDefinition) (unFinishedElement, int) {
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
				Def: ir.ElementDefinitions[ir.TK_NORMAL_TEXT],
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

func parseSelfStandingElements(i int, tkns []ir.Token, def ir.ElementDefinition) (unFinishedElement, int) {
	b := unFinishedElement{
		Def: def,
		V:   []ir.Token{},
	}
	return b, i + 1
}

func parseInlineElements(el *unFinishedElement) *ir.MarkdownElement {
	if el.Def.ContentType == ir.CONTENT_TYPE_NONE {
		return ir.NewMarkDownElement(el.Def, "", nil)
	}
	if el.Def.ContentType == ir.CONTENT_TYPE_PLAIN_TEXT {
		return ir.NewMarkDownElement(el.Def, joinToString(el.V), nil)
	}

	if len(el.V) == 0 {
		return ir.NewMarkDownElement(el.Def, "", nil)
	}

	return ir.NewMarkDownElement(el.Def, "", parseInlineElementsTokens(el.V))
}

func parseInlineElementsTokens(tkns []ir.Token) []*ir.MarkdownElement {
	/*
		- Iterate throught tokens
		- if starts with a NORMAL_TEXT
			- When an inline element token is found or it ends create a normal_text element
		- If an token of an inline element is found
			- Start looking for its end token
			- If not found then the found token is a normal text element
			- if end token found, create that inline element and
				- call parseInlineElementsTokens on the tokens in between start and end token
				- and assign the returned children to it

	*/
	panic("Method Not implemented!!")
}

func joinToString(tkns []ir.Token) string {
	b := strings.Builder{}
	for _, tkn := range tkns {
		b.WriteString(tkn.V)
	}
	return b.String()
}

func NewParser(tks []ir.Token) *Parser {
	return &Parser{
		tks:  tks,
		elms: make([]*ir.MarkdownElement, len(tks)),
	}
}
