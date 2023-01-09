package requireannotation

import (
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func main() {
	value := "hello"
	valuePtr := &value

	Expect(*valuePtr).To(Equal("hello"))               // want `use PointTo instead of dereferencing`
	gomega.Expect(*valuePtr).To(gomega.Equal("hello")) // want `use PointTo instead of dereferencing`

	Expect(valuePtr).To(PointTo(Equal("hello")))
	gomega.Expect(valuePtr).To(PointTo(gomega.Equal("hello")))
}
