package strstr

type StrStrTestCase struct {
	name           string
	haystack       string
	needle         string
	expectedOutput int
}

var strStrTestCases = []StrStrTestCase{
	{"#1", "hello", "ll", 2},
	{"#2", "aaaaa", "bba", -1},
	{"#3", "", "", 0},
	{"#4", "a", "", 0},
	{"#5", "a", "aa", -1},
	{"#6", "a", "a", 0},
	{"#7", "abc", "c", 2},
	{"#8", "mississippi", "pi", 9},
	{"#9", "mississippi", "issipi", -1},
}
