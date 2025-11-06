package isvalid

var openLookup = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	open := []rune{}

	for _, c := range s {
		expectedOpenPair, found := openLookup[c]

		if !found {
			// opening bracket
			open = append(open, c)
			continue
		}

		l := len(open)

		if l > 0 && expectedOpenPair == open[l-1] {
			// valid closing bracket
			open = open[:l-1]
		} else {
			// invalid closing bracket
			return false
		}
	}

	return len(open) == 0
}
