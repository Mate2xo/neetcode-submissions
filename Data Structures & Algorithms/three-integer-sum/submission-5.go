import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := [][]int{}

	for firstIdx := 0; firstIdx < len(nums)-2; firstIdx++ {
		first := nums[firstIdx]
		if first > 0 {
			// every folowing sums will be > 0
			break
		}
		if firstIdx > 0 && first == nums[firstIdx-1] {
			// skip duplicate value
			continue
		}

		secondIdx, thirdIdx := firstIdx+1, len(nums)-1
		for secondIdx < thirdIdx {
			second, third := nums[secondIdx], nums[thirdIdx]
			if (first + second + third) == 0 {
				result = append(result, []int{first, second, third})
				secondIdx++
				for secondIdx < thirdIdx && nums[secondIdx] == nums[secondIdx-1] {
					secondIdx++
				}
				thirdIdx--
			} else if first+second+third > 0 {
				thirdIdx--
			} else if first+second+third < 0 {
				secondIdx++
			}
		}
	}

	return result
}
