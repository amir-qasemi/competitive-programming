// https://leetcode.com/problems/rotate-array/

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 || k == 0 {
		return
	}

	rotate1(nums, k)
}

// Solution with O(n) time complexity and O(n) space complexity
func rotate1(nums []int, k int) {
	res := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		res[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, res)
}

// Solution with O(nk) time complexity and O(1) space complexity
func rotate2(nums []int, k int) {
	rotateOnce := func(nums []int) {
		temp := nums[len(nums)-1]
		var temp2 int
		for i := 0; i < len(nums); i++ {
			temp2 = nums[i]
			nums[i] = temp
			temp = temp2
		}
	}
	for i := 0; i < k; i++ {
		rotateOnce(nums)
	}
}
