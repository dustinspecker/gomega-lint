package usenewwitht

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"

	"customasserter"
)

func main() {
	t := &testing.T{}

	gomega.NewGomegaWithT(t) // want `NewGomegaWithT is deprecated. Use NewWithT instead.`
	NewGomegaWithT(t)        // want `NewGomegaWithT is deprecated. Use NewWithT instead.`

	gomega.NewWithT(t)
	NewWithT(t)

	customasserter.NewGomegaWithT(t)
}
