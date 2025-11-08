package generateparentheses

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	for _, tc := range generateParenthesisTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := generateParenthesis(tc.n)

			if !checkEquality(actualOutput, tc.expectedOutput) {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}

func checkEquality(v1 []string, v2 []string) bool {
	if len(v1) != len(v2) {
		return false
	}

	sort.Strings(v1)
	sort.Strings(v2)

	return slices.Equal(v1, v2)
}
