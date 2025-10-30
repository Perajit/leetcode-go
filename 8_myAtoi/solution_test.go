package myatoi

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	for i, tc := range myAtoiTestCases {
		tcName := fmt.Sprintf("Case_%d_str_%s", i, tc.str)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := myAtoi(tc.str)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
