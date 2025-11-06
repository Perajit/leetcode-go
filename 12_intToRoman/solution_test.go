package inttoroman

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	for _, tc := range intToRomanTestCases {
		tcName := fmt.Sprintf("Case_%s", tc.name)

		t.Run(tcName, func(t *testing.T) {
			actualOutput := intToRoman(tc.num)

			if actualOutput != tc.expectedOutput {
				t.Errorf("FAIL: Got %#v, want %#v", actualOutput, tc.expectedOutput)
			}
		})
	}
}
