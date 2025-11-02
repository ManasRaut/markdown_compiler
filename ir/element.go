package ir

import (
	"fmt"
	"strings"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ElementName string

const (
	EL_HEADING_1        ElementName = "HEADING_1"
	EL_HEADING_2        ElementName = "HEADING_2"
	EL_HEADING_3        ElementName = "HEADING_3"
	EL_HEADING_4        ElementName = "HEADING_4"
	EL_HEADING_5        ElementName = "HEADING_5"
	EL_HEADING_6        ElementName = "HEADING_6"
	EL_NORMAL_TEXT      ElementName = "NORMAL_TEXT"
	EL_LINE_BREAK       ElementName = "LINE_BREAK"
	EL_HORIZONTAL_LINE  ElementName = "HORIZONTAL_LINE"
	EL_BLOCK_QUOTE      ElementName = "BLOCK_QUOTE"
	EL_BULLET_POINT     ElementName = "BULLET_POINT"
	EL_LIST_SEQUENCE    ElementName = "LIST_SEQUENCE"
	EL_CODE_BLOCK       ElementName = "CODE_BLOCK"
	EL_IMAGE            ElementName = "IMAGE"
	EL_BOLD             ElementName = "BOLD"
	EL_UNDERLINE        ElementName = "UNDERLINE"
	EL_ITALIC           ElementName = "ITALIC"
	EL_BOLD_AND_ITALIC  ElementName = "BOLD_AND_ITALIC"
	EL_STRIKETHROUGH    ElementName = "STRIKETHROUGH"
	EL_EMPHASIS         ElementName = "EMPHASIS"
	EL_HYPER_LINK       ElementName = "HYPER_LINK"
	EL_ESCAPE_CHARACTER ElementName = "ESCAPE_CHARACTER"
	EL_CHECKED_BOX      ElementName = "CHECKED_BOX"
	EL_UNCHECKED_BOX    ElementName = "UNCHECKED_BOX"
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

type MarkdownElement struct {
	Id       string
	Def      ElementDefinition
	V        string
	C        []*MarkdownElement
	Metadata string
}

func (a *MarkdownElement) Equal(b *MarkdownElement) bool {
	return a.Id == b.Id
}

func (m MarkdownElement) String() string {
	childsString := "nil"
	if m.C != nil {
		b := strings.Builder{}
		b.WriteString("[")
		for _, c := range m.C {
			if c != nil {
				b.WriteString(c.String())
				b.WriteString(", ")
			}
		}
		b.WriteString("]")
		childsString = b.String()
	}
	// return fmt.Sprintf("%s::%s(`%s`,`%s`)", m.Def.T, m.Id, m.V, childsString)
	return fmt.Sprintf("%s(`%s`,`%s`)", m.Def.T, m.V, childsString)
}

func NewMarkDownElement(Def ElementDefinition, V string, C []*MarkdownElement) *MarkdownElement {
	id, err := gonanoid.New(10)
	if err != nil {
		panic(fmt.Sprintf("Error while creating new markdown element : %v", err))
	}
	return &MarkdownElement{
		Id:  id,
		Def: Def,
		V:   V,
		C:   C,
	}
}

// ************************************************************
// 						Inline Elements
// ************************************************************

// type InlineElement struct {
// 	T ElementName
// 	V string
// }

// func (a *InlineElement) Equal(b *InlineElement) bool {
// 	return a.T == b.T && a.V == b.V
// }
