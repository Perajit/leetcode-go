package removeduplicates

type RemoveDuplicatesTestCase struct {
	name           string
	nums           []int
	expectedResult []int
}

var removeDuplicatesTestCases = []RemoveDuplicatesTestCase{
	{"#1", []int{1, 1, 2}, []int{1, 2}},
	{"#2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
	{"#3", []int{1, 2, 3}, []int{1, 2, 3}},
	{"#4", []int{1, 1}, []int{1}},
	{"#5", []int{1}, []int{1}},
	{"#6", []int{}, []int{}},
	{"#7", []int{-1, 0, 0, 0, 0, 3, 3}, []int{-1, 0, 3}},
}
