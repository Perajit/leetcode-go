package convert

import (
	"strings"
)

func convert(s string, numRows int) string {
	L := len(s)

	if numRows == 0 || L == 0 {
		return ""
	}

	if numRows == 1 || L <= numRows {
		return s
	}

	groupSize := 2*numRows - 2
	totalGroups := L / groupSize

	if L%groupSize > 0 {
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
