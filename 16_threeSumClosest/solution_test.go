package threesumclosest

import (
	"fmt"
	"testing"
)

func TestThreeSumCloset(t *testing.T) {
	for _, tc := range threeSumClosestTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := threeSumClosest(tc.nums, tc.target)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
