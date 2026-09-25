func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	fmt.Print(i, j)
	for i != j {
		left, right := nums[i], nums[j]
		sum := left + right
		if sum == target {
			break
		} else if sum < target {
			if left < right {
				i++	
			} else  {
				j--
			}
		} else if sum > target {
			if left < right {
				j--
			} else  {
				i++	
			}
		}
	}

	return []int{i, j}
}
