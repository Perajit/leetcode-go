package twosum

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	for _, tc := range twoSumTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := twoSum(tc.nums, tc.target)
			sort.Ints(actualOutput)

			if !slices.Equal(actualOutput, tc.expectedOutput) {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
