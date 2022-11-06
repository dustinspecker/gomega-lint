package requireannotation

import (
	"fmt"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func main() {
	Expect("hello").To(Equal("hello"), fmt.Sprintf("hello %s", "dustin"))                      // want `Annotation should not use fmt.Sprintf`
	gomega.Expect("hello").NotTo(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))     // want `Annotation should not use fmt.Sprintf`
	gomega.Expect("hello").Should(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))    // want `Annotation should not use fmt.Sprintf`
	gomega.Expect("hello").ShouldNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin")) // want `Annotation should not use fmt.Sprintf`
	gomega.Expect("hello").To(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))        // want `Annotation should not use fmt.Sprintf`
	gomega.Expect("hello").ToNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))     // want `Annotation should not use fmt.Sprintf`

	Ω("hello").To(Equal("hello"), fmt.Sprintf("hello %s", "dustin"))                      // want `Annotation should not use fmt.Sprintf`
	gomega.Ω("hello").NotTo(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))     // want `Annotation should not use fmt.Sprintf`
	gomega.Ω("hello").Should(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))    // want `Annotation should not use fmt.Sprintf`
	gomega.Ω("hello").ShouldNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin")) // want `Annotation should not use fmt.Sprintf`
	gomega.Ω("hello").To(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))        // want `Annotation should not use fmt.Sprintf`
	gomega.Ω("hello").ToNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))     // want `Annotation should not use fmt.Sprintf`

	Eventually("hello").Should(Equal("hello"), fmt.Sprintf("hello %s", "dustin"))                  // want `Annotation should not use fmt.Sprintf`
	gomega.Eventually("hello").Should(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))    // want `Annotation should not use fmt.Sprintf`
	gomega.Eventually("hello").ShouldNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin")) // want `Annotation should not use fmt.Sprintf`

	Consistently("hello").Should(Equal("hello"), fmt.Sprintf("hello %s", "dustin"))                  // want `Annotation should not use fmt.Sprintf`
	gomega.Consistently("hello").Should(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin"))    // want `Annotation should not use fmt.Sprintf`
	gomega.Consistently("hello").ShouldNot(gomega.Equal("hello"), fmt.Sprintf("hello %s", "dustin")) // want `Annotation should not use fmt.Sprintf`

	Expect("hello").To(Equal("hello"), "hello %s", "dustin")
	gomega.Expect("hello").NotTo(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Expect("hello").Should(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Expect("hello").ShouldNot(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Expect("hello").To(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Expect("hello").ToNot(gomega.Equal("hello"), "hello %s", "dustin")

	Ω("hello").To(Equal("hello"), "hello %s", "dustin")
	gomega.Ω("hello").NotTo(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Ω("hello").Should(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Ω("hello").ShouldNot(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Ω("hello").To(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Ω("hello").ToNot(gomega.Equal("hello"), "hello %s", "dustin")

	Eventually("hello").Should(Equal("hello"), "hello %s", "dustin")
	gomega.Eventually("hello").Should(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Eventually("hello").ShouldNot(gomega.Equal("hello"), "hello %s", "dustin")

	Consistently("hello").Should(Equal("hello"), "hello %s", "dustin")
	gomega.Consistently("hello").Should(gomega.Equal("hello"), "hello %s", "dustin")
	gomega.Consistently("hello").ShouldNot(gomega.Equal("hello"), "hello %s", "dustin")
}
