func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]

	currentSum := 0
	for _, num := range nums {
		currentSum += num
		maxSum = max(maxSum, currentSum)

		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}
