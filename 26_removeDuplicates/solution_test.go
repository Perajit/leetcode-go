package removeduplicates

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	for _, tc := range removeDuplicatesTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := removeDuplicates(tc.nums)
			actualResult := tc.nums[:actualOutput]

			if !slices.Equal(actualResult, tc.expectedResult) {
				t.Errorf("FAIL: Got %#v, want %#v", actualResult, tc.expectedResult)
			}
		})
	}
}
