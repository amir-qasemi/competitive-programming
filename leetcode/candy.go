// https://leetcode.com/problems/candy

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}

	res := make([]int, len(ratings))
	for i, _ := range ratings {
		res[i] = 1
	}

	if ratings[0] > ratings[1] {
		res[0] += 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}

	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && res[i] <= res[i+1] {
			res[i] = res[i+1] + 1
		}

	}

	sum := 0
	for _, val := range res {
		sum += val
	}

	return sum
}