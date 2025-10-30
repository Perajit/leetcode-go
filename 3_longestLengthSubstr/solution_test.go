package lengthoflongestsubstring

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range lengthOfLongestSubstringTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := lengthOfLongestSubstring(tc.s)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
