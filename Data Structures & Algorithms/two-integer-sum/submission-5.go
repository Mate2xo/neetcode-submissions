func twoSum(nums []int, target int) []int {
   for i, left := range nums {
		for j, right := range nums {
			if i == j {
				continue
			}

			if left+right == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
