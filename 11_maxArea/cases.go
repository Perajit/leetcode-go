package maxarea

type MaxAreaTestCase struct {
	name           string
	height         []int
	expectedOutput int
}

var maxAreaTestCases = []MaxAreaTestCase{
	{"#1", []int{}, 0},
	{"#2", []int{10}, 0},
	{"#3", []int{10, 10}, 10},
	{"#4", []int{1, 10}, 1},
	{"#5", []int{1, 10, 2}, 2},
	{"#6", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
}
