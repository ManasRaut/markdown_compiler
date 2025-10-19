package parser

import "github.com/ManasRaut/lexe/ir"

type elementDefinition struct {
	T           ir.ElementType
	StartToken  ir.TokenType
	EndToken    ir.TokenType
	Category    ir.ElementCategory
	ContentType ir.ElementContentType
}

var elementDefinitions map[ir.TokenType]elementDefinition = map[ir.TokenType]elementDefinition{
	// ---------------------- Single Line Block elements ----------------------
	ir.TK_HEADING_1: {
		T:           ir.EL_HEADING_1,
		StartToken:  ir.TK_HEADING_1,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_HEADING_2: {
		T:           ir.EL_HEADING_2,
		StartToken:  ir.TK_HEADING_2,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_HEADING_3: {
		T:           ir.EL_HEADING_3,
		StartToken:  ir.TK_HEADING_3,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_HEADING_4: {
		T:           ir.EL_HEADING_4,
		StartToken:  ir.TK_HEADING_4,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_HEADING_5: {
		T:           ir.EL_HEADING_5,
		StartToken:  ir.TK_HEADING_5,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_HEADING_6: {
		T:           ir.EL_HEADING_6,
		StartToken:  ir.TK_HEADING_6,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_BULLET_POINT: {
		T:           ir.EL_BULLET_POINT,
		StartToken:  ir.TK_BULLET_POINT,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_LIST_SEQUENCE: {
		T:           ir.EL_LIST_SEQUENCE,
		StartToken:  ir.TK_LIST_SEQUENCE,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_NORMAL_TEXT: {
		T:           ir.EL_NORMAL_TEXT,
		StartToken:  ir.TK_NORMAL_TEXT,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	// ---------------------- Multi Line Block elements ----------------------
	ir.TK_CODE_BLOCK: {
		T:           ir.EL_CODE_BLOCK,
		StartToken:  ir.TK_CODE_BLOCK,
		EndToken:    ir.TK_CODE_BLOCK,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_PLAIN_TEXT,
	},
	ir.TK_BLOCK_QUOTE: {
		T:           ir.EL_BLOCK_QUOTE,
		StartToken:  ir.TK_BLOCK_QUOTE,
		EndToken:    ir.TK_BLOCK_QUOTE,
		Category:    ir.CATEGORY_BLOCK,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	// ---------------------- Self contained elements ----------------------
	ir.TK_LINE_BREAK: {
		T:           ir.EL_LINE_BREAK,
		StartToken:  ir.TK_LINE_BREAK,
		EndToken:    ir.TK_LINE_BREAK,
		Category:    ir.CATEGORY_SELF_CONTAINED,
		ContentType: ir.CONTENT_TYPE_NONE,
	},
	ir.TK_HORIZONTAL_LINE: {
		T:           ir.EL_HORIZONTAL_LINE,
		StartToken:  ir.TK_HORIZONTAL_LINE,
		EndToken:    ir.TK_HORIZONTAL_LINE,
		Category:    ir.CATEGORY_SELF_CONTAINED,
		ContentType: ir.CONTENT_TYPE_NONE,
	},
	ir.TK_IMAGE: {
		T:           ir.EL_IMAGE,
		StartToken:  ir.TK_IMAGE,
		EndToken:    ir.TK_IMAGE,
		Category:    ir.CATEGORY_SELF_CONTAINED,
		ContentType: ir.CONTENT_TYPE_NONE,
	},
	// ---------------------- Inline elements ----------------------
	ir.TK_CHECKED_BOX: {
		T:           ir.EL_CHECKED_BOX,
		StartToken:  ir.TK_CHECKED_BOX,
		EndToken:    ir.TK_CHECKED_BOX,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_UNCHECKED_BOX: {
		T:           ir.EL_UNCHECKED_BOX,
		StartToken:  ir.TK_UNCHECKED_BOX,
		EndToken:    ir.TK_UNCHECKED_BOX,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_BOLD_AND_ITALIC: {
		T:           ir.EL_BOLD_AND_ITALIC,
		StartToken:  ir.TK_BOLD_AND_ITALIC,
		EndToken:    ir.TK_BOLD_AND_ITALIC,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_BOLD: {
		T:           ir.EL_BOLD,
		StartToken:  ir.TK_BOLD,
		EndToken:    ir.TK_BOLD,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_ITALIC: {
		T:           ir.EL_ITALIC,
		StartToken:  ir.TK_ITALIC,
		EndToken:    ir.TK_ITALIC,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_EMPHASIS: {
		T:           ir.EL_EMPHASIS,
		StartToken:  ir.TK_EMPHASIS,
		EndToken:    ir.TK_EMPHASIS,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_ESCAPE_CHARACTER: {
		T:           ir.EL_ESCAPE_CHARACTER,
		StartToken:  ir.TK_ESCAPE_CHARACTER,
		EndToken:    ir.TK_ESCAPE_CHARACTER,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
	ir.TK_STRIKETHROUGH: {
		T:           ir.EL_STRIKETHROUGH,
		StartToken:  ir.TK_STRIKETHROUGH,
		EndToken:    ir.TK_STRIKETHROUGH,
		Category:    ir.CATEGORY_INLINE,
		ContentType: ir.CONTENT_TYPE_INLINE_ELEMENTS,
	},
}
