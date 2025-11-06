package lettercombinations

var valueLookup = map[byte]string{
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	L := len(digits)

	if L < 1 {
		return []string{}
	}

	out := []string{""}

	for i := 0; i < L; i++ {
		s := valueLookup[digits[i]]
		out = distribute(s, out)
	}

	return out
}

func distribute(s string, initial []string) []string {
	out := make([]string, len(s)*len(initial))

	for is, cs := range s {
		offset := is * len(initial)

		for ii, ci := range initial {
			out[offset+ii] = string(ci) + string(cs)
		}
	}

	return out
}
