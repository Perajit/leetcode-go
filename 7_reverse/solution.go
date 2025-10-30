package reverse

const Min = -(1 << 31)
const Max = (1 << 31) - 1

func reverse(x int) int {
	abs := absolute(x)

	if abs < 10 {
		return x
	}

	digits := make([]int, 10) // max at 2^30
	remaining := abs
	m := 1
	count := 0

	for remaining > 0 {
		v := remaining % (m * 10)
		digits[count] = v / m
		remaining -= v
		m *= 10
		count++
	}

	m = 1
	for i := count - 1; i >= 0; i-- {
		remaining += digits[i] * m
		m *= 10
	}

	if x < 0 {
		remaining = -remaining
	}

	if remaining < Min || remaining > Max {
		return 0
	}

	return remaining
}

func absolute(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
