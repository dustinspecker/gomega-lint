package rules_test

import (
	"testing"

	"github.com/dustinspecker/gomega-lint/internal/rules"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestUseNewWithT(t *testing.T) {
	testDataDir := analysistest.TestData()

	analysistest.Run(t, testDataDir, rules.UseNewWithT, "./src/usenewwitht/")
}

func TestUseNewWithTAutoFix(t *testing.T) {
	testDataDir := analysistest.TestData()

	analysistest.RunWithSuggestedFixes(t, testDataDir, rules.UseNewWithT, "./src/usenewwitht/")
}
