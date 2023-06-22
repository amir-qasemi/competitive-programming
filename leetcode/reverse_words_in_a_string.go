// https://leetcode.com/problems/reverse-words-in-a-string

func reverseWords(s string) string {
	var sb strings.Builder
	spaceWritten := true
	var detectedWordSb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && !spaceWritten {
			sb.WriteString(reverseWord(detectedWordSb))
			detectedWordSb.Reset()
			sb.WriteByte(' ')
			spaceWritten = true
		} else if s[i] != ' ' {
			detectedWordSb.WriteByte(s[i])
			spaceWritten = false
		}
	}
	sb.WriteString(reverseWord(detectedWordSb))
	res := sb.String()
	if res[len(res)-1] == ' ' {
		return res[:len(res)-1]
	} else {
		return res
	}
}

func reverseWord(sb strings.Builder) string {
	s := sb.String()

	var res strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		res.WriteByte(s[i])
	}

	return res.String()
}