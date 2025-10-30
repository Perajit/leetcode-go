package twosum

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	for i, tc := range twoSumTestCases {
		tcName := fmt.Sprintf("Case_%d_Nums_%#v_Target_%d", i, tc.nums, tc.target)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := twoSum(tc.nums, tc.target)
			sort.Ints(actualOutput)

			if !slices.Equal(actualOutput, tc.expectedOutput) {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
