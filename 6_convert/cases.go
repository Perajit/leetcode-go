package convert

type ConvertTestCase struct {
	s              string
	numRows        int
	expectedOutput string
}

var convertTestCases []ConvertTestCase = []ConvertTestCase{
	{"", 2, ""},
	{"A", 1, "A"},
	{"AB", 3, "AB"},
	{"AB", 1, "AB"},
	{"ABC", 2, "ACB"},
	{"PAYPALISHIRING", 0, ""},
	{"PAYPALISHIRING", 1, "PAYPALISHIRING"},
	{"PAYPALISHIRING", 2, "PYAIHRNAPLSIIG"},
	{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
	{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
	{"PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
}
