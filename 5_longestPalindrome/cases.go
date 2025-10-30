package longestpalindrome

type LongestPalindromeTestCases struct {
	name            string
	s               string
	expectedOutputs []string
}

var longestPalindromeTestCases = []LongestPalindromeTestCases{
	{"#1", "babad", []string{"aba", "bab"}},
	{"#2", "cbbd", []string{"bb"}},
	{"#3", "", []string{""}},
	{"#4", "a", []string{"a"}},
	{"#5", "bb", []string{"bb"}},
	{"#6", "abbcbb", []string{"bbcbb"}},
	{"#7", "abadd", []string{"aba"}},
	{"#8", "cocoachocolate", []string{"coc", "oco"}},
	{"#9", "mochacappuppacahcom", []string{"mochacappuppacahcom"}},
	{"#10", "crabbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbbbbbbbbbbbbba", []string{"abbbbbbbbbbbbbbbbbbbbbbbbbbbabbbbbbbbbbbbbbbbbbbbbbbbbbba"}},
	{"#11", "have a nice day yad ecin a evah", []string{"have a nice day yad ecin a evah"}},
}

type IsPalindromeTestCase struct {
	name           string
	s              string
	start          int
	end            int
	expectedOutput bool
}

var isPalindromeTestCases []IsPalindromeTestCase = []IsPalindromeTestCase{
	// ""
	{"#1", "", 0, 0, true},
	{"#2", "a", 0, 0, true},
	// "a"
	{"#3", "a", 0, 1, true},
	{"#4", "abc", 0, 1, true},
	{"#5", "cab", 1, 2, true},
	{"#6", "cba", 2, 3, true},
	// bb
	{"#7", "bb", 0, 2, true},
	// bbabb
	{"#8", "bbabb", 0, 5, true},
	// bbaabb
	{"#9", "bbaabb", 0, 6, true},
	// abbcbba
	{"#10", "abbcbba", 0, 7, true},
	{"#11", "aaabbcbbacc", 2, 9, true},
	// abbcdbba
	{"#12", "abbcdbba", 0, 8, false},
	// aabb
	{"#13", "aaabbcd", 1, 5, false},
}
