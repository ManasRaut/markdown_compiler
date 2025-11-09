package converters

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ManasRaut/md_lex/ir"
)

type htmlContext struct {
}

type HTMLMarkup string

type HTMLConverter struct {
}

func (c HTMLConverter) Convert(e []*ir.MarkdownElement) (*HTMLMarkup, error) {

	html, err := traverseAndConvert(&htmlContext{}, e)
	if err != nil {
		return nil, err
	}

	return &html, nil
}

func traverseAndConvert(ctx *htmlContext, elements []*ir.MarkdownElement) (HTMLMarkup, error) {
	html := strings.Builder{}

	for i := range elements {
		element := elements[i]

		childrenHTML, err := traverseAndConvert(ctx, element.C)

		if err != nil {
			return "", err
		}
		if len(element.C) == 0 {
			childrenHTML = HTMLMarkup(element.V)
		}

		startTag, endTag, childrenHTML := handleMetadata(element, childrenHTML)

		html.WriteString(fmt.Sprintf("%s%s%s", startTag, childrenHTML, endTag))
	}

	return HTMLMarkup(html.String()), nil
}

func handleMetadata(element *ir.MarkdownElement, childrenHTML HTMLMarkup) (string, string, HTMLMarkup) {

	var startTag, endTag string

	switch element.Def {
	case ir.LIST_SEQUENCE_DEFINITION:
		startTag, endTag = getTags(element.Def, fmt.Sprintf(`value="%s"`, element.Metadata))
		startTag = "<ol>" + startTag
		endTag = endTag + "</ol>"
	case ir.BULLET_POINT_DEFINITION:
		startTag, endTag = getTags(element.Def, fmt.Sprintf(`value="%s"`, element.Metadata))
		startTag = "<ul>" + startTag
		endTag = endTag + "</ul>"
	case ir.IMAGE_DEFINITION:
		metadata := make(map[string]string)
		json.Unmarshal([]byte(element.Metadata), &metadata)
		startTag, endTag = getTags(element.Def, fmt.Sprintf(`alt="%s" src="%s"`, metadata["label"], metadata["url"]))
	default:
		startTag, endTag = getTags(element.Def, "")
	}

	return startTag, endTag, childrenHTML
}

func getTags(def ir.ElementDefinition, attributes string) (string, string) {
	switch def {
	case ir.HEADING_1_DEFINITION:
		return "<h1>", "</h1>"
	case ir.HEADING_2_DEFINITION:
		return "<h2>", "</h2>"
	case ir.HEADING_3_DEFINITION:
		return "<h3>", "</h3>"
	case ir.HEADING_4_DEFINITION:
		return "<h4>", "</h4>"
	case ir.HEADING_5_DEFINITION:
		return "<h5>", "</h5>"
	case ir.HEADING_6_DEFINITION:
		return "<h6>", "</h6>"
	case ir.BULLET_POINT_DEFINITION:
		return "<li>", "</li>"
	case ir.LIST_SEQUENCE_DEFINITION:
		return fmt.Sprintf("<li %s>", attributes), "</li>"
	case ir.PLAIN_TEXT_DEFINITION:
		return "", ""
	case ir.CHECKED_BOX_DEFINITION:
		return `<input type="checkbox" checked>`, "</input>"
	case ir.UNCHECKED_BOX_DEFINITION:
		return `<input type="checkbox">`, "</input>"
	case ir.CODE_BLOCK_DEFINITION:
		return "<pre>", "</pre>"
	case ir.BLOCK_QUOTE_DEFINITION:
		return "<blockquote>", "</blockquote>"
	case ir.LINE_BREAK_DEFINITION:
		return "<br>", ""
	case ir.HORIZONTAL_LINE_DEFINITION:
		return "<hr>", ""
	case ir.IMAGE_DEFINITION:
		return fmt.Sprintf("<img %s />", attributes), ""
	case ir.BOLD_AND_ITALIC_DEFINITION:
		return "<b><i>", "</i></b>"
	case ir.UNDERLINE_DEFINITION:
		return "<u>", "</u>"
	case ir.BOLD_DEFINITION:
		return "<b>", "</b>"
	case ir.ITALIC_DEFINITION:
		return "<i>", "</i>"
	case ir.EMPHASIS_DEFINITION:
		return "<em>", "</em>"
	case ir.ESCAPE_CHARACTER_DEFINITION:
		return "<pre>", "</pre>"
	case ir.STRIKETHROUGH_DEFINITION:
		return "<strike>", "</strike>"
	default:
		return "", ""
	}
}
