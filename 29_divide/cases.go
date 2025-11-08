package divide

type DivideTestCase struct {
	name           string
	dividend       int
	divisor        int
	expectedOutput int
}

var divideTestCases = []DivideTestCase{
	{"#1", 10, 3, 3},
	{"#2", 7, -3, -2},
	{"#3", -7, 3, -2},
	{"#4", -7, -3, 2},
	{"#5", 2, 3, 0},
	{"#6", 2, -3, 0},
	{"#7", 4294967296, 2, 2147483647},
	{"#8", 4294967296, -2, -2147483648},
	{"#9", -2147483648, 2, -1073741824},
}
