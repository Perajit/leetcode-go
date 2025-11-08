package removeelement

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, tc := range removeElementTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := removeElement(tc.nums, tc.val)
			actualResult := tc.nums[:actualOutput]
			sort.Ints(actualResult)
			sort.Ints(tc.expectedResult)

			if !slices.Equal(actualResult, tc.expectedResult) {
				t.Errorf("FAIL: Got %#v, want %#v", actualResult, tc.expectedResult)
			}
		})
	}
}
