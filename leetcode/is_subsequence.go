// https://leetcode.com/problems/is-subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 && len(s) == 0 {
		return true
	}

	if len(t) != 0 && len(s) == 0 {
		return false
	}
	i := 0
	sRune := []rune(s)
	for _, val := range t {
		if val == sRune[i] {
			i++
		}

		if i >= len(sRune) {
			return true
		}
	}

	return false
}