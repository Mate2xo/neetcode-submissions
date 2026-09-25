func hasDuplicate(nums []int) bool {
	frequency := map[int]int{}
	for _, num := range nums {
		if _, ok := frequency[num]; ok {
			return true
		} else {
			frequency[num] = 1
		}
	}
	return false
}