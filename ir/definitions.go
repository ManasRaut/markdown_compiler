package ir

type ElementDefinition struct {
	T           ElementName
	StartToken  TokenType
	EndToken    TokenType
	Category    ElementCategory
	ContentType ElementContentType
}

var HEADING_1_DEFINITION = ElementDefinition{
	T:           EL_HEADING_1,
	StartToken:  TK_HEADING_1,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var HEADING_2_DEFINITION = ElementDefinition{
	T:           EL_HEADING_2,
	StartToken:  TK_HEADING_2,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var HEADING_3_DEFINITION = ElementDefinition{
	T:           EL_HEADING_3,
	StartToken:  TK_HEADING_3,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var HEADING_4_DEFINITION = ElementDefinition{
	T:           EL_HEADING_4,
	StartToken:  TK_HEADING_4,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var HEADING_5_DEFINITION = ElementDefinition{
	T:           EL_HEADING_5,
	StartToken:  TK_HEADING_5,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var HEADING_6_DEFINITION = ElementDefinition{
	T:           EL_HEADING_6,
	StartToken:  TK_HEADING_6,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var BULLET_POINT_DEFINITION = ElementDefinition{
	T:           EL_BULLET_POINT,
	StartToken:  TK_BULLET_POINT,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var LIST_SEQUENCE_DEFINITION = ElementDefinition{
	T:           EL_LIST_SEQUENCE,
	StartToken:  TK_LIST_SEQUENCE,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var NORMAL_TEXT_DEFINITION = ElementDefinition{
	T:           EL_NORMAL_TEXT,
	StartToken:  TK_NORMAL_TEXT,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var CHECKED_BOX_DEFINITION = ElementDefinition{
	T:           EL_CHECKED_BOX,
	StartToken:  TK_CHECKED_BOX,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var UNCHECKED_BOX_DEFINITION = ElementDefinition{
	T:           EL_UNCHECKED_BOX,
	StartToken:  TK_UNCHECKED_BOX,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}

// ---------------------- Multi Line Block elements ----------------------
var CODE_BLOCK_DEFINITION = ElementDefinition{
	T:           EL_CODE_BLOCK,
	StartToken:  TK_CODE_BLOCK,
	EndToken:    TK_CODE_BLOCK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_PLAIN_TEXT,
}
var BLOCK_QUOTE_DEFINITION = ElementDefinition{
	T:           EL_BLOCK_QUOTE,
	StartToken:  TK_BLOCK_QUOTE,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_BLOCK,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}

// ---------------------- Self contained elements ----------------------
var LINE_BREAK_DEFINITION = ElementDefinition{
	T:           EL_LINE_BREAK,
	StartToken:  TK_LINE_BREAK,
	EndToken:    TK_LINE_BREAK,
	Category:    CATEGORY_SELF_CONTAINED,
	ContentType: CONTENT_TYPE_NONE,
}
var HORIZONTAL_LINE_DEFINITION = ElementDefinition{
	T:           EL_HORIZONTAL_LINE,
	StartToken:  TK_HORIZONTAL_LINE,
	EndToken:    TK_HORIZONTAL_LINE,
	Category:    CATEGORY_SELF_CONTAINED,
	ContentType: CONTENT_TYPE_NONE,
}
var IMAGE_DEFINITION = ElementDefinition{
	T:           EL_IMAGE,
	StartToken:  TK_IMAGE,
	EndToken:    TK_IMAGE,
	Category:    CATEGORY_SELF_CONTAINED,
	ContentType: CONTENT_TYPE_NONE,
}

// ---------------------- Inline elements ----------------------
//
//	var CHECKED_BOX_DEFINITION = ElementDefinition{
//		T:           EL_CHECKED_BOX,
//		StartToken:  TK_CHECKED_BOX,
//		EndToken:    TK_CHECKED_BOX,
//		Category:    CATEGORY_INLINE,
//		ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
//	}
//
//	var UNCHECKED_BOX_DEFINITION = ElementDefinition{
//		T:           EL_UNCHECKED_BOX,
//		StartToken:  TK_UNCHECKED_BOX,
//		EndToken:    TK_UNCHECKED_BOX,
//		Category:    CATEGORY_INLINE,
//		ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
//	}
var BOLD_AND_ITALIC_DEFINITION = ElementDefinition{
	T:           EL_BOLD_AND_ITALIC,
	StartToken:  TK_BOLD_AND_ITALIC,
	EndToken:    TK_BOLD_AND_ITALIC,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var BOLD_DEFINITION = ElementDefinition{
	T:           EL_BOLD,
	StartToken:  TK_BOLD,
	EndToken:    TK_BOLD,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var ITALIC_DEFINITION = ElementDefinition{
	T:           EL_ITALIC,
	StartToken:  TK_ITALIC,
	EndToken:    TK_ITALIC,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var EMPHASIS_DEFINITION = ElementDefinition{
	T:           EL_EMPHASIS,
	StartToken:  TK_EMPHASIS,
	EndToken:    TK_EMPHASIS,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var ESCAPE_CHARACTER_DEFINITION = ElementDefinition{
	T:           EL_ESCAPE_CHARACTER,
	StartToken:  TK_ESCAPE_CHARACTER,
	EndToken:    TK_ESCAPE_CHARACTER,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}
var STRIKETHROUGH_DEFINITION = ElementDefinition{
	T:           EL_STRIKETHROUGH,
	StartToken:  TK_STRIKETHROUGH,
	EndToken:    TK_STRIKETHROUGH,
	Category:    CATEGORY_INLINE,
	ContentType: CONTENT_TYPE_INLINE_ELEMENTS,
}

var ElementDefinitions map[TokenType]ElementDefinition = map[TokenType]ElementDefinition{
	// ---------------------- Single Line Block elements ----------------------
	TK_HEADING_1:     HEADING_1_DEFINITION,
	TK_HEADING_2:     HEADING_2_DEFINITION,
	TK_HEADING_3:     HEADING_3_DEFINITION,
	TK_HEADING_4:     HEADING_4_DEFINITION,
	TK_HEADING_5:     HEADING_5_DEFINITION,
	TK_HEADING_6:     HEADING_6_DEFINITION,
	TK_BULLET_POINT:  BULLET_POINT_DEFINITION,
	TK_LIST_SEQUENCE: LIST_SEQUENCE_DEFINITION,
	TK_NORMAL_TEXT:   NORMAL_TEXT_DEFINITION,
	TK_CHECKED_BOX:   CHECKED_BOX_DEFINITION,
	TK_UNCHECKED_BOX: UNCHECKED_BOX_DEFINITION,
	// ---------------------- Multi Line Block elements ----------------------
	TK_CODE_BLOCK:  CODE_BLOCK_DEFINITION,
	TK_BLOCK_QUOTE: BLOCK_QUOTE_DEFINITION,
	// ---------------------- Self contained elements ----------------------
	TK_LINE_BREAK:      LINE_BREAK_DEFINITION,
	TK_HORIZONTAL_LINE: HORIZONTAL_LINE_DEFINITION,
	TK_IMAGE:           IMAGE_DEFINITION,
	// ---------------------- Inline elements ----------------------
	// TK_CHECKED_BOX:      CHECKED_BOX_DEFINITION,
	// TK_UNCHECKED_BOX:    UNCHECKED_BOX_DEFINITION,
	TK_BOLD_AND_ITALIC:  BOLD_AND_ITALIC_DEFINITION,
	TK_BOLD:             BOLD_DEFINITION,
	TK_ITALIC:           ITALIC_DEFINITION,
	TK_EMPHASIS:         EMPHASIS_DEFINITION,
	TK_ESCAPE_CHARACTER: ESCAPE_CHARACTER_DEFINITION,
	TK_STRIKETHROUGH:    STRIKETHROUGH_DEFINITION,
}
