package rules

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const (
	newGomegaWithTFuncName = "NewGomegaWithT"
)

var UseNewWithT = &analysis.Analyzer{
	Name: "usenewwitht",
	Run:  useNewWithTRun,
	Doc:  "Use NewWithT instead of the deprecated NewGomegaWithT",
}

func useNewWithTRun(pass *analysis.Pass) (interface{}, error) {
	var newGomegaWithT types.Type

	for _, pkg := range pass.Pkg.Imports() {
		// use HasSuffix because unit tests will use package vendor/github.com/onsi/gomega
		if strings.HasSuffix(pkg.Path(), "github.com/onsi/gomega") {
			newGomegaWithT = pkg.Scope().Lookup(newGomegaWithTFuncName).Type()
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// examine every function call
			callExpr, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			// cases like NewGomegaWithT and NewWithT
			callExprIdent, isCallExprIdent := callExpr.Fun.(*ast.Ident)

			// cases like gomega.NewGomegaWithT and gomega.NewWithT
			callExprSelector, isCallExprSelector := callExpr.Fun.(*ast.SelectorExpr)

			if !isCallExprIdent && !isCallExprSelector {
				return true
			}

			// validate function being examined is from gomega package
			callExprFunType := pass.TypesInfo.TypeOf(callExpr.Fun)
			if callExprFunType != newGomegaWithT {
				return true
			}

			var name string
			if isCallExprIdent {
				name = callExprIdent.Name
			} else {
				name = callExprSelector.Sel.Name
			}

			// Unfortunately have to look up name because NewGomegaWithT points at NewWithT, so they
			// are the same thing when parsed.
			if name == newGomegaWithTFuncName {
				var start token.Pos

				// get starting position for text replacement
				// want to keep import name intact if using named import
				if isCallExprIdent {
					start = callExpr.Pos()
				} else {
					start = callExprSelector.Sel.NamePos
				}

				pass.Report(analysis.Diagnostic{
					Message: "NewGomegaWithT is deprecated. Use NewWithT instead.",
					Pos:     node.Pos(),
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "Replace NewGomegaWithT with NewWithT",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     start,
									End:     callExpr.Lparen,
									NewText: []byte("NewWithT"),
								},
							},
						},
					},
				})
			}

			return true
		})
	}

	return nil, nil
}
