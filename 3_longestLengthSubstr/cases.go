package lengthoflongestsubstring

type LengthOfLongestSubstringTestCases struct {
	s              string
	expectedOutput int
}

var lengthOfLongestSubstringTestCases []LengthOfLongestSubstringTestCases = []LengthOfLongestSubstringTestCases{
	{"abcabcbb", 3},          // abc
	{"bbbbb", 1},             // b
	{"pwwkew", 3},            // wke
	{"qwwerttttteeatopp", 5}, // eatop
	{"dvdf", 3},              // vdf
	{"ddddvdf", 3},           // vdf
	{"a", 1},                 // a
	{"", 0},
}
