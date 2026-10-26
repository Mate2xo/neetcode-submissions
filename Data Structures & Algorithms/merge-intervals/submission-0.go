import "slices"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	result, previous := [][]int{}, intervals[0]
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		if previous[1] < current[0] {
			result = append(result, previous)
			previous = current
			continue
		}
		previous = []int{min(previous[0], current[0]), max(previous[1], current[1])}
	}
	result = append(result, previous)
	return result

}
