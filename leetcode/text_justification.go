// https://leetcode.com/problems/text-justification
amanaplanacanalpanama
func fullJustify(words []string, maxWidth int) []string {
	lens := make([]int, len(words))
	var result []string
	for i, val := range words {
		lens[i] = len(val)
	}

	buildLine := func(i, numOfCurrentLetters, numOfCurrentWords int, lastLine bool) string {
		var line strings.Builder
		for j := i - numOfCurrentWords; j < i; j++ {
			line.WriteString(words[j])
			if j < i-1 && !lastLine {
				line.WriteString(strings.Repeat(" ", (maxWidth-numOfCurrentLetters)/(numOfCurrentWords-1)))
				if j-(i-numOfCurrentWords) < (maxWidth-numOfCurrentLetters)%(numOfCurrentWords-1) {
					line.WriteString(" ")
				}
			} else if numOfSpaces := maxWidth - (numOfCurrentLetters + numOfCurrentWords - 1); numOfCurrentWords == 1 && !lastLine && numOfSpaces >= 0 {
				line.WriteString(strings.Repeat(" ", numOfSpaces))
			} else if lastLine && j < i-1 {
				line.WriteString(" ")
			}
		}

		if numOfSpaces := maxWidth - (numOfCurrentLetters + numOfCurrentWords - 1); lastLine && numOfSpaces >= 0 {
			line.WriteString(strings.Repeat(" ", numOfSpaces))
		}
		return line.String()
	}

	numOfCurrentWords := 0
	numOfCurrentLetters := 0
	for i := 0; i < len(words); i++ {
		if len(words[i])+numOfCurrentLetters+numOfCurrentWords <= maxWidth {
			numOfCurrentLetters += len(words[i])
			numOfCurrentWords++
		} else {
			result = append(result, buildLine(i, numOfCurrentLetters, numOfCurrentWords, false))

			numOfCurrentLetters = len(words[i])
			numOfCurrentWords = 1
		}
	}

	if numOfCurrentWords > 0 {
		fmt.Println("Last Line")
		result = append(result, buildLine(len(words), numOfCurrentLetters, numOfCurrentWords, true))
	}

	return result
}