package longestpalindrome

func longestPalindrome(s string) string {
	L := len(s)

	if L < 2 {
		return s
	}

	i := 0
	j := L
	start := 0
	end := 0

	for i < j && i < L {
		switch {
		case isPalindrome(s, i, j):
			// update start/end
			if end-start < j-i {
				start = i
				end = j
			}
			fallthrough
		case j-i < end-start:
			// stop current i, start next one
			i++
			j = L
		default:
			// check next j for current i
			j--
		}
	}

	return s[start:end]
}

func isPalindrome(s string, start int, end int) bool {
	i := start
	j := end

	for i < j {
		if s[i] != s[j-1] {
			return false
		}

		i++
		j--
	}

	return true
}
