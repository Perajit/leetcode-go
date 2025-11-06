package inttoroman

type IntToRomanTestCase struct {
	name           string
	num            int
	expectedOutput string
}

var intToRomanTestCases = []IntToRomanTestCase{
	{"#1", 3749, "MMMDCCXLIX"},
	{"#2", 58, "LVIII"},
	{"#3", 1994, "MCMXCIV"},
	{"#4", 0, ""},
	{"#5", 1, "I"},
	{"#6", 2, "II"},
	{"#7", 3, "III"},
	{"#8", 4, "IV"},
	{"#9", 5, "V"},
	{"#10", 6, "VI"},
	{"#11", 7, "VII"},
	{"#12", 8, "VIII"},
	{"#13", 9, "IX"},
	{"#14", 10, "X"},
	{"#15", 20, "XX"},
	{"#16", 30, "XXX"},
	{"#17", 40, "XL"},
	{"#18", 50, "L"},
	{"#19", 60, "LX"},
	{"#20", 70, "LXX"},
	{"#21", 80, "LXXX"},
	{"#22", 90, "XC"},
	{"#23", 100, "C"},
	{"#24", 200, "CC"},
	{"#25", 300, "CCC"},
	{"#26", 400, "CD"},
	{"#27", 500, "D"},
	{"#28", 600, "DC"},
	{"#29", 700, "DCC"},
	{"#30", 800, "DCCC"},
	{"#31", 900, "CM"},
	{"#32", 1000, "M"},
	{"#33", 13, "XIII"},
	{"#34", 38, "XXXVIII"},
	{"#35", 88, "LXXXVIII"},
	{"#36", 808, "DCCCVIII"},
	{"#37", 888, "DCCCLXXXVIII"},
	{"#38", 14, "XIV"},
	{"#39", 19, "XIX"},
	{"#40", 409, "CDIX"},
	{"#41", 999, "CMXCIX"},
	{"#42", 3999, "MMMCMXCIX"},
}
