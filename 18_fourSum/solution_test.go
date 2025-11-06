package foursum

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestFourSum(t *testing.T) {
	for _, tc := range fourSumTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := fourSum(tc.nums, tc.target)

			if !checkEquality(actualOutput, tc.expectedOutput) {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}

func checkEquality(v1 [][]int, v2 [][]int) bool {
	if len(v1) != len(v2) {
		return false
	}

	for i := 0; i < len(v1); i++ {
		included := false

		for j := 0; j < len(v2); j++ {
			sort.Ints(v1[i])
			sort.Ints(v2[j])

			if slices.Equal(v1[i], v2[j]) {
				included = true
				break
			}
		}

		if !included {
			return false
		}
	}

	return true
}
