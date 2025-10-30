package reverse

type ReverseTestCase struct {
	name           string
	x              int
	expectedOutput int
}

var reverseTestCases = []ReverseTestCase{
	{"#1", 0, 0},
	{"#2", 1, 1},
	{"#3", -1, -1},
	{"#4", -10, -1},
	{"#5", 123, 321},
	{"#6", 120, 21},
	{"#7", 1534236469, 0},
	{"#8", -8463847412, -2147483648},
	{"#9", 8463847412, 0},
	{"#10", -9463847412, 0},
}
