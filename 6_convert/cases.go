package convert

type ConvertTestCase struct {
	name           string
	s              string
	numRows        int
	expectedOutput string
}

var convertTestCases = []ConvertTestCase{
	{"#1", "", 2, ""},
	{"#2", "A", 1, "A"},
	{"#3", "AB", 3, "AB"},
	{"#4", "AB", 1, "AB"},
	{"#5", "ABC", 2, "ACB"},
	{"#6", "PAYPALISHIRING", 0, ""},
	{"#7", "PAYPALISHIRING", 1, "PAYPALISHIRING"},
	{"#8", "PAYPALISHIRING", 2, "PYAIHRNAPLSIIG"},
	{"#9", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"#10", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"#11", "PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
}
