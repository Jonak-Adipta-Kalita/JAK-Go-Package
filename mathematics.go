package jak_go_package

type Mathematics struct {
	number1 int
	number2 int
}

func (x Mathematics) add() int {
	return x.number1 + x.number2
}

func (x Mathematics) sub() int {
	return x.number1 - x.number2
}

func (x Mathematics) mul() int {
	return x.number1 * x.number2
}

func (x Mathematics) div() int {
	return x.number1 / x.number2
}
