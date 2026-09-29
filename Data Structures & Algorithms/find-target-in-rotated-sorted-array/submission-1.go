import "slices"

func search(nums []int, target int) int {
	if len(nums) == 1 && target == nums[0] {
		return 0
	}

	minIdx := slices.Index(nums, slices.Min(nums))
	left, right := 0, len(nums)-1
	if nums[right] < target && minIdx != 0 {
		right = minIdx - 1
	} else {
		left = minIdx
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
