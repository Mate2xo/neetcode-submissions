import "slices"

func eraseOverlapIntervals(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		previous, current := result[len(result)-1], intervals[i]
		if current[0] < previous[1] {
			var shortest []int
			if previous[1] < current[1] {
				shortest = previous
			} else {
				shortest = current
			}
			result[len(result)-1] = shortest
		} else {
			result = append(result, current)
		}
	}
	return len(intervals) - len(result)
}
