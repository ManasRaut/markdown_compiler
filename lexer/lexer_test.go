package lexer_test

// import (
// 	_ "embed"
// 	"fmt"
// 	"strings"
// 	"testing"

// 	"github.com/ManasRaut/lexe/lexer"
// )

// //go:embed test_files/headings.md
// var headingInput string

// //go:embed test_files/headings.tokens
// var headingOutput string

// func TestHeadings(t *testing.T) {
// 	lex := lexer.InitLexer(headingInput)
// 	tokens := lex.Parse()
// }

// func loadOutput(out string) []lexer.Token {
// 	const SEP = "--"
// 	lines := strings.Split(out, "\n")
// 	for _, line := range lines {
// 		arr := strings.Split(line, SEP)
// 		ttypeStr := arr[0]
// 		tValStr := arr[1]
// 	}

// 	tokens := make([]lexer.Token, 20)
// 	for _, tstr := range tokenstring {
// 		switch tstr {
// 		case "HEADING_1":
// 			tokens = append(tokens, lexer.HEADING_1)
// 			break
// 		}
// 	}
// }

// func getToken(val string) lexer.TokenType {
// 	switch val {
// 	case "HEADING_1":
// 		return lexer.HEADING_1

// 	case "HEADING_2":
// 		return lexer.HEADING_2
// 	case "HEADING_3":
// 		return lexer.HEADING_3
// 	case "HEADING_4":
// 		return lexer.HEADING_4
// 	case "HEADING_5":
// 		return lexer.HEADING_5
// 	case "HEADING_6":
// 		return lexer.HEADING_6
// 	case "NORMAL_TEXT":
// 		return lexer.NORMAL_TEXT
// 	case "LINE_BREAK":
// 		return lexer.LINE_BREAK
// 	case "HORIZONTAL_LINE":
// 		return lexer.HORIZONTAL_LINE
// 	case "BLOCK_QUOTE":
// 		return lexer.BLOCK_QUOTE
// 	case "BULLET_POINT":
// 		return lexer.BULLET_POINT
// 	case "LIST_SEQUENCE":
// 		return lexer.LIST_SEQUENCE
// 	case "CODE_BLOCK":
// 		return lexer.CODE_BLOCK
// 	case "IMAGE":
// 		return lexer.IMAGE
// 	case "BOLD":
// 		return lexer.BOLD
// 	case "ITALIC":
// 		return lexer.ITALIC
// 	case "BOLD_AND_ITALIC":
// 		return lexer.BOLD_AND_ITALIC
// 	case "STRIKETHROUGH":
// 		return lexer.STRIKETHROUGH
// 	case "EMPHASIS":
// 		return lexer.EMPHASIS
// 	case "HYPER_LINK":
// 		return lexer.HYPER_LINK
// 	case "ESCAPE_CHARACTER":
// 		return lexer.ESCAPE_CHARACTER
// 	case "CHECKED_BOX":
// 		return lexer.CHECKED_BOX
// 	case "UNCHECKED_BOX":
// 		return lexer.UNCHECKED_BOX
// 	default:
// 		break
// 	}
// 	fmt.Printf("invalid token value in test case=%s\n", val)
// 	panic("invalid token value in test case")
// }
