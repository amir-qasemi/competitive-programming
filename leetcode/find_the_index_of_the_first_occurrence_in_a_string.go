// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string

func strStr(haystack string, needle string) int {
	indToMatch := 0
	for i := 0; i < len(haystack); i++ {
		if needle[indToMatch] == haystack[i] {
			indToMatch++
			if indToMatch == len(needle) {
				return i - len(needle) + 1
			}
		} else {
			i -= indToMatch
			indToMatch = 0
		}
	}

	return -1
}