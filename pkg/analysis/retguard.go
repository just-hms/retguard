package analysis

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/just-hms/retguard/pkg/analysis/astx"
	"golang.org/x/tools/go/analysis"
)

var RetGuard = &analysis.Analyzer{
	Name: "retguard",
	Doc:  "check for unassigned named return values",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			if fn.Type.Results == nil {
				return true
			}

			// should check the assignment only if the error is nil or false

			// squash returns values
			toReport := map[token.Pos][]string{}
			for _, result := range fn.Type.Results.List {
				for _, name := range result.Names {
					if pos := firstReturnWithoutAssignment(fn.Body, name.Name); pos != nil {
						toReport[*pos] = append(toReport[*pos], name.Name)
					}
				}
			}

			for pos, r := range toReport {
				if len(r) == 1 {
					pass.Reportf(pos, "named return value %s is never assigned", r[0])
				} else {
					pass.Reportf(pos, "named return values %s are never assigned", strings.Join(r, ", "))
				}
			}

			return true
		})
	}
	return nil, nil
}

func firstReturnWithoutAssignment(block *ast.BlockStmt, name string) *token.Pos {

	var blocks astx.Blocks
	var pos *token.Pos

	if name == "e" {
		fmt.Println()
	}

	ast.Inspect(block, func(n ast.Node) bool {

		switch n := n.(type) {

		case *ast.BlockStmt:
			blocks.Insert(n)
		case *ast.DeclStmt:
			// find out if "name" variable was shadowed
			genDecl, ok := n.Decl.(*ast.GenDecl)
			if !ok {
				return true
			}
			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				for _, ident := range valueSpec.Names {
					if ident.Name != name {
						continue
					}
					blocks.Shadow(n)
					return false
				}
			}
		case *ast.IncDecStmt:
			if ident, ok := n.X.(*ast.Ident); ok && ident.Name == name {
				blocks.Assign(n)
				return false
			}
		case *ast.AssignStmt:

			// find out if "name" variable was shadowed
			for _, lhs := range n.Lhs {
				switch lhs := lhs.(type) {
				case *ast.Ident:
					if lhs.Name == name {
						if n.Tok == token.ASSIGN {
							blocks.Assign(n)
						}
						if n.Tok == token.DEFINE {
							blocks.Shadow(n)
						}
					}
				case *ast.SelectorExpr:
					if ident, ok := lhs.X.(*ast.Ident); ok && ident.Name == name {
						blocks.Assign(n)
					}
				case *ast.IndexExpr:
					if xIdent, ok := lhs.X.(*ast.Ident); ok && xIdent.Name == name {
						blocks.Assign(n)
					}
				}
			}
			return false

		case *ast.ReturnStmt:
			if blocks.GetState(n) != astx.NOTHING {
				return false
			}
			for _, result := range n.Results {
				if ident, ok := result.(*ast.Ident); ok && ident.Name == name {
					pos = &n.Return
					return false
				}
			}
			if len(n.Results) == 0 {
				pos = &n.Return
				return false
			}
			return false
		case *ast.CallExpr:
			if selExpr, ok := n.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == name {
					blocks.Assign(n)
					return false
				}
			}

			for _, arg := range n.Args {
				if ident, ok := arg.(*ast.Ident); ok && ident.Name == name {
					blocks.Assign(n)
					return false
				}

				switch arg := arg.(type) {
				case *ast.UnaryExpr:
					if arg.Op == token.AND {
						if ident, ok := arg.X.(*ast.Ident); ok && ident.Name == name {
							blocks.Assign(n)
							return false
						}
					}
				case *ast.SelectorExpr:
					if xIdent, ok := arg.X.(*ast.Ident); ok && xIdent.Name == name {
						blocks.Assign(n)
						return false
					}
				}

			}

		}

		return true
	})

	return pos
}
