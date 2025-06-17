package analys

import (
//	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
//	"strings"
)

type AnalysResult struct {
	DeclCount    int
	CallCount    int
	AssignCount  int
	ImportsCount int
}

func Analys(filePath string) (AnalysResult, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return AnalysResult{}, err
	}

	result := AnalysResult{}

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			if x.Tok == token.VAR || x.Tok == token.CONST || x.Tok == token.TYPE {
				result.DeclCount += len(x.Specs)
			}

		case *ast.AssignStmt:
			if x.Tok == token.ASSIGN {
				result.AssignCount++
			}

		case *ast.CallExpr:
			result.CallCount++

		case *ast.ImportSpec:
			result.ImportsCount++
		}
		return true
	})

	return result, nil
}

func isBuiltin(name string) bool {
	builtins := []string{
		"append", "cap", "close", "complex", "copy", "delete", "imag", "len",
		"make", "new", "panic", "print", "println", "real", "recover",
	}
	for _, b := range builtins {
		if b == name {
			return true
		}
	}
	return false
}
