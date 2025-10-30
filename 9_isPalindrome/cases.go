package ispalindrome

type IsPalindromeTestCase struct {
	x              int
	expectedOutput bool
}

var isPalindromeTestCases []IsPalindromeTestCase = []IsPalindromeTestCase{
	{0, true},
	{1, true},
	{11, true},
	{121, true},
	{-121, false},
	{10, false},
	{101, true},
	{1001, true},
	{10011, false},
}
