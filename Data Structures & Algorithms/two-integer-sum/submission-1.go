func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	fmt.Print(left, right)
	for left != right {
		sum := nums[left] + nums[right]
		print("left: ", left, ", right: ", right, "\n")
		if sum == target {
			break
		} else if sum < target {
			print("sum ", sum, " < ", target, " target\n")
			left++
		} else if sum > target {
			print("sum ", sum, " > ", target, " target\n")
			right--
		}
	}

	return []int{left, right}
}
