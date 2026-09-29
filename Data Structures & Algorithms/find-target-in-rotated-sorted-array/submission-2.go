func search(nums []int, target int) int {
	if len(nums) == 1 && target == nums[0] {
		return 0
	}

	left, right := 0, len(nums)-1
	for left < right {
		middle := left + (right-left)/2
		if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			right = middle
		}
	}
	minimum := left

	left, right = 0, len(nums)-1
	if nums[right] < target && minimum != 0 {
		right = minimum -1
	} else {
		left = minimum
	}

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			return middle
		}
		if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return -1
}