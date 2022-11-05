package main

import (
	"github.com/dustinspecker/gomega-lint/internal/rules"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(rules.RequireAnnotationAnalyzer)
}
