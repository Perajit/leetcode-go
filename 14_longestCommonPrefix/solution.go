package longestcommonprefix

import "math"

func longestCommonPrefix(strs []string) string {
	N := len(strs)

	if N == 1 {
		return strs[0]
	}

	// find shortest length
	shortest := math.MaxInt

	for i := 0; i < N; i++ {
		l := len(strs[i])

		if l < shortest {
			shortest = l
		}
	}

	if shortest == 0 {
		return ""
	}

	// check prefix until shortest
	out := ""

	for i := 0; i < shortest; i++ {
		included := true
		cRef := strs[0][i]

		for n := 1; n < N; n++ {
			c := strs[n][i]

			if c != cRef {
				included = false
				break
			}
		}

		if !included {
			break
		}

		out += string(cRef)
	}

	return out
}
