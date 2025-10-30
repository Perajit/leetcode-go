package findmediansortedarrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	for i, tc := range findMedianSortedArraysTestCases {
		tcName := fmt.Sprintf("Case_%d_nums1_%#v_nums2_%#v", i, tc.nums1, tc.nums2)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := findMedianSortedArrays(tc.nums1, tc.nums2)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
