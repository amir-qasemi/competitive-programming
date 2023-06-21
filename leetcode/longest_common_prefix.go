// https://leetcode.com/problems/longest-common-prefix

func longestCommonPrefix(strs []string) string {
	var commonPrefix strings.Builder

OUTER:
	for i := 0; i < len(strs[0]); i++ {
		charToConsider := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j])-1 < i {
				break OUTER
			}

			if strs[j][i] != charToConsider {
				break OUTER
			}
		}
		commonPrefix.WriteByte(charToConsider)
	}

	return commonPrefix.String()
}