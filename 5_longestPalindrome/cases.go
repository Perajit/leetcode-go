package longestpalindrome

type LongestPalindromeTestCases struct {
	s               string
	expectedOutputs []string
}

var longestPalindromeTestCases []LongestPalindromeTestCases = []LongestPalindromeTestCases{
	{"babad", []string{"aba", "bab"}},
	{"cbbd", []string{"bb"}},
	{"", []string{""}},
	{"a", []string{"a"}},
	{"bb", []string{"bb"}},
	{"abbcbb", []string{"bbcbb"}},
	{"abadd", []string{"aba"}},
	{"cocoachocolate", []string{"coc", "oco"}},
	{"mochacappuppacahcom", []string{"mochacappuppacahcom"}},
	{"crabbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbbbbbbbbbbbbba", []string{"abbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbbbbbbbbbbbbba"}},
	{"have a nice day yad ecin a evah", []string{"have a nice day yad ecin a evah"}},
}

type IsPalindromeTestCase struct {
	s              string
	start          int
	end            int
	expectedOutput bool
}

var isPalindromeTestCases []IsPalindromeTestCase = []IsPalindromeTestCase{
	// ""
	{"", 0, 0, true},
	{"a", 0, 0, true},
	// "a"
	{"a", 0, 1, true},
	{"abc", 0, 1, true},
	{"cab", 1, 2, true},
	{"cba", 2, 3, true},
	// bb
	{"bb", 0, 2, true},
	// bbabb
	{"bbabb", 0, 5, true},
	// bbaabb
	{"bbaabb", 0, 6, true},
	// abbcbba
	{"abbcbba", 0, 7, true},
	{"aaabbcbbacc", 2, 9, true},
	// abbcdbba
	{"abbcdbba", 0, 8, false},
	// aabb
	{"aaabbcd", 1, 5, false},
}
