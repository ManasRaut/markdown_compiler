package ir

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	TK_HEADING_1        TokenType = "HEADING_1"
	TK_HEADING_2        TokenType = "HEADING_2"
	TK_HEADING_3        TokenType = "HEADING_3"
	TK_HEADING_4        TokenType = "HEADING_4"
	TK_HEADING_5        TokenType = "HEADING_5"
	TK_HEADING_6        TokenType = "HEADING_6"
	TK_NORMAL_TEXT      TokenType = "NORMAL_TEXT"
	TK_LINE_BREAK       TokenType = "LINE_BREAK"
	TK_HORIZONTAL_LINE  TokenType = "HORIZONTAL_LINE"
	TK_BLOCK_QUOTE      TokenType = "BLOCK_QUOTE"
	TK_BULLET_POINT     TokenType = "BULLET_POINT"
	TK_LIST_SEQUENCE    TokenType = "LIST_SEQUENCE"
	TK_CODE_BLOCK       TokenType = "CODE_BLOCK"
	TK_IMAGE            TokenType = "IMAGE"
	TK_BOLD             TokenType = "BOLD"
	TK_ITALIC           TokenType = "ITALIC"
	TK_BOLD_AND_ITALIC  TokenType = "BOLD_AND_ITALIC"
	TK_STRIKETHROUGH    TokenType = "STRIKETHROUGH"
	TK_EMPHASIS         TokenType = "EMPHASIS"
	TK_HYPER_LINK       TokenType = "HYPER_LINK"
	TK_ESCAPE_CHARACTER TokenType = "ESCAPE_CHARACTER"
	TK_CHECKED_BOX      TokenType = "CHECKED_BOX"
	TK_UNCHECKED_BOX    TokenType = "UNCHECKED_BOX"
	TK_UNKNOWN          TokenType = "UNKNOWN"
)

type Token struct {
	T TokenType
	V string
}

func (t Token) String() string {
	return fmt.Sprintf("%v(%s)", t.T, escape(t.V))
}

func (a *Token) Equal(b *Token) bool {
	return a.T == b.T && a.V == b.V
}

func escape(v string) string {
	return strings.ReplaceAll(strings.ReplaceAll(v, "\n", "\\n"), "\t", "\\t")
}

func GetTokentype(v string) TokenType {
	switch v {
	case "HEADING_1":
		return TK_HEADING_1
	case "HEADING_2":
		return TK_HEADING_2
	case "HEADING_3":
		return TK_HEADING_3
	case "HEADING_4":
		return TK_HEADING_4
	case "HEADING_5":
		return TK_HEADING_5
	case "HEADING_6":
		return TK_HEADING_6
	case "NORMAL_TEXT":
		return TK_NORMAL_TEXT
	case "LINE_BREAK":
		return TK_LINE_BREAK
	case "HORIZONTAL_LINE":
		return TK_HORIZONTAL_LINE
	case "BLOCK_QUOTE":
		return TK_BLOCK_QUOTE
	case "BULLET_POINT":
		return TK_BULLET_POINT
	case "LIST_SEQUENCE":
		return TK_LIST_SEQUENCE
	case "CODE_BLOCK":
		return TK_CODE_BLOCK
	case "IMAGE":
		return TK_IMAGE
	case "BOLD":
		return TK_BOLD
	case "ITALIC":
		return TK_ITALIC
	case "BOLD_AND_ITALIC":
		return TK_BOLD_AND_ITALIC
	case "STRIKETHROUGH":
		return TK_STRIKETHROUGH
	case "EMPHASIS":
		return TK_EMPHASIS
	case "HYPER_LINK":
		return TK_HYPER_LINK
	case "ESCAPE_CHARACTER":
		return TK_ESCAPE_CHARACTER
	case "CHECKED_BOX":
		return TK_CHECKED_BOX
	case "UNCHECKED_BOX":
		return TK_UNCHECKED_BOX
	default:
		return TK_UNKNOWN
	}
}
