package longestpalindrome

import (
	"fmt"
	"slices"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	for i, tc := range longestPalindromeTestCases {
		tcName := fmt.Sprintf("Case_%d_s_%s", i, tc.s)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := longestPalindrome(tc.s)

			if !slices.Contains(tc.expectedOutputs, actualOutput) {
				t.Errorf("FAIL: Got %#v, want one of %#v", actualOutput, tc.expectedOutputs)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	for i, tc := range isPalindromeTestCases {
		tcName := fmt.Sprintf("Case_%d_s_%s", i, tc.s)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := isPalindrome(tc.s, tc.start, tc.end)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
