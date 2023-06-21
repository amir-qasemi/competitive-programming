// https://leetcode.com/problems/trapping-rain-water

func trap(height []int) int {
	return trap2(height)
}

// O(n^^2) time, O(n^^2) space
func trap1(height []int) int {
	max := height[0]
	for _, val := range height {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
	fmt.Println(len(height))
	matrix := make([][]bool, max)
	for i, _ := range matrix {
		matrix[i] = make([]bool, len(height))
	}

	for j := 0; j < len(height); j++ {
		for i := 0; i < max; i++ {
			if i > max-height[j]-1 {
				matrix[i][j] = true
			}
		}
	}

	sum := 0
	for i := 0; i < max; i++ {
		prevPopulatedCol := -1
		for j := 0; j < len(height); j++ {
			if matrix[i][j] {
				if prevPopulatedCol != -1 {
					sum += j - prevPopulatedCol - 1
				}
				prevPopulatedCol = j
			}
		}
	}

	return sum
}

// O(n) time, O(n) space
func trap2(height []int) int {
	result, leftMax, rightMax := make([]int, len(height)), make([]int, len(height)), make([]int, len(height))
	leftMaxSoFar, rightMaxSoFar := height[0], height[len(height)-1]

	for i := 0; i < len(height); i++ {
		if height[i] > leftMaxSoFar {
			leftMaxSoFar = height[i]
		}
		leftMax[i] = leftMaxSoFar

		if height[len(height)-1-i] > rightMaxSoFar {
			rightMaxSoFar = height[len(height)-1-i]
		}
		rightMax[len(height)-1-i] = rightMaxSoFar
	}

	result[0], result[len(height)-1] = 0, 0
	for i := 1; i < len(height)-1; i++ {
		if rightMax[i] > height[i] && leftMax[i] > height[i] {
			result[i] = int(math.Min(float64(rightMax[i]), float64(leftMax[i]))) - height[i]
		}
	}

	sum := 0
	for _, val := range result {
		sum += val
	}

	return sum
}