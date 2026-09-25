func twoSum(nums []int, target int) []int {
	sortedWithIdx := make([][2]int, len(nums))
	for i, num := range nums {
		sortedWithIdx[i] = [2]int{num, i}
	}

	sort.Slice(sortedWithIdx, func(i, j int) bool {
		return sortedWithIdx[i][0] < sortedWithIdx[j][0]
	})

	left, right := 0, len(nums)-1
	for left < right {
		sum := sortedWithIdx[left][0] + sortedWithIdx[right][0]
		if sum == target {
			if sortedWithIdx[left][1] < sortedWithIdx[right][1] {
				return []int{sortedWithIdx[left][1], sortedWithIdx[right][1]}
			} else {
				return []int{sortedWithIdx[right][1], sortedWithIdx[left][1]}
			}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}

	return []int{}
}
