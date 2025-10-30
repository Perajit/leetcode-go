package addtwonumbers

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	for _, tc := range addTwoNumbersTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := addTwoNumbers(tc.l1, tc.l2)
			actualStr := toStr(actualOutput)
			expectedStr := toStr(tc.expectedOutput)

			if actualStr != expectedStr {
				t.Errorf("FAIL: Got %#v, want %#v", actualStr, expectedStr)
			}
		})
	}
}

func toStr(l *ListNode) string {
	n := l
	s := ""

	for n != nil {
		s = strconv.Itoa(n.Val) + s
		n = n.Next
	}

	return s
}
