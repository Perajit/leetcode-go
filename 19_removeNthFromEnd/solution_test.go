package removenthfromend

import (
	"fmt"
	"slices"
	"strconv"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	for _, tc := range removeNthFromEndTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := removeNthFromEnd(tc.head, tc.n)
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
