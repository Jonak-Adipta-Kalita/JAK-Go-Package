package jak_go_package

type mathematics struct {
	number1 int
	number2 int
}

func (x mathematics) add() int {
	return x.number1 + x.number2
}

func (x mathematics) sub() int {
	return x.number1 - x.number2
}

func (x mathematics) mul() int {
	return x.number1 * x.number2
}

func (x mathematics) div() int {
	return x.number1 / x.number2
}
