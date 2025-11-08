package strstr

func strStr(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)

	if ln == 0 || haystack == needle {
		return 0
	}

	if lh < ln {
		return -1
	}

	for i := 0; i < lh; i++ {
		if i+ln <= lh && haystack[i] == needle[0] {
			matched := true

			for j := 1; j < ln; j++ {
				if haystack[i+j] != needle[j] {
					matched = false
					break
				}
			}

			if matched {
				return i
			}
		}
	}

	return -1
}
