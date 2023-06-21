// https://leetcode.com/problems/length-of-last-word

func lengthOfLastWord(s string) int {
	runes := []rune(s)
	startId, endId := 0, -1
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] != ' ' && endId == -1 {
			endId = i
		} else if runes[i] == ' ' && endId != -1 {
			startId = i + 1
			break
		}
	}

	return endId - startId + 1
}