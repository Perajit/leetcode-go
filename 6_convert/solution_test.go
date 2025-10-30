package convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	for _, tc := range convertTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := convert(tc.s, tc.numRows)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
