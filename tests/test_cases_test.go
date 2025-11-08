package tests_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ManasRaut/md_lex/ir"
)

const token_seperator = "<<%>>"

func readOutputFile(path string) string {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	text := strings.Builder{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text.WriteString(scanner.Text())
		text.WriteString("\n")
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return text.String()
}

func loadTokens(str string) []ir.Token {

	tokens := make([]ir.Token, 0)

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			tokens = append(tokens, ir.Token{T: ir.TK_LINE_BREAK, V: "\n"})
			continue
		}
		unparsedTokens := strings.Split(line, token_seperator)
		for i, ununparsedToken := range unparsedTokens {
			name, value, found := strings.Cut(ununparsedToken, " ")
			if !found {
				panic(fmt.Sprintf("At line %d, Space seperator expected in unparsedToken %s", i, ununparsedToken))
			}
			tokenType := ir.GetTokentype(name)
			if tokenType == ir.TK_UNKNOWN {
				panic(fmt.Sprintf("At line %d, Unknown token type found in expected output %s", i, name))
			}
			tokens = append(tokens, ir.Token{T: tokenType, V: value})
		}
		tokens = append(tokens, ir.Token{T: ir.TK_LINE_BREAK, V: "\n"})
	}

	return tokens
}

type testElement struct {
	Name     string
	Value    string
	Children []*testElement
	Metadata string
}

// type parsingStackFrame struct {
// 	i             int
// 	currArr       []*testElement
// 	children      []*ir.MarkdownElement
// 	result        []*ir.MarkdownElement
// 	finished      bool
// 	returnedValue []*ir.MarkdownElement
// }

func loadElements(str string) []*ir.MarkdownElement {
	testElements := make([]*testElement, 0)

	json.Unmarshal([]byte(str), &testElements)

	return loadElement(testElements)

	// TODO: convert this recursive solution into a linear loop
	// stack := types.NewStack[parsingStackFrame](0)
	// stack.Push(parsingStackFrame{
	// 	i:        0,
	// 	currArr:  testElements,
	// 	result:   make([]*ir.MarkdownElement, 0),
	// 	children: make([]*ir.MarkdownElement, 0),
	// 	finished: false,
	// 	returnedValue: nil,
	// })
	// var frame *parsingStackFrame

	// for {
	// 	if stack.Len() == 0 {
	// 		break
	// 	}
	// 	frame, _ = stack.Top()

	// 	result := frame.result
	// 	children := frame.children
	// 	i := frame.i
	// 	currArr := frame.currArr

	// 	if currArr == nil {
	// 		return nil
	// 	}
	// 	for {
	// 		if !frame.finished && i < len(currArr) {
	// 			frame.finished = true
	// 			frame.returnedValue = result
	// 			continue
	// 		}
	// 		currEl = currArr[i]
	// 		// children = loadElement(currEl.Children)
	// 		if !frame.finished {
	// 			stack.Push(parsingStackFrame{
	// 				i:        0,
	// 				currArr:  currEl.Children,
	// 				children: make([]*ir.MarkdownElement, 0),
	// 				result:   make([]*ir.MarkdownElement, 0),
	// 				finished: false,
	// 			})
	// 			continue
	// 		}

	// 		me := ir.NewMarkDownElement(ir.ElementDefinitions[ir.TokenType(currEl.Name)], currEl.Value, children)
	// 		result = append(result, me)
	// 		i++
	// 	}
	// }
}

func loadElement(els []*testElement) []*ir.MarkdownElement {
	result := make([]*ir.MarkdownElement, 0)
	var children []*ir.MarkdownElement
	var el *testElement
	i := 0
	if els == nil {
		return nil
	}
	for {
		if i >= len(els) {
			return result
		}
		el = els[i]
		children = loadElement(el.Children)
		if definition, ok := ir.ElementDefinitions[ir.TokenType(el.Name)]; ok {
			me := ir.NewMarkDownElement(definition, el.Value, children)
			me.Metadata = el.Metadata
			result = append(result, me)
			i++
		} else {
			panic("Unknown Token " + el.Name)
		}
	}
}

func deepCompare(w *ir.MarkdownElement, g *ir.MarkdownElement) bool {

	if w == nil && g == nil {
		return true
	}
	if (w == nil && g != nil) || (w != nil && g == nil) || (w.V != g.V) || len(w.C) != len(g.C) {
		return false
	}
	if w.C == nil && g.C == nil {
		return true
	}
	if (w.C == nil && g.C != nil) || (w.C != nil && g.C == nil) {
		return false
	}
	for i := range w.C {
		a := w.C[i]
		b := g.C[i]
		if !deepCompare(a, b) {
			return false
		}
	}
	return true
}
