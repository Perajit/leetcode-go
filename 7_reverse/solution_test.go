package reverse

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	for i, tc := range reverseTestCases {
		tcName := fmt.Sprintf("Case_%d_x_%d", i, tc.x)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := reverse(tc.x)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
