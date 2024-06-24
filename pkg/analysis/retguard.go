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
			for _, result := range fn.Type.Results.List {
				if len(result.Names) == 0 {
					continue
				}
				for _, name := range result.Names {
					if !isAssigned(fn.Body, name.Name) && isReturned(fn.Body, name.Name) {
						pass.Reportf(name.Pos(), "named return value %s is never assigned", name.Name)
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func isAssigned(block *ast.BlockStmt, name string) bool {
	assigned := false
	isShadowed := false

	ast.Inspect(block, func(n ast.Node) bool {
		// skip the block if the variable was shadowed
		// after shadowing, even if there is an assignment involving a variable named "name", the first variable named "name" will never be assigned
		if isShadowed {
			return false
		}

		switch n := n.(type) {

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
							isShadowed = true
							return false
						}
					}
				}
			}
		// find out if "name" variable was assigned
		case *ast.AssignStmt:
			for _, lhs := range n.Lhs {
				switch lhs := lhs.(type) {
				case *ast.Ident:
					if lhs.Name == name {
						if n.Tok == token.DEFINE {
							isShadowed = true
							return false
						}
						assigned = true
						return false
					}
				case *ast.SelectorExpr:
					if ident, ok := lhs.X.(*ast.Ident); ok && ident.Name == name {
						assigned = true
						return false
					}
				}
			}

		// find out if "name" was passed as reference
		case *ast.CallExpr:
			for _, arg := range n.Args {
				switch arg := arg.(type) {
				case *ast.UnaryExpr:
					if arg.Op == token.AND {
						if ident, ok := arg.X.(*ast.Ident); ok && ident.Name == name {
							assigned = true
							return false
						}
					}
				}
			}
		}
		return true
	})

	return assigned
}

func isReturned(block *ast.BlockStmt, name string) bool {
	isShadowed := false
	isReturned := false

	ast.Inspect(block, func(n ast.Node) bool {
		if isShadowed {
			return false
		}

		switch n := n.(type) {

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
							isShadowed = true
							return false
						}
					}
				}
			}
		case *ast.AssignStmt:
			// find out if "name" variable was shadowed
			for _, lhs := range n.Lhs {
				switch lhs := lhs.(type) {
				case *ast.Ident:
					if lhs.Name == name && n.Tok == token.DEFINE {
						isShadowed = true
						return false
					}
				}
			}
		case *ast.ReturnStmt:
			if len(n.Results) == 0 {
				isReturned = true
				return false
			}
			for _, result := range n.Results {
				if ident, ok := result.(*ast.Ident); ok && ident.Name == name {
					isReturned = true
					return false
				}
			}
		}

		return true
	})

	return isReturned
}
