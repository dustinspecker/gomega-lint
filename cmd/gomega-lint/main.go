package main

import (
	"github.com/dustinspecker/gomega-lint/internal/rules"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(rules.RequireAnnotationAnalyzer, rules.NoAnnotationFormat)
}
