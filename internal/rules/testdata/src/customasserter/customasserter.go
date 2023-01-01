package customasserter

import "testing"

type assert struct{}

func Expect() assert {
	return assert{}
}

func (a assert) To(_ string) {
}

func NewGomegaWithT(t *testing.T) {
}
