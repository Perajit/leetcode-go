package divide

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	for _, tc := range divideTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := divide(tc.dividend, tc.divisor)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
