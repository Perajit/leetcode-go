package nextpermutation

type NextPermutationTestCase struct {
	name           string
	nums           []int
	expectedResult []int
}

var nextPermutationTestCases = []NextPermutationTestCase{
	{"#1", []int{1, 2, 3}, []int{1, 3, 2}},
	{"#2", []int{3, 2, 1}, []int{1, 2, 3}},
	{"#3", []int{1, 1, 5}, []int{1, 5, 1}},
	{"#4", []int{1, 3, 2}, []int{2, 1, 3}},
	{"#5", []int{1, 3, 9, 8, 6}, []int{1, 6, 3, 8, 9}},
	{"#6", []int{1, 6, 9, 8, 3}, []int{1, 8, 3, 6, 9}},
	{"#7", []int{1, 6, 9, 5, 3}, []int{1, 9, 3, 5, 6}},
	{"#8", []int{1, 6, 9, 8, 3, 2}, []int{1, 8, 2, 3, 6, 9}},
	{"#8", []int{1, 8, 9, 6, 3, 2}, []int{1, 9, 2, 3, 6, 8}},
}

type ReverseTestCase struct {
	name           string
	nums           []int
	start          int
	end            int
	expectedResult []int
}

var reverseTestCases = []ReverseTestCase{
	{"#1", []int{1, 8, 9, 6, 3}, 2, 5, []int{1, 8, 3, 6, 9}},
	{"#2", []int{1, 8, 9, 6, 3}, 1, 4, []int{1, 6, 9, 8, 3}},
	{"#3", []int{1, 6, 9, 8, 3, 2}, 2, 6, []int{1, 6, 2, 3, 8, 9}},
	{"#4", []int{1, 6, 9, 8, 3, 2}, 1, 5, []int{1, 3, 8, 9, 6, 2}},
	{"#5", []int{1, 6, 9, 8, 3, 2}, 1, 6, []int{1, 2, 3, 8, 9, 6}},
}
