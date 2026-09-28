import "slices"

func longestConsecutive(nums []int) int {
	longest, currentMax := 1, 1
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return currentMax
	}
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if math.Abs(float64(nums[i-1]-nums[i])) == 1 {
			currentMax++
		} else if nums[i-1] == nums[i] {
			continue
		} else {
			currentMax = 1
		}
		setLongest(&currentMax, &longest)
	}

	return longest
}

func setLongest(current, longest *int) {
	if *current > *longest {
		*longest = *current
	}
}