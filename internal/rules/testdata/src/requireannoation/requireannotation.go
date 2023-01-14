package requireannotation

import (
	"testing"
	"time"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"

	"customasserter"
)

func main() {
	t := &testing.T{}
	g := gomega.NewWithT(t)

	g.Expect("hello").To(Equal("hello")) // want `To should have an annotation`

	ExpectWithOffset(1, "hello").To(Equal("hello")) // want `To should have an annotation`

	Expect("hello").To(Equal("hello"))                      // want `To should have an annotation`
	gomega.Expect("hello").NotTo(gomega.Equal("hello"))     // want `NotTo should have an annotation`
	gomega.Expect("hello").Should(gomega.Equal("hello"))    // want `Should should have an annotation`
	gomega.Expect("hello").To(gomega.Equal("hello"))        // want `To should have an annotation`
	gomega.Expect("hello").ToNot(gomega.Equal("hello"))     // want `ToNot should have an annotation`
	gomega.Expect("hello").ShouldNot(gomega.Equal("hello")) // want `ShouldNot should have an annotation`

	Ω("hello").To(Equal("hello"))                      // want `To should have an annotation`
	gomega.Ω("hello").NotTo(gomega.Equal("hello"))     // want `NotTo should have an annotation`
	gomega.Ω("hello").Should(gomega.Equal("hello"))    // want `Should should have an annotation`
	gomega.Ω("hello").To(gomega.Equal("hello"))        // want `To should have an annotation`
	gomega.Ω("hello").ToNot(gomega.Equal("hello"))     // want `ToNot should have an annotation`
	gomega.Ω("hello").ShouldNot(gomega.Equal("hello")) // want `ShouldNot should have an annotation`

	Eventually("hello").Should(Equal("hello"))                          // want `Should should have an annotation`
	Eventually("hello").WithPolling(time.Minute).Should(Equal("hello")) // want `Should should have an annotation`
	gomega.Eventually("hello").Should(gomega.Equal("hello"))            // want `Should should have an annotation`
	gomega.Eventually("hello").ShouldNot(gomega.Equal("hello"))         // want `ShouldNot should have an annotation`

	Consistently("hello").Should(Equal("hello"))                  // want `Should should have an annotation`
	gomega.Consistently("hello").Should(gomega.Equal("hello"))    // want `Should should have an annotation`
	gomega.Consistently("hello").ShouldNot(gomega.Equal("hello")) // want `ShouldNot should have an annotation`

	ExpectWithOffset(1, "hello").To(Equal("hello"), "hello")

	Expect("hello").To(Equal("hello"), "hello")
	gomega.Expect("hello").NotTo(gomega.Equal("hello"), "hello")
	gomega.Expect("hello").Should(gomega.Equal("hello"), "hello")
	gomega.Expect("hello").ShouldNot(gomega.Equal("hello"), "hello")
	gomega.Expect("hello").To(gomega.Equal("hello"), "hello")
	gomega.Expect("hello").ToNot(gomega.Equal("hello"), "hello")

	g.Expect("hello").To(Equal("hello"), "hello")

	Ω("hello").To(Equal("hello"), "hello")
	gomega.Ω("hello").NotTo(gomega.Equal("hello"), "hello")
	gomega.Ω("hello").Should(gomega.Equal("hello"), "hello")
	gomega.Ω("hello").ShouldNot(gomega.Equal("hello"), "hello")
	gomega.Ω("hello").To(gomega.Equal("hello"), "hello")
	gomega.Ω("hello").ToNot(gomega.Equal("hello"), "hello")

	Eventually("hello").Should(Equal("hello"), "hello")
	gomega.Eventually("hello").Should(gomega.Equal("hello"), "hello")
	gomega.Eventually("hello").ShouldNot(gomega.Equal("hello"), "hello")

	Consistently("hello").Should(Equal("hello"), "hello")
	gomega.Consistently("hello").Should(gomega.Equal("hello"), "hello")
	gomega.Consistently("hello").ShouldNot(gomega.Equal("hello"), "hello")

	customasserter.Expect().To("hey")

	Expect("hello").Error()
}
