package lettercombinations

type LetterCombinationsTestCase struct {
	name           string
	digits         string
	expectedOutput []string
}

var letterCombinationsTestCases = []LetterCombinationsTestCase{
	{"#1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"#2", "", []string{}},
	{"#3", "2", []string{"a", "b", "c"}},
	{"#4", "3", []string{"d", "e", "f"}},
	{"#5", "4", []string{"g", "h", "i"}},
	{"#6", "5", []string{"j", "k", "l"}},
	{"#7", "6", []string{"m", "n", "o"}},
	{"#8", "7", []string{"p", "q", "r", "s"}},
	{"#9", "8", []string{"t", "u", "v"}},
	{"#10", "9", []string{"w", "x", "y", "z"}},
	{"#11", "22", []string{
		"aa", "ab", "ac",
		"ba", "bb", "bc",
		"ca", "cb", "cc",
	}},
	{"#12", "223", []string{
		"aad", "aae", "aaf",
		"abd", "abe", "abf",
		"acd", "ace", "acf",
		"bad", "bae", "baf",
		"bbd", "bbe", "bbf",
		"bcd", "bce", "bcf",
		"cad", "cae", "caf",
		"cbd", "cbe", "cbf",
		"ccd", "cce", "ccf",
	}},
}
