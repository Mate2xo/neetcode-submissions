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
			sum := first + nums[secondIdx] + nums[thirdIdx]
			if sum == 0 {
				result = append(result, []int{first, nums[secondIdx], nums[thirdIdx]})
				secondIdx++
				for secondIdx < thirdIdx && nums[secondIdx] == nums[secondIdx-1] {
					secondIdx++
				}
				thirdIdx--
			} else if sum > 0 {
				thirdIdx--
			} else if sum < 0 {
				secondIdx++
			}
		}
	}

	return result
}
