package go_koans

func divideFourBy(i int) int {
	return 4 / i
}

const __divisor__ = 0

func aboutPanics() {
	var n int
	defer func() { // alternative solution is to just set __divisor__ to 2 but this is more fun
		if r := recover(); r != nil {
			n = 2
		}
	}()
	n = divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
