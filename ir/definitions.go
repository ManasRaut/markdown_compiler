package ir

import (
	"fmt"
	"regexp"
)

type ElementDefinition struct {
	T                        ElementName
	StartToken               TokenType
	EndToken                 TokenType
	Category                 ElementCategory
	ContentType              ElementContentType
	IncludeStartTokenInValue bool
}

var HEADING_1_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_1,
	StartToken:               TK_HEADING_1,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var HEADING_2_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_2,
	StartToken:               TK_HEADING_2,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var HEADING_3_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_3,
	StartToken:               TK_HEADING_3,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var HEADING_4_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_4,
	StartToken:               TK_HEADING_4,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var HEADING_5_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_5,
	StartToken:               TK_HEADING_5,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var HEADING_6_DEFINITION = ElementDefinition{
	T:                        EL_HEADING_6,
	StartToken:               TK_HEADING_6,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var BULLET_POINT_DEFINITION = ElementDefinition{
	T:                        EL_BULLET_POINT,
	StartToken:               TK_BULLET_POINT,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var LIST_SEQUENCE_DEFINITION = ElementDefinition{
	T:                        EL_LIST_SEQUENCE,
	StartToken:               TK_LIST_SEQUENCE,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var PLAIN_TEXT_DEFINITION = ElementDefinition{
	T:                        EL_PLAIN_TEXT,
	StartToken:               TK_PLAIN_TEXT,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: true,
}
var CHECKED_BOX_DEFINITION = ElementDefinition{
	T:                        EL_CHECKED_BOX,
	StartToken:               TK_CHECKED_BOX,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var UNCHECKED_BOX_DEFINITION = ElementDefinition{
	T:                        EL_UNCHECKED_BOX,
	StartToken:               TK_UNCHECKED_BOX,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}

// ---------------------- Multi Line Block elements ----------------------
var CODE_BLOCK_DEFINITION = ElementDefinition{
	T:                        EL_CODE_BLOCK,
	StartToken:               TK_CODE_BLOCK,
	EndToken:                 TK_CODE_BLOCK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_PLAIN_TEXT,
	IncludeStartTokenInValue: false,
}
var BLOCK_QUOTE_DEFINITION = ElementDefinition{
	T:                        EL_BLOCK_QUOTE,
	StartToken:               TK_BLOCK_QUOTE,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_BLOCK,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}

// ---------------------- Self contained elements ----------------------
var LINE_BREAK_DEFINITION = ElementDefinition{
	T:                        EL_LINE_BREAK,
	StartToken:               TK_LINE_BREAK,
	EndToken:                 TK_LINE_BREAK,
	Category:                 CATEGORY_SELF_CONTAINED,
	ContentType:              CONTENT_TYPE_NONE,
	IncludeStartTokenInValue: false,
}
var HORIZONTAL_LINE_DEFINITION = ElementDefinition{
	T:                        EL_HORIZONTAL_LINE,
	StartToken:               TK_HORIZONTAL_LINE,
	EndToken:                 TK_HORIZONTAL_LINE,
	Category:                 CATEGORY_SELF_CONTAINED,
	ContentType:              CONTENT_TYPE_NONE,
	IncludeStartTokenInValue: false,
}
var IMAGE_DEFINITION = ElementDefinition{
	T:                        EL_IMAGE,
	StartToken:               TK_IMAGE,
	EndToken:                 TK_IMAGE,
	Category:                 CATEGORY_SELF_CONTAINED,
	ContentType:              CONTENT_TYPE_MEDIA,
	IncludeStartTokenInValue: false,
}

// ---------------------- Inline elements ----------------------
var BOLD_AND_ITALIC_DEFINITION = ElementDefinition{
	T:                        EL_BOLD_AND_ITALIC,
	StartToken:               TK_BOLD_AND_ITALIC,
	EndToken:                 TK_BOLD_AND_ITALIC,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var BOLD_DEFINITION = ElementDefinition{
	T:                        EL_BOLD,
	StartToken:               TK_BOLD,
	EndToken:                 TK_BOLD,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var UNDERLINE_DEFINITION = ElementDefinition{
	T:                        EL_UNDERLINE,
	StartToken:               TK_UNDERLINE,
	EndToken:                 TK_UNDERLINE,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var ITALIC_DEFINITION = ElementDefinition{
	T:                        EL_ITALIC,
	StartToken:               TK_ITALIC,
	EndToken:                 TK_ITALIC,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var EMPHASIS_DEFINITION = ElementDefinition{
	T:                        EL_EMPHASIS,
	StartToken:               TK_EMPHASIS,
	EndToken:                 TK_EMPHASIS,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var ESCAPE_CHARACTER_DEFINITION = ElementDefinition{
	T:                        EL_ESCAPE_CHARACTER,
	StartToken:               TK_ESCAPE_CHARACTER,
	EndToken:                 TK_ESCAPE_CHARACTER,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
}
var STRIKETHROUGH_DEFINITION = ElementDefinition{
	T:                        EL_STRIKETHROUGH,
	StartToken:               TK_STRIKETHROUGH,
	EndToken:                 TK_STRIKETHROUGH,
	Category:                 CATEGORY_INLINE,
	ContentType:              CONTENT_TYPE_INLINE_ELEMENTS,
	IncludeStartTokenInValue: false,
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
	TK_PLAIN_TEXT:    PLAIN_TEXT_DEFINITION,
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
	TK_UNDERLINE:        UNDERLINE_DEFINITION,
	TK_BOLD:             BOLD_DEFINITION,
	TK_ITALIC:           ITALIC_DEFINITION,
	TK_EMPHASIS:         EMPHASIS_DEFINITION,
	TK_ESCAPE_CHARACTER: ESCAPE_CHARACTER_DEFINITION,
	TK_STRIKETHROUGH:    STRIKETHROUGH_DEFINITION,
}

func MetadataHandler(def ElementDefinition, s string) string {
	if def == IMAGE_DEFINITION {
		return ImageMetadataHandler(s)
	}
	return s
}

func ImageMetadataHandler(s string) string {
	regex := regexp.MustCompile(`!\[(.+)\]\((.+)\)`)
	res := regex.FindStringSubmatch(s)
	if len(res) < 2 {
		return s
	}
	if len(res) < 3 {
		return fmt.Sprintf(`{"label":"%s","url":""}`, res[1])
	}
	return fmt.Sprintf(`{"label":"%s","url":"%s"}`, res[1], res[2])
}
