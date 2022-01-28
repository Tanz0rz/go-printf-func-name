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
	Doc:      "Checks that all addresses in strings are strictly lowercase",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.File)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		file := node.(*ast.File)

		for _, decl := range file.Decls {

			// Check for any package-level strings
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				checkGenDecl(pass, genDecl)
				continue
			}

			// Check inside all functions
			if funcDecl, ok := decl.(*ast.FuncDecl); ok {
				for _, stmt := range funcDecl.Body.List {

					// Check in constants
					if declStmt, ok := stmt.(*ast.DeclStmt); ok {
						if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok {
							checkGenDecl(pass, genDecl)
							continue
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
			}
		}
		return
	})

	return nil, nil
}

func reportFailure(pass *analysis.Pass, position token.Pos, value string) {
	pass.Reportf(position, "string contains address with capital letters: %s", value)
}

func checkGenDecl(pass *analysis.Pass, genDecl *ast.GenDecl) {
	variable, ok := genDecl.Specs[0].(*ast.ValueSpec); if !ok {
		return
	}

	if len(variable.Values) == 0 {
		return
	}

	basicLit, ok := variable.Values[0].(*ast.BasicLit); if !ok {
		return
	}

	if basicLit.Kind == token.STRING {
		if !validateString(basicLit.Value) {
			reportFailure(pass, basicLit.Pos(), basicLit.Value)
		}
		return
	}
}
