package longestcommonprefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	for _, tc := range longestCommonPrefixTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := longestCommonPrefix(tc.strs)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
