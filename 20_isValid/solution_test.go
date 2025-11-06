package isvalid

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	for _, tc := range isValidTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := isValid(tc.s)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
