package findmediansortedarrays

type FindMedianSortedArraysTestCase struct {
	name           string
	nums1          []int
	nums2          []int
	expectedOutput float64
}

var findMedianSortedArraysTestCases = []FindMedianSortedArraysTestCase{
	{"#1", []int{1, 3}, []int{2}, 2},
	{"#2", []int{1, 2}, []int{3, 4}, 2.5},
	{"#3", []int{0, 0}, []int{0, 0}, 0},
	{"#4", []int{}, []int{1}, 1},
	{"#5", []int{2}, []int{}, 2},
	{"#6", []int{}, []int{}, 0},
	{"#7", []int{1, 1, 6}, []int{2, 2, 4, 6, 8}, 3},
}
