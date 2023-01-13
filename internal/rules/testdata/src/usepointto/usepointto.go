package requireannotation

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func main() {
	t := &testing.T{}
	g := gomega.NewWithT(t)

	value := "hello"
	valuePtr := &value

	g.Expect(*valuePtr).To(Equal("hello")) // want `use PointTo instead of dereferencing`

	ExpectWithOffset(1, *valuePtr).To(Equal("hello")) // want `use PointTo instead of dereferencing`

	Expect(*valuePtr).To(Equal("hello"))               // want `use PointTo instead of dereferencing`
	gomega.Expect(*valuePtr).To(gomega.Equal("hello")) // want `use PointTo instead of dereferencing`

	g.Expect(valuePtr).To(PointTo(Equal("hello")))

	ExpectWithOffset(1, valuePtr).To(PointTo(Equal("hello")))

	Expect(valuePtr).To(PointTo(Equal("hello")))
	gomega.Expect(valuePtr).To(PointTo(gomega.Equal("hello")))
}
