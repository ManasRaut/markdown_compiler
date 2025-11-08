package lexer

import (
	_ "embed"
	"regexp"
	"slices"

	"github.com/ManasRaut/md_lex/ir"
)

type TokenMatcher struct {
	tType   ir.TokenType
	regex   *regexp.Regexp
	handler func(tType ir.TokenType, value string) ir.Token
}

type Lexer struct {
	source   string
	pos      int
	tokens   []ir.Token
	matchers []TokenMatcher
}

func InitLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		pos:    0,
		tokens: make([]ir.Token, 0),
		matchers: []TokenMatcher{
			{tType: ir.TK_HORIZONTAL_LINE, regex: regexp.MustCompile(`^-{3,}`), handler: commonHandler},
			{tType: ir.TK_BLOCK_QUOTE, regex: regexp.MustCompile(`^>+ `), handler: commonHandler},
			{tType: ir.TK_BULLET_POINT, regex: regexp.MustCompile(`^- `), handler: commonHandler},
			{tType: ir.TK_LIST_SEQUENCE, regex: regexp.MustCompile(`^[0-9]+. `), handler: commonHandler},
			{tType: ir.TK_CODE_BLOCK, regex: regexp.MustCompile("^```"), handler: commonHandler},
			{tType: ir.TK_IMAGE, regex: regexp.MustCompile(`^!\[.+\]\(.+\)`), handler: commonHandler},
			{tType: ir.TK_CHECKED_BOX, regex: regexp.MustCompile(`^\[\.\]`), handler: commonHandler},
			{tType: ir.TK_UNCHECKED_BOX, regex: regexp.MustCompile(`^\[ \]`), handler: commonHandler},
			{tType: ir.TK_UNDERLINE, regex: regexp.MustCompile(`^__`), handler: commonHandler},
			{tType: ir.TK_BOLD_AND_ITALIC, regex: regexp.MustCompile(`^\*{3}`), handler: commonHandler},
			{tType: ir.TK_BOLD, regex: regexp.MustCompile(`^\*{2}`), handler: commonHandler},
			{tType: ir.TK_ITALIC, regex: regexp.MustCompile(`^\*{1}`), handler: commonHandler},
			{tType: ir.TK_EMPHASIS, regex: regexp.MustCompile("^`"), handler: commonHandler},
			{tType: ir.TK_ESCAPE_CHARACTER, regex: regexp.MustCompile(`^\\`), handler: commonHandler},
			{tType: ir.TK_STRIKETHROUGH, regex: regexp.MustCompile(`^~{2}`), handler: commonHandler},

			{tType: ir.TK_HEADING_1, regex: regexp.MustCompile(`^# `), handler: commonHandler},
			{tType: ir.TK_HEADING_2, regex: regexp.MustCompile(`^## `), handler: commonHandler},
			{tType: ir.TK_HEADING_3, regex: regexp.MustCompile(`^### `), handler: commonHandler},
			{tType: ir.TK_HEADING_4, regex: regexp.MustCompile(`^#### `), handler: commonHandler},
			{tType: ir.TK_HEADING_5, regex: regexp.MustCompile(`^##### `), handler: commonHandler},
			{tType: ir.TK_HEADING_6, regex: regexp.MustCompile(`^###### `), handler: commonHandler},
			{tType: ir.TK_LINE_BREAK, regex: regexp.MustCompile("^\n"), handler: commonHandler},
			// {tType: NORMAL_TEXT, regex: regexp.MustCompile(`.+`), handler: textHandler},
			// {tType: NORMAL_TEXT, regex: regexp.MustCompile(`(.+?)[~~,\\,` + "`" + `,\*{3,3},\*{2,2},\*{1,1}]+`), handler: textHandler},
			{tType: ir.TK_PLAIN_TEXT, regex: regexp.MustCompile(`^[^\\~_` + "`" + `*\r\n]+`), handler: textHandler},
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

func (lex *Lexer) Parse() []ir.Token {

	// TODO: in infinite because of inline matching. FIX IT!!!
	for !lex.isEnd() {
		currSource := lex.next()
		tokenMatched := false
		for _, matcher := range lex.matchers {
			loc := matcher.regex.FindStringIndex(currSource)
			if loc != nil {
				value := currSource[loc[0]:loc[1]]
				token := matcher.handler(matcher.tType, value)
				lex.tokens = append(lex.tokens, token)
				lex.move(len(value))
				tokenMatched = true
				break
			}
		}

		if tokenMatched {
			// Combine any consecutive normal texts
			tokenCount := len(lex.tokens)
			i := tokenCount - 1
			combinedString := ""
			for i >= 0 && lex.tokens[i].T == ir.TK_PLAIN_TEXT {
				combinedString = lex.tokens[i].V + combinedString
				i--
			}
			if tokenCount-i-1 > 1 {
				lex.tokens[i+1].V = combinedString
				lex.tokens = slices.Delete(lex.tokens, i+2, tokenCount)
			}
			continue
		}

		// If no matcher matched , even normal text then get only the first character as a normal text
		// To prevent a infinite loop
		lex.tokens = append(lex.tokens, textHandler(ir.TK_PLAIN_TEXT, currSource[0:1]))
		lex.move(1)
	}

	return lex.tokens
}

func commonHandler(tType ir.TokenType, value string) ir.Token {
	return ir.Token{T: tType, V: value}
}

func textHandler(tType ir.TokenType, value string) ir.Token {
	return ir.Token{T: ir.TK_PLAIN_TEXT, V: value}
}
