package rules

import (
	"go/ast"

	"github.com/dustinspecker/gomega-lint/internal/visitor"
	"golang.org/x/tools/go/analysis"
)

var UsePointTo = &analysis.Analyzer{
	Name: "usepointto",
	Doc:  "use PointTo instead of dereferencing",
	Run:  runUsePointTo,
}

func runUsePointTo(pass *analysis.Pass) (interface{}, error) {
	visitor.VistGomegaAssertions(pass, func(gomegaAssertion visitor.GomegaAssertion) *analysis.Diagnostic {
		_, isStarExpr := gomegaAssertion.Subject.(*ast.StarExpr)
		if isStarExpr {
			return &analysis.Diagnostic{
				Message: "use PointTo instead of dereferencing",
				Pos:     gomegaAssertion.Node.Pos(),
			}
		}

		return nil
	})

	return nil, nil
}
