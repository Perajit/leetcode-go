package findmediansortedarrays

type FindMedianSortedArraysTestCase struct {
	nums1          []int
	nums2          []int
	expectedOutput float64
}

var findMedianSortedArraysTestCases []FindMedianSortedArraysTestCase = []FindMedianSortedArraysTestCase{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{0, 0}, []int{0, 0}, 0},
	{[]int{}, []int{1}, 1},
	{[]int{2}, []int{}, 2},
	{[]int{}, []int{}, 0},
	{[]int{1, 1, 6}, []int{2, 2, 4, 6, 8}, 3},
}
