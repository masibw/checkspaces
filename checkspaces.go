package checkspaces

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "checkspaces is a checker for go:embed comment."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "checkspaces",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Comment)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Comment:
			if strings.Contains(n.Text, "// go:embed") {
				pass.Reportf(n.Pos(), "There is a space between slash and go:embed")
			}
		}
	})

	return nil, nil
}
