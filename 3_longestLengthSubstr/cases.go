package lengthoflongestsubstring

type LengthOfLongestSubstringTestCases struct {
	name           string
	s              string
	expectedOutput int
}

var lengthOfLongestSubstringTestCases = []LengthOfLongestSubstringTestCases{
	{"#1", "abcabcbb", 3},          // abc
	{"#2", "bbbbb", 1},             // b
	{"#3", "pwwkew", 3},            // wke
	{"#4", "qwwerttttteeatopp", 5}, // eatop
	{"#5", "dvdf", 3},              // vdf
	{"#6", "ddddvdf", 3},           // vdf
	{"#7", "a", 1},                 // a
	{"#8", "", 0},
}
