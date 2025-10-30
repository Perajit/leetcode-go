package convert

import (
	"strings"
)

func convert(s string, numRows int) string {
	N := len(s)

	if numRows == 0 || N == 0 {
		return ""
	}

	if numRows == 1 || N <= numRows {
		return s
	}

	groupSize := 2*numRows - 2
	totalGroups := N / groupSize

	if N%groupSize > 0 {
		totalGroups++
	}

	rows := make([]string, numRows)

	for i, c := range s {
		// calculate row index
		r := i % groupSize

		if r+1 > numRows {
			r = groupSize - r
		}

		// update rows
		rows[r] += string(c)
	}

	return strings.Join(rows, "")
}

/*
P A H N
APLSIIG
Y I R

P  I  N
A LS IG
YA HR
P  I
*/
