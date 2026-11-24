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
		secondIdx, thirdIdx := firstIdx+1, len(nums)-1

		for secondIdx != thirdIdx {
			first, second, third := sortedWithIndices[firstIdx][0], sortedWithIndices[secondIdx][0], sortedWithIndices[thirdIdx][0]
			if (first + second + third) == 0 {
				// result = append(result, []int{nums[sortedWithIndices[firstIdx][1]], nums[sortedWithIndices[secondIdx][1]], nums[sortedWithIndices[thirdIdx][1]]})
				resultSet[[3]int{first, second, third}] = true
				break
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
