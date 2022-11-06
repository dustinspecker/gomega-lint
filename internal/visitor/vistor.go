package visitor

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/tools/go/analysis"
)

type GomegaAssertion struct {
	AssertionArgs       []ast.Expr
	AssertionMethodName string
	Node                ast.Node
}

var assertionMethods = []string{
	"NotTo",
	"Should",
	"ShouldNot",
	"To",
	"ToNot",
}

func VistGomegaAssertions(pass *analysis.Pass, visitFunc func(gomegaAssertion GomegaAssertion) *analysis.Diagnostic) {
	var consistentlyFuncType types.Type
	var eventuallyFuncType types.Type
	var expectFuncType types.Type
	var omegaFuncType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		// use HasSuffix because unit tests will use package vendor/github.com/onsi/gomega
		if strings.HasSuffix(pkg.Path(), "github.com/onsi/gomega") {
			consistentlyFuncType = pkg.Scope().Lookup("Consistently").Type()
			eventuallyFuncType = pkg.Scope().Lookup("Eventually").Type()
			expectFuncType = pkg.Scope().Lookup("Expect").Type()
			omegaFuncType = pkg.Scope().Lookup("Ω").Type()
		}
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			// examine every function call
			callExpr, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Look for cases like Expect(...).To(...)
			selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			// The subject of the call should be Expect(...), etc.
			subjectCallExpr, ok := selExpr.X.(*ast.CallExpr)
			if !ok {
				return true
			}

			// Verify the subject is one of Expect, etc.
			subjectCallExprFunType := pass.TypesInfo.TypeOf(subjectCallExpr.Fun)
			if subjectCallExprFunType != consistentlyFuncType && subjectCallExprFunType != eventuallyFuncType && subjectCallExprFunType != expectFuncType && subjectCallExprFunType != omegaFuncType {
				return true
			}

			// Verify the the method being called is an assertion method
			if slices.Contains(assertionMethods, selExpr.Sel.Name) {
				gomegaAssertion := GomegaAssertion{
					AssertionArgs:       callExpr.Args,
					AssertionMethodName: selExpr.Sel.Name,
					Node:                node,
				}

				diagnostic := visitFunc(gomegaAssertion)
				if diagnostic != nil {
					pass.Report(*diagnostic)
				}
			}

			return true
		})
	}
}
