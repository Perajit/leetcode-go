package ispalindrome

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	for _, tc := range isPalindromeTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := isPalindrome(tc.x)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
