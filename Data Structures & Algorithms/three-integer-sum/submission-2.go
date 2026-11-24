import "slices"

func threeSum(nums []int) [][]int {
	sortedWithIndices := [][2]int{}
	resultSet := map[[3]int]bool{}
	for idx, num := range nums {
		sortedWithIndices = append(sortedWithIndices, [2]int{num, idx})
	}
	slices.SortFunc(sortedWithIndices, func(a [2]int, b [2]int) int {
		return a[0] - b[0]
	})

	for firstIdx := 0; firstIdx < len(nums)-2; firstIdx++ {
		first := sortedWithIndices[firstIdx][0]
		for secondIdx := firstIdx + 1; secondIdx < len(nums)-1; secondIdx++ {
			second := sortedWithIndices[secondIdx][0]
			thirdIdx := len(nums) - 1
			for thirdIdx != secondIdx {
				third := sortedWithIndices[thirdIdx][0]
				if (first + second + third) == 0 {
					resultSet[[3]int{first, second, third}] = true
					break
				} else if first+second+third > 0 {
					thirdIdx--
				} else {
					break
				}
			}
		}
	}

	result := [][]int{}
	for k := range resultSet {
		match := []int{k[0], k[1], k[2]}
		result = append(result, match)
	}
	return result
}
