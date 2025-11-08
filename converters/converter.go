package converters

import "github.com/ManasRaut/md_lex/ir"

// Convert MarkdownElements into final UI elements
type Converter[R any] interface {
	Convert(e []*ir.MarkdownElement) (*R, error)
}
