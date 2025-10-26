package ir

import "fmt"

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
	return fmt.Sprintf("%v", t.T)
}

func (a *Token) Equal(b *Token) bool {
	return a.T == b.T && a.V == b.V
}
