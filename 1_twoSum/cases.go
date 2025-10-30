package twosum

type TwoSumTestCase struct {
	name           string
	nums           []int
	target         int
	expectedOutput []int
}

var twoSumTestCases = []TwoSumTestCase{
	{"#1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"#2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"#3", []int{3, 3}, 6, []int{0, 1}},
	{"#4", []int{11, 7, 2, 1, 15}, 17, []int{2, 4}},
}
