package ispalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for i, tc := range isPalindromeTestCases {
		tcName := fmt.Sprintf("Case_%d_x_%d", i, tc.x)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := isPalindrome(tc.x)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
