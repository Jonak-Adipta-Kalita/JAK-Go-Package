package jak_go_package

type Mathematics struct {
	Number1 int
	Number2 int
}

func (x Mathematics) Add() int {
	return x.Number1 + x.Number2
}

func (x Mathematics) Sub() int {
	return x.Number1 - x.Number2
}

func (x Mathematics) Mul() int {
	return x.Number1 * x.Number2
}

func (x Mathematics) Div() int {
	return x.Number1 / x.Number2
}
