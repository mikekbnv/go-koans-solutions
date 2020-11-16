package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const divisor = 2

func aboutPanics() {

	n := divideFourBy(divisor)

	assert(n == 2) // panics are exceptional errors at runtime
}
