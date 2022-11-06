package rules

import (
	"go/ast"
	"go/types"

	"github.com/dustinspecker/gomega-lint/internal/visitor"
	"golang.org/x/tools/go/analysis"
)

var NoAnnotationFormat = &analysis.Analyzer{
	Name: "noannotationformat",
	Doc:  "Avoid fmt.Sprintf in annotations",
	Run:  noAnnotationFormatRun,
}

func noAnnotationFormatRun(pass *analysis.Pass) (interface{}, error) {
	var fmtSprintFyncType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() == "fmt" {
			fmtSprintFyncType = pkg.Scope().Lookup("Sprintf").Type()
		}
	}

	visitor.VistGomegaAssertions(pass, func(gomegaAssertion visitor.GomegaAssertion) *analysis.Diagnostic {
		if len(gomegaAssertion.AssertionArgs) > 1 {
			callExpr, ok := gomegaAssertion.AssertionArgs[1].(*ast.CallExpr)
			if !ok {
				return nil
			}

			if fmtSprintFyncType != pass.TypesInfo.TypeOf(callExpr.Fun) {
				return nil
			}

			return &analysis.Diagnostic{
				Message: "Annotation should not use fmt.Sprintf",
				Pos:     gomegaAssertion.Node.Pos(),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						TextEdits: []analysis.TextEdit{
							{
								Pos:     gomegaAssertion.AssertionArgs[1].Pos(),
								End:     callExpr.Lparen + 1,
								NewText: []byte(""),
							},
							{
								Pos:     callExpr.Rparen,
								End:     callExpr.Rparen + 1,
								NewText: []byte(""),
							},
						},
					},
				},
			}
		}

		return nil
	})
	return nil, nil
}
