package longestcommonprefix

type LongestCommonPrefixTestCase struct {
	name           string
	strs           []string
	expectedOutput string
}

var longestCommonPrefixTestCases = []LongestCommonPrefixTestCase{
	{"#1", []string{""}, ""},
	{"#2", []string{"test"}, "test"},
	{"#3", []string{"happy", "happy"}, "happy"},
	{"#4", []string{"flower", "flow", "flight"}, "fl"},
	{"#5", []string{"dog", "racecar", "car"}, ""},
}
