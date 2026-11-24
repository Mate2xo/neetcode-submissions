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
		secondIdx, thirdIdx := firstIdx+1, len(nums)-1
		for secondIdx != thirdIdx {
			second, third := sortedWithIndices[secondIdx][0], sortedWithIndices[thirdIdx][0]
			if (first + second + third) == 0 {
				resultSet[[3]int{first, second, third}] = true
				secondIdx++
				thirdIdx = len(nums) - 1
			} else if first+second+third > 0 {
				thirdIdx--
			} else {
				secondIdx++
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
