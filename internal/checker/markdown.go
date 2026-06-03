package checker

import (
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

func ExtractMarkdownLinks(content []byte) []string {
	md := goldmark.New()
	doc := md.Parser().Parse(text.NewReader(content))
	var links []string
	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if entering {
			if link, ok := n.(*ast.Link); ok {
				links = append(links, string(link.Destination))
			}
		}
		return ast.WalkContinue, nil
	})
	return links
}
