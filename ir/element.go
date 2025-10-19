package ir

type ElementType string

const (
	EL_HEADING_1        ElementType = "HEADING_1"
	EL_HEADING_2        ElementType = "HEADING_2"
	EL_HEADING_3        ElementType = "HEADING_3"
	EL_HEADING_4        ElementType = "HEADING_4"
	EL_HEADING_5        ElementType = "HEADING_5"
	EL_HEADING_6        ElementType = "HEADING_6"
	EL_NORMAL_TEXT      ElementType = "NORMAL_TEXT"
	EL_LINE_BREAK       ElementType = "LINE_BREAK"
	EL_HORIZONTAL_LINE  ElementType = "HORIZONTAL_LINE"
	EL_BLOCK_QUOTE      ElementType = "BLOCK_QUOTE"
	EL_BULLET_POINT     ElementType = "BULLET_POINT"
	EL_LIST_SEQUENCE    ElementType = "LIST_SEQUENCE"
	EL_CODE_BLOCK       ElementType = "CODE_BLOCK"
	EL_IMAGE            ElementType = "IMAGE"
	EL_BOLD             ElementType = "BOLD"
	EL_ITALIC           ElementType = "ITALIC"
	EL_BOLD_AND_ITALIC  ElementType = "BOLD_AND_ITALIC"
	EL_STRIKETHROUGH    ElementType = "STRIKETHROUGH"
	EL_EMPHASIS         ElementType = "EMPHASIS"
	EL_HYPER_LINK       ElementType = "HYPER_LINK"
	EL_ESCAPE_CHARACTER ElementType = "ESCAPE_CHARACTER"
	EL_CHECKED_BOX      ElementType = "CHECKED_BOX"
	EL_UNCHECKED_BOX    ElementType = "UNCHECKED_BOX"
)

type ElementCategory int

const (
	CATEGORY_BLOCK ElementCategory = iota
	CATEGORY_SELF_CONTAINED
	CATEGORY_INLINE
)

type ElementContentType int

const (
	CONTENT_TYPE_PLAIN_TEXT ElementContentType = iota
	CONTENT_TYPE_INLINE_ELEMENTS
	CONTENT_TYPE_NONE
)

// ************************************************************
// 						Block Elements
// ************************************************************

type BlockElement struct {
	T ElementType
	V []InlineElement
}

func (a *BlockElement) Equal(b *BlockElement) bool {

	if a.T != b.T || len(a.V) != len(b.V) {
		return false
	}

	for i := range len(a.V) {
		if !a.V[i].Equal(&b.V[i]) {
			return false
		}
	}

	return true
}

// ************************************************************
// 						Inline Elements
// ************************************************************

type InlineElement struct {
	T ElementType
	V string
}

func (a *InlineElement) Equal(b *InlineElement) bool {
	return a.T == b.T && a.V == b.V
}
