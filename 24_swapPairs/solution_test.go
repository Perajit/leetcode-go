package swappairs

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	for _, tc := range swapPairsTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := swapPairs(tc.head)
			actualSlice := toSlice(actualOutput)
			expectedSlice := toSlice(tc.expectedOutput)

			if !slices.Equal(actualSlice, expectedSlice) {
				t.Errorf("FAIL: Got %#v, want %#v", actualSlice, expectedSlice)
			}
		})
	}
}

func toSlice(l *ListNode) []string {
	n := l
	strs := []string{}

	for n != nil {
		strs = append(strs, strconv.Itoa(n.Val))
		n = n.Next
	}

	return strs
}
