import "slices"

func longestConsecutive(nums []int) int {
	longest, currentMax := 0, 1
	if len(nums) == 0 {
		return longest
	}
	slices.Sort(nums)

	for i := 1; i < len(nums); i++ {
		if math.Abs(float64(nums[i-1]-nums[i])) == 1 {
			currentMax++
		} else if nums[i-1] == nums[i] {
			continue
		} else {
			currentMax = 0
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