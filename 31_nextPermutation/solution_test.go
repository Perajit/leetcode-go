package nextpermutation

import (
	"fmt"
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	for _, tc := range nextPermutationTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			nextPermutation(tc.nums)

			if !slices.Equal(tc.nums, tc.expectedResult) {
				t.Errorf("FAIL: Got %#v, want %#v", tc.nums, tc.expectedResult)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range reverseTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			reverse(&tc.nums, tc.start, tc.end)

			if !slices.Equal(tc.nums, tc.expectedResult) {
				t.Errorf("FAIL: Got %#v, want %#v", tc.nums, tc.expectedResult)
			}
		})
	}
}
