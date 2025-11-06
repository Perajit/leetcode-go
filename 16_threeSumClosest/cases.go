package threesumclosest

type ThreeSumClosestTestCase struct {
	name           string
	nums           []int
	target         int
	expectedOutput int
}

var threeSumClosestTestCases = []ThreeSumClosestTestCase{
	{"#1", []int{-1, 2, 1, 4}, 1, 2},
	{"#2", []int{-1, 2, 1, 4}, 2, 2},
	{"#3", []int{-1, 2, 1, 4}, 4, 4},
	{"#4", []int{1, 0, 1, 2}, 2, 2},
	{"#5", []int{1, -2, -6, 2}, 0, 1},
	{"#6", []int{1, 2, 4, 8, 16, 32, 64, 128}, 82, 82},
	{"#7", []int{-1, 0, 1, 2, -1, -4}, 0, 0},
	{"#8", []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}, 0, 0},
	{"#9", []int{1, 2, 2, 4, 6, 6}, 6, 7},
	{"#10", []int{1, 1, 2, 2, 8, 8}, 8, 10},
	{"#11", []int{4, 0, 5, -5, 3, 3, 0, -4, -5}, -2, -2},
}
