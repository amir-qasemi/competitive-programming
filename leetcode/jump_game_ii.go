// https://leetcode.com/problems/jump-game-ii

// O(n^^2) time, O(n) memory
func jump(nums []int) int {
	numOfJumps := make([]int, len(nums))
	numOfJumps[0] = 0
	for i := 1; i < len(nums); i++ {
		numOfJumps[i] = math.MaxInt
	}

	for i := 0; i < len(nums); i++ {
		for j := 1; j <= nums[i]; j++ {
			if i+j < len(nums) && numOfJumps[i]+1 < numOfJumps[i+j] {
				numOfJumps[i+j] = numOfJumps[i] + 1
			}
		}
	}

	return numOfJumps[len(numOfJumps)-1]
}