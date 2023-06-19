// https://leetcode.com/problems/jump-game
func canJump(nums []int) bool {
	reachable := make([]bool, len(nums))
	reachable[0] = true
	for i := 0; i < len(nums); i++ {
		if reachable[i] {
			for j := 1; j <= nums[i]; j++ {
				if i+j < len(nums) {
					reachable[i+j] = true
				}
			}
		}
	}

	return reachable[len(nums)-1]
}