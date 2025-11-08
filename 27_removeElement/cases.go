package removeelement

type RemoveElementTestCase struct {
	name           string
	nums           []int
	val            int
	expectedResult []int
}

var removeElementTestCases = []RemoveElementTestCase{
	{"#1", []int{3, 2, 2, 3}, 3, []int{2, 2}},
	{"#2", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 4, 0, 3}},
	{"#3", []int{3}, 3, []int{}},
	{"#4", []int{3}, 2, []int{3}},
	{"#5", []int{3}, 4, []int{3}},
	{"#6", []int{3, 3}, 2, []int{3, 3}},
	{"#7", []int{3, 3}, 4, []int{3, 3}},
	{"#8", []int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3}},
	{"#9", []int{5, 2, 3, 2, 3, 4, 1}, 3, []int{1, 2, 2, 5, 4}},
	{"#10", []int{3, 3, 3}, 3, []int{}},
	{"#11", []int{3, 3, 3}, 0, []int{3, 3, 3}},
}
