package strstr

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	for _, tc := range strStrTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := strStr(tc.haystack, tc.needle)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
