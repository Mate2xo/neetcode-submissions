func eraseOverlapIntervals(intervals [][]int) int {
	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		previous, current := result[len(result)-1], intervals[i]
		if current[0] < previous[1] {
			result[len(result)-1] = []int{min(previous[0], current[0]), min(previous[1], current[1])}
		} else {
			result = append(result, current)
		}
	}
	return len(intervals) - len(result)
}
