// https://leetcode.com/problems/valid-palindrome

func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	clean := strings.ReplaceAll(strings.ToLower(reg.ReplaceAllString(s, "")), " ", "")
	fmt.Println(clean)
	for i := 0; i < len(clean)/2; i++ {
		if clean[i] != clean[len(clean)-1-i] {
			fmt.Println(i)
			return false
		}
	}

	return true
}