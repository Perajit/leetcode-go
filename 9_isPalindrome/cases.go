package ispalindrome

type IsPalindromeTestCase struct {
	name           string
	x              int
	expectedOutput bool
}

var isPalindromeTestCases = []IsPalindromeTestCase{
	{"#1", 0, true},
	{"#2", 1, true},
	{"#3", 11, true},
	{"#4", 121, true},
	{"#5", -121, false},
	{"#6", 10, false},
	{"#7", 101, true},
	{"#8", 1001, true},
	{"#9", 10011, false},
}
