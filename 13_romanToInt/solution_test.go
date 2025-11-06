package romantoint

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	for _, tc := range romanToIntTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := romanToInt(tc.s)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
