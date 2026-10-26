import "slices"

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) == 0 {
		return intervals
	}
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	leftP, rightP := 0, len(intervals)-1
	for newInterval[0] > intervals[leftP][1] {
		leftP++
	}
	if leftP != 0 {
		leftP--
	}
	for newInterval[1] < intervals[rightP][0] {
		rightP--
	}
	if rightP != len(intervals)-1 {
		rightP++
	}

	var leftOverlap, rightOverlap bool
	if newInterval[0] <= intervals[leftP][1] {
		leftOverlap = true
	}
	if newInterval[1] >= intervals[rightP][0] {
		rightOverlap = true
	}

	if !leftOverlap && !rightOverlap {
		intervals = slices.Delete(intervals, leftP+1, rightP)
		intervals = slices.Insert(intervals, leftP+1, newInterval)
	}
	if leftOverlap && rightOverlap {
		intervals[rightP][0] = intervals[leftP][0]
		intervals[rightP][1] = max(intervals[rightP][1], newInterval[1])
		intervals = slices.Delete(intervals, leftP, rightP)
	} else if leftOverlap {
		intervals[leftP][1] = max(newInterval[1], intervals[leftP][1])
		intervals = slices.Delete(intervals, leftP+1, rightP)
	} else if rightOverlap {
		intervals[rightP][0] = intervals[leftP][0]
		// intervals[rightP][0] = max(newInterval[0], intervals[rightP][0])
		intervals = slices.Delete(intervals, leftP+1, rightP)

	}
	return intervals
}
