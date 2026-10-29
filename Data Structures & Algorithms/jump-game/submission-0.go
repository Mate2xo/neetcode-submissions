func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	start, end, foundPath := len(nums)-2, len(nums)-1, false
	for start >= 0 {
		if start+nums[start] >= end {
			end = start
			start--
			foundPath = true
		} else {
			start--
			foundPath = false
		}
	}

	return foundPath
}
