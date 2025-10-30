package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	for _, tc := range reverseTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := reverse(tc.x)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
