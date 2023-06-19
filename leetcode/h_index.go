// https://leetcode.com/problems/h-index

// O(n) time, O(1) space
func hIndex(citations []int) int {
	if citations == nil || len(citations) == 0 {
		return 0
	}

	hIndex := 0

	sort.Sort(sort.IntSlice(citations))
	for i := len(citations) - 1; i >= 0; i-- {
		proposedHIndex := 0
		if (len(citations) - i) >= citations[i] {
			proposedHIndex = citations[i]
		} else {
			proposedHIndex = len(citations) - i
		}

		if proposedHIndex > hIndex {
			hIndex = proposedHIndex
		}
	}

	return hIndex
}