package visitor

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type GomegaAssertion struct {
	AssertionArgs       []ast.Expr
	AssertionMethodName string
	Node                ast.Node
	Subject             ast.Expr
}

func VistGomegaAssertions(pass *analysis.Pass, visitFunc func(gomegaAssertion GomegaAssertion) *analysis.Diagnostic) {
	var assertionType types.Type
	var asyncAssertionType types.Type

	for _, pkg := range pass.Pkg.Imports() {
		// use HasSuffix because unit tests will use package vendor/github.com/onsi/gomega
		if strings.HasSuffix(pkg.Path(), "github.com/onsi/gomega") {
			assertionType = pkg.Scope().Lookup("Assertion").Type()
			asyncAssertionType = pkg.Scope().Lookup("AsyncAssertion").Type()
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
			signature, ok := subjectCallExprFunType.(*types.Signature)
			if !ok {
				return true
			}
			if signature.Results().Len() == 0 {
				return true
			}

			returnVal := signature.Results().At(0).Type()
			if returnVal != assertionType && returnVal != asyncAssertionType {
				return true
			}

			// we want the last argument in cases of `ExpectWithOffset`, etc. where the last
			// arg is the subject under test
			lastArg := subjectCallExpr.Args[len(subjectCallExpr.Args)-1]

			gomegaAssertion := GomegaAssertion{
				AssertionArgs:       callExpr.Args,
				AssertionMethodName: selExpr.Sel.Name,
				Node:                node,
				Subject:             lastArg,
			}

			diagnostic := visitFunc(gomegaAssertion)
			if diagnostic != nil {
				pass.Report(*diagnostic)
			}

			return true
		})
	}
}
