// https://leetcode.com/problems/product-of-array-except-self

func productExceptSelf(nums []int) []int {
	result, prefixMult, postfixMult := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	prefixMult[0], postfixMult[len(nums)-1] = nums[0], nums[len(nums)-1]

	for i := 1; i < len(nums)-1; i++ {
		prefixMult[i] = nums[i] * prefixMult[i-1]
		postfixInd := len(nums) - 1 - i
		postfixMult[postfixInd] = nums[postfixInd] * postfixMult[postfixInd+1]
	}

	result[0] = postfixMult[1]
	result[len(nums)-1] = prefixMult[len(nums)-2]
	for i := 1; i < len(nums)-1; i++ {
		result[i] = prefixMult[i-1] * postfixMult[i+1]
	}
	return result
}