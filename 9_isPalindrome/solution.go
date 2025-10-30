package ispalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	digits := make([]int, 10)
	count := 0
	m := 1
	remaining := x

	for remaining > 0 {
		v := remaining % (m * 10)
		digits[count] = v / m
		remaining -= v
		m *= 10
		count++
	}

	mid := count / 2

	for i := 0; i < mid; i++ {
		if digits[i] != digits[count-1-i] {
			return false
		}
	}

	return true
}
