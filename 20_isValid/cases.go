package isvalid

type IsValidTestCase struct {
	name           string
	s              string
	expectedOutput bool
}

var isValidTestCases = []IsValidTestCase{
	{"#1", "(", false},
	{"#2", "{", false},
	{"#3", "[", false},
	{"#4", ")", false},
	{"#5", "}", false},
	{"#6", "]", false},
	{"#7", "()", true},
	{"#8", "{}", true},
	{"#9", "[]", true},
}
