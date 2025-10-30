package reverse

type ReverseTestCase struct {
	x              int
	expectedOutput int
}

var reverseTestCases []ReverseTestCase = []ReverseTestCase{
	{0, 0},
	{1, 1},
	{-1, -1},
	{-10, -1},
	{123, 321},
	{120, 21},
	{1534236469, 0},
	{-8463847412, -2147483648},
	{8463847412, 0},
	{-9463847412, 0},
}
