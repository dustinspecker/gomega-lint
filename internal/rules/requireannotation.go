package rules

import (
	"fmt"

	"github.com/dustinspecker/gomega-lint/internal/visitor"
	"golang.org/x/tools/go/analysis"
)

var RequireAnnotationAnalyzer = &analysis.Analyzer{
	Name: "requiredescription",
	Doc:  "Require all assertions to have a description",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	visitor.VistGomegaAssertions(pass, func(gomegaAssertion visitor.GomegaAssertion) *analysis.Diagnostic {
		if len(gomegaAssertion.AssertionArgs) == 1 {
			return &analysis.Diagnostic{
				Message: fmt.Sprintf("%s should have an annotation", gomegaAssertion.AssertionMethodName),
				Pos:     gomegaAssertion.Node.Pos(),
			}
		}

		return nil
	})

	return nil, nil
}
