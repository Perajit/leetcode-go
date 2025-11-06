package lettercombinations

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	for _, tc := range letterCombinationsTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := letterCombinations(tc.digits)

			if !checkEquality(actualOutput, tc.expectedOutput) {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}

func checkEquality(v1 []string, v2 []string) bool {
	sort.Strings(v1)
	sort.Strings(v2)

	return slices.Equal(v1, v2)
}
