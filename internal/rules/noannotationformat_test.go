package rules_test

import (
	"testing"

	"github.com/dustinspecker/gomega-lint/internal/rules"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNoAnnotationFormat(t *testing.T) {
	testDataDir := analysistest.TestData()

	analysistest.Run(t, testDataDir, rules.NoAnnotationFormat, "./src/noannotationformat/")
}
