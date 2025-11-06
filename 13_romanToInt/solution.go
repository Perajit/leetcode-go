package romantoint

var valueLookup = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	L := len(s)
	out := 0
	prevV := 0

	for i := L - 1; i >= 0; i-- {
		v := valueLookup[rune(s[i])]

		if v < prevV {
			out -= v
		} else {
			out += v
		}

		prevV = v
	}

	return out
}
