package addtwonumbers

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	for i, tc := range addTwoNumbersTestCases {
		tcName := fmt.Sprintf("Case_%d_l1_%s_l2_%s", i, toStr(tc.l1), toStr(tc.l2))

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
		s = strconv.Itoa(n.val) + s
		n = n.next
	}

	return s
}
