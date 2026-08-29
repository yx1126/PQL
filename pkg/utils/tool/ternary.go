package tool

func Flag[T any](flag bool, trueValue, falseValue T) T {
	if flag {
		return trueValue
	}
	return falseValue
}

func FlagFn[T any](flagFn func(trueValue, falseValue T) bool, trueValue, falseValue T) T {
	if flagFn(trueValue, falseValue) {
		return trueValue
	}
	return falseValue
}
