package romantoint

type RomanToIntTestCase struct {
	name           string
	s              string
	expectedOutput int
}

var romanToIntTestCases = []RomanToIntTestCase{
	{"#1", "", 0},
	{"#2", "I", 1},
	{"#3", "II", 2},
	{"#4", "III", 3},
	{"#5", "IV", 4},
	{"#6", "V", 5},
	{"#7", "VI", 6},
	{"#8", "X", 10},
	{"#9", "XX", 20},
	{"#10", "XXX", 30},
	{"#11", "XL", 40},
	{"#12", "L", 50},
	{"#13", "LX", 60},
	{"#14", "LXX", 70},
	{"#15", "LXXX", 80},
	{"#16", "XC", 90},
	{"#17", "C", 100},
	{"#18", "CC", 200},
	{"#19", "CCC", 300},
	{"#20", "CD", 400},
	{"#21", "D", 500},
	{"#22", "DC", 600},
	{"#23", "DCC", 700},
	{"#24", "DCCC", 800},
	{"#25", "CM", 900},
	{"#26", "M", 1000},
	{"#27", "MM", 2000},
	{"#28", "MMM", 3000},
	{"#29", "LVIII", 58},
	{"#30", "MCMXCIV", 1994},
	{"#31", "XIII", 13},
	{"#32", "XCIII", 93},
	{"#33", "DCVI", 606},
	{"#34", "DCLXXXVI", 686},
	{"#35", "CDXII", 412},
	{"#36", "LXXXVIII", 88},
	{"#37", "DCCCVIII", 808},
	{"#38", "DCCCLXXXVIII", 888},
	{"#39", "MCMXCIX", 1999},
	{"#40", "MMMCMXCIX", 3999},
}
