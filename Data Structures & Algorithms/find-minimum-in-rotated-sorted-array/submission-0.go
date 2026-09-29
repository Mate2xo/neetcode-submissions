func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	first, last := nums[0], nums[len(nums)-1]
	if first < last {
		return first
	}

	middleIdx := len(nums)/2
	middle := nums[middleIdx]
	if middle < last {
		return findMin(nums[:middleIdx+1])
	} else {
		return findMin(nums[middleIdx:])
	}

}
