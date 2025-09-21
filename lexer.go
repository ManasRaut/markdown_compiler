package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

type TokenType string

const (
	// headings
	HEADING_1        TokenType = "HEADING_1"
	HEADING_2        TokenType = "HEADING_2"
	HEADING_3        TokenType = "HEADING_3"
	HEADING_4        TokenType = "HEADING_4"
	HEADING_5        TokenType = "HEADING_5"
	HEADING_6        TokenType = "HEADING_6"
	NORMAL_TEXT      TokenType = "NORMAL_TEXT"
	LINE_BREAK       TokenType = "LINE_BREAK"
	HORIZONTAL_LINE  TokenType = "HORIZONTAL_LINE"
	BLOCK_QUOTE      TokenType = "BLOCK_QUOTE"
	BULLET_POINT     TokenType = "BULLET_POINT"
	LIST_SEQUENCE    TokenType = "LIST_SEQUENCE"
	CODE_BLOCK       TokenType = "CODE_BLOCK"
	IMAGE            TokenType = "IMAGE"
	BOLD             TokenType = "BOLD"
	ITALIC           TokenType = "ITALIC"
	BOLD_AND_ITALIC  TokenType = "BOLD_AND_ITALIC"
	STRIKETHROUGH    TokenType = "STRIKETHROUGH"
	EMPHASIS         TokenType = "EMPHASIS"
	HYPER_LINK       TokenType = "HYPER_LINK"
	ESCAPE_CHARACTER TokenType = "ESCAPE_CHARACTER"
	CHECKED_BOX      TokenType = "CHECKED_BOX"
	UNCHECKED_BOX    TokenType = "UNCHECKED_BOX"
)

type Token struct {
	tType TokenType
	value string
}

func (t Token) String() string {
	// return fmt.Sprintf("%v\t\t ==> %v", t.tType, escape(t.value))
	return fmt.Sprintf("%v", t.tType)
}

type TokenMatcher struct {
	tType   TokenType
	regex   *regexp.Regexp
	handler func(tType TokenType, value string) Token
}

type Lexer struct {
	source   string
	pos      int
	tokens   []Token
	matchers []TokenMatcher
}

func initLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		pos:    0,
		tokens: make([]Token, 0),
		matchers: []TokenMatcher{
			{tType: HORIZONTAL_LINE, regex: regexp.MustCompile(`^-{3,}`), handler: commonHandler},
			{tType: BLOCK_QUOTE, regex: regexp.MustCompile(`^>+ `), handler: commonHandler},
			{tType: BULLET_POINT, regex: regexp.MustCompile(`^- `), handler: commonHandler},
			{tType: LIST_SEQUENCE, regex: regexp.MustCompile(`^[0-9]+. `), handler: commonHandler},
			{tType: CODE_BLOCK, regex: regexp.MustCompile("^```"), handler: commonHandler},
			{tType: IMAGE, regex: regexp.MustCompile(`^!\[.+\]\(.+\)`), handler: commonHandler},
			{tType: CHECKED_BOX, regex: regexp.MustCompile(`^\[ \]`), handler: commonHandler},
			{tType: UNCHECKED_BOX, regex: regexp.MustCompile(`^\[.\]`), handler: commonHandler},
			// {tType: BOLD_AND_ITALIC, regex: regexp.MustCompile(`^\*{3,3}`), handler: commonHandler},
			// {tType: BOLD, regex: regexp.MustCompile(`^\*{2,2}`), handler: commonHandler},
			// {tType: ITALIC, regex: regexp.MustCompile(`^\*{1,1}`), handler: commonHandler},
			// {tType: EMPHASIS, regex: regexp.MustCompile("^`"), handler: commonHandler},
			// {tType: ESCAPE_CHARACTER, regex: regexp.MustCompile(`^\\`), handler: commonHandler},
			// {tType: STRIKETHROUGH, regex: regexp.MustCompile(`^~{2,2}`), handler: commonHandler},

			{tType: HEADING_1, regex: regexp.MustCompile(`^# `), handler: commonHandler},
			{tType: HEADING_2, regex: regexp.MustCompile(`^## `), handler: commonHandler},
			{tType: HEADING_3, regex: regexp.MustCompile(`^### `), handler: commonHandler},
			{tType: HEADING_4, regex: regexp.MustCompile(`^#### `), handler: commonHandler},
			{tType: HEADING_5, regex: regexp.MustCompile(`^##### `), handler: commonHandler},
			{tType: HEADING_6, regex: regexp.MustCompile(`^###### `), handler: commonHandler},
			{tType: LINE_BREAK, regex: regexp.MustCompile("^\n"), handler: commonHandler},
			{tType: NORMAL_TEXT, regex: regexp.MustCompile(`.+`), handler: textHandler},
			// {tType: NORMAL_TEXT, regex: regexp.MustCompile(`(.+?)[~~,\\,` + "`" + `,\*{3,3},\*{2,2},\*{1,1}]+`), handler: textHandler},
			// },
		},
	}
}

func (lex *Lexer) next() string {
	return lex.source[lex.pos:]
}

func (lex *Lexer) move(n int) {
	lex.pos += n
}

func (lex *Lexer) isEnd() bool {
	return lex.pos >= len(lex.source)
}

func (lex *Lexer) parse() []Token {

	// TODO: in infinite because of inline matching. FIX IT!!!
	for !lex.isEnd() {
		for _, matcher := range lex.matchers {
			currSource := lex.next()
			loc := matcher.regex.FindStringIndex(currSource)
			if loc != nil {
				value := currSource[loc[0]:loc[1]]
				token := matcher.handler(matcher.tType, value)
				fmt.Printf("Matched at `%d:%d` for >>>%s<<< to %v\n", loc[0], loc[1], value, token)
				lex.tokens = append(lex.tokens, token)
				lex.move(len(value))
				fmt.Printf("\tMoved by %d\n", len(value))
				break
			}
		}
	}

	return lex.tokens
}

func commonHandler(tType TokenType, value string) Token {
	return Token{tType, value}
}

func textHandler(tType TokenType, value string) Token {
	return Token{NORMAL_TEXT, value}
}

//go:embed markdown_elements2.md
var exmapleSource string

func main() {
	lex := initLexer(exmapleSource)
	tokens := lex.parse()
	printTokens(tokens)
}

func printTokens(tokens []Token) {
	for _, token := range tokens {
		if token.tType == LINE_BREAK {
			fmt.Println()
		} else {
			fmt.Printf("%s ", token.String())
		}
	}
}

func escape(s string) string {
	return strings.ReplaceAll(s, "\n", "\\n")
}
