package foursum

type FourSumTestCase struct {
	name           string
	nums           []int
	target         int
	expectedOutput [][]int
}

var fourSumTestCases = []FourSumTestCase{
	{"#1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}},
	{"#2", []int{2, 2, 2, 2, 2}, 8, [][]int{
		{2, 2, 2, 2},
	}},
	{"#3", []int{1, -2}, -1, [][]int{}},
	{"#4", []int{1, 0, 1, 2}, 0, [][]int{}},
	{"#5", []int{1, 0, 1, 2}, 4, [][]int{
		{1, 0, 1, 2},
	}},
	{"#6", []int{-3, -1, 0, 2, 4, 5}, 2, [][]int{
		{-3, -1, 2, 4},
	}},
}
