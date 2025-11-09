package converters

import "github.com/ManasRaut/markdown_compiler/ir"

// Convert MarkdownElements into final UI elements
type Converter[R any] interface {
	Convert(e []*ir.MarkdownElement) (*R, error)
}
