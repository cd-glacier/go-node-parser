package nodeparser

import (
	"errors"
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

func ParseStmt(stmtStr string) (*ast.Stmt, error) {
	header := `
package main

func main() {
	`

	footer := `
}
	`

	f, err := parser.ParseFile(token.NewFileSet(), "main.go", header+stmtStr+footer, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		d, ok := decl.(*ast.FuncDecl)
		if !ok {
			return nil, errors.New("Failed to convert to *ast.FuncDecl")
		}

		if len(d.Body.List) != 1 {
			return nil, errors.New("stmt len is not 1")
		}
		return &d.Body.List[0], nil
	}

	return nil, nil
}
