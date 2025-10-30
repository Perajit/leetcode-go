package convert

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	for i, tc := range convertTestCases {
		tcName := fmt.Sprintf("Case_%d_s_%s_numRows_%d", i, tc.s, tc.numRows)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := convert(tc.s, tc.numRows)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
