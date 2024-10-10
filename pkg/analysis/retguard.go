package analysis

import (
	"go/ast"
	"go/token"

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

			// todo: squash returns values
			for _, result := range fn.Type.Results.List {
				if len(result.Names) == 0 {
					continue
				}
				for _, name := range result.Names {
					if pos := firstReturnWithoutAssignment(fn.Body, name.Name); pos != nil {
						pass.Reportf(*pos, "named return value %s is never assigned", name.Name)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func firstReturnWithoutAssignment(block *ast.BlockStmt, name string) *token.Pos {
	var pos *token.Pos

	var assignments, shadows []*ast.BlockStmt

	isInside := func(pos ast.Node, zone []*ast.BlockStmt) bool {
		for _, s := range zone {
			if pos.Pos() >= s.Lbrace && pos.End() <= s.Rbrace {
				return true
			}
		}
		return false
	}

	var curBlock *ast.BlockStmt
	ast.Inspect(block, func(n ast.Node) bool {

		switch n := n.(type) {

		// todo: check me
		case *ast.BlockStmt:
			curBlock = n
		case *ast.DeclStmt:
			// find out if "name" variable was shadowed
			genDecl, ok := n.Decl.(*ast.GenDecl)
			if !ok {
				return true
			}
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					for _, ident := range valueSpec.Names {
						if ident.Name == name {
							shadows = append(shadows, curBlock)
							return false
						}
					}
				}
			}
		case *ast.IncDecStmt:
			if ident, ok := n.X.(*ast.Ident); ok && ident.Name == name {
				assignments = append(assignments, curBlock)
				return false
			}
		case *ast.AssignStmt:

			// find out if "name" variable was shadowed
			for _, lhs := range n.Lhs {
				switch lhs := lhs.(type) {
				case *ast.Ident:
					if lhs.Name == name {
						if n.Tok == token.ASSIGN {
							assignments = append(assignments, curBlock)
						}
						if n.Tok == token.DEFINE {
							shadows = append(shadows, curBlock)
						}
					}
				case *ast.SelectorExpr:
					if ident, ok := lhs.X.(*ast.Ident); ok && ident.Name == name {
						assignments = append(assignments, curBlock)
					}
				case *ast.IndexExpr:
					if xIdent, ok := lhs.X.(*ast.Ident); ok && xIdent.Name == name {
						assignments = append(assignments, curBlock)
					}
				}
				return false
			}

		case *ast.ReturnStmt:
			if isInside(n, shadows) {
				return false
			}
			if isInside(n, assignments) {
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
			for _, arg := range n.Args {
				switch arg := arg.(type) {
				case *ast.UnaryExpr:
					if arg.Op == token.AND {
						if ident, ok := arg.X.(*ast.Ident); ok && ident.Name == name {
							assignments = append(assignments, curBlock)
							return false
						}
					}
				case *ast.SelectorExpr:
					if xIdent, ok := arg.X.(*ast.Ident); ok && xIdent.Name == name {
						assignments = append(assignments, curBlock)
						return false
					}
				}
			}

		}

		return true
	})

	return pos
}
