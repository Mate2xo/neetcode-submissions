import "slices"

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}

	for i, currentInterval := range intervals {
		if newInterval[1] < currentInterval[0] {
			result = append(result, newInterval)
			return slices.Concat(result, intervals[i:])
		} else if newInterval[0] > currentInterval[1] {
			result = append(result, currentInterval)
			continue
		}

		newInterval[0] = min(newInterval[0], currentInterval[0])
		newInterval[1] = max(newInterval[1], currentInterval[1])
	}
	result = append(result, newInterval)

	return result
}
