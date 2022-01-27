package analyzer

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "addresslowercase",
	Doc:      "Checks that all addresses are strictly lowercase",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.BlockStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		blockStmt := node.(*ast.BlockStmt)

		if len(blockStmt.List) == 0 {
			return
		}

		for _, stmt := range blockStmt.List {

			// Check in constants
			if declStmt, ok := stmt.(*ast.DeclStmt); ok {
				if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok {
					variable, ok := genDecl.Specs[0].(*ast.ValueSpec); if !ok {
						continue
					}

					if len(variable.Values) == 0 {
						continue
					}

					variableValues, ok := variable.Values[0].(*ast.BasicLit); if !ok {
						continue
					}

					if variableValues.Kind == token.STRING {
						if !validateString(variableValues.Value) {
							reportFailure(pass, variableValues.Pos(), variableValues.Value)
						}
						continue
					}
				}
			}

			// Check in assignments
			if assignStmt, ok := stmt.(*ast.AssignStmt); ok {
				if len(assignStmt.Rhs) == 0 {
					continue
				}
				for _, expr := range assignStmt.Rhs {

					if basicLit, ok := expr.(*ast.BasicLit); ok {
						if basicLit.Kind == token.STRING {
							if !validateString(basicLit.Value) {
								reportFailure(pass, basicLit.Pos(), basicLit.Value)
							}
							continue
						}
					}
				}
			}
			continue
		}
		return
	})

	return nil, nil
}

func reportFailure(pass *analysis.Pass, position token.Pos, value string) {
	pass.Reportf(position, "string contains address with capital letters: %s", value)
}
