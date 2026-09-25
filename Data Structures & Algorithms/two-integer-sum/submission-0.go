func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	fmt.Print(left, right)
	for left != right {
		sum := nums[left] + nums[right]
		if sum == target {
			break
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{left, right}
}
