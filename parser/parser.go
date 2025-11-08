package parser

import (
	"strings"

	"github.com/ManasRaut/md_lex/ir"
)

type unFinishedElement struct {
	Def      ir.ElementDefinition
	Metadata string
	V        []ir.Token
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
		parsedElement := parseInlineElements(&el)
		parsedElement.Metadata = el.Metadata
		p.elms = append(p.elms, parsedElement)
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
			Def: ir.ElementDefinitions[ir.TK_PLAIN_TEXT],
			V: []ir.Token{{
				T: ir.TK_PLAIN_TEXT, V: t.V,
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
				Def: ir.ElementDefinitions[ir.TK_PLAIN_TEXT],
				V:   []ir.Token{{T: ir.TK_PLAIN_TEXT, V: tkns[i].V}},
			}, i + 1
		}
		nextIndex++
	}

	s := i + 1
	if def.T == ir.EL_PLAIN_TEXT {
		s--
	}
	b := unFinishedElement{
		Def:      def,
		V:        tkns[s:j],
		Metadata: tkns[i].V,
	}
	return b, nextIndex
}

func parseSelfStandingElements(i int, tkns []ir.Token, def ir.ElementDefinition) (unFinishedElement, int) {
	b := unFinishedElement{
		Def:      def,
		V:        []ir.Token{},
		Metadata: tkns[i].V,
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

// TODO: Refactor this entire function later
func parseInlineElementsTokens(tkns []ir.Token) []*ir.MarkdownElement {
	children := make([]*ir.MarkdownElement, 0)

	startToken := ir.TK_UNKNOWN
	startIdx := 0

	i := 0
	l := len(tkns)

	for i <= l {

		if i >= l {
			if startToken == ir.TK_UNKNOWN {
				break
			}

			if startToken == ir.TK_PLAIN_TEXT {
				childValue := joinToString(tkns[startIdx:l])
				child := ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, childValue, nil)
				children = append(children, child)
				break
			}

			startToken = ir.TK_PLAIN_TEXT
			i = startIdx + 1
			continue
		}

		currTkn := tkns[i]
		currTknDef := ir.ElementDefinitions[currTkn.T]

		if startToken == ir.TK_UNKNOWN {
			startToken = currTknDef.StartToken
			startIdx = i
			i++
			continue
		}

		if startToken == ir.TK_PLAIN_TEXT {
			if currTkn.T != ir.TK_PLAIN_TEXT {
				childValue := joinToString(tkns[startIdx:i])
				children = append(children, ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, childValue, nil))
				startToken = ir.TK_UNKNOWN
				continue
			}
			i++
			continue
		}

		startTokenDef := ir.ElementDefinitions[startToken]

		if startTokenDef.EndToken == currTkn.T {

			grandChildren := parseInlineElementsTokens(tkns[startIdx+1 : i])

			hasAllNormalText := true
			childValue := strings.Builder{}
			for _, tkn := range grandChildren {
				if tkn.Def != ir.PLAIN_TEXT_DEFINITION {
					hasAllNormalText = false
					break
				}
				childValue.WriteString(tkn.V)
			}
			if hasAllNormalText {
				children = append(children, ir.NewMarkDownElement(startTokenDef, childValue.String(), nil))
			} else {
				children = append(children, ir.NewMarkDownElement(startTokenDef, "", grandChildren))
			}
			startToken = ir.TK_UNKNOWN
		}

		i++
	}

	// TODO: Refactor later
	compactedChildren := make([]*ir.MarkdownElement, 0, len(children))
	checkpointIdx := 0
	someChildrensValue := strings.Builder{}
	for idx := 0; idx <= len(children); idx++ {
		var child *ir.MarkdownElement = nil
		if idx < len(children) {
			child = children[idx]
		}
		if child != nil && child.Def == ir.PLAIN_TEXT_DEFINITION {
			someChildrensValue.WriteString(child.V)
		} else if checkpointIdx != idx {
			compactedChildren = append(compactedChildren, ir.NewMarkDownElement(ir.PLAIN_TEXT_DEFINITION, someChildrensValue.String(), nil))
			someChildrensValue = strings.Builder{}
			checkpointIdx = idx
			idx--
		} else if child != nil {
			compactedChildren = append(compactedChildren, child)
			checkpointIdx = idx + 1
		}
	}

	return compactedChildren
	// return children
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
		elms: make([]*ir.MarkdownElement, 0),
	}
}
