package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	L := len(s)

	if L < 2 {
		return L
	}

	runes := []rune(s)
	counterLookup := map[rune]int{runes[0]: 1}
	i := 0
	j := 1
	longest := 1

	for j < L {
		counter, ok := counterLookup[runes[j]]

		// duplicated
		if ok && counter > 0 {
			// reduce counter at i and move i to the right
			counterLookup[runes[i]] -= 1
			i++
		} else {
			// update longest
			l := j - i + 1
			if longest < l {
				longest = j - i + 1
			}

			// increase counter at j and move j to the right
			counterLookup[runes[j]] = 1
			j++
		}
	}

	return longest
}
