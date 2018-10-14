package nodeparser

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(fileName string) (*ast.File, error) {
	return parser.ParseFile(token.NewFileSet(), fileName, nil, parser.AllErrors)
}

func ParseDecl(declStr string) (*ast.Decl, error) {
	header := `
package main
	`

	f, err := parser.ParseFile(token.NewFileSet(), "main.go", header+declStr, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		return &decl, nil
	}

	return nil, nil
}
