package customasserter

type assert struct{}

func Expect() assert {
	return assert{}
}

func (a assert) To(_ string) {
}
