/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
import "slices"

func minMeetingRooms(intervals []Interval) int {
	maxRooms := 0
	if len(intervals) == 0 {
		return maxRooms
	}

	meetings := [][]int{}
	for _, meeting := range intervals {
		meetings = append(meetings, []int{meeting.start, +1})
		meetings = append(meetings, []int{meeting.end, -1})
	}

	usedRooms := 0
	slices.SortFunc(meetings, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}

		return a[0] - b[0]
	})

	for _, meeting := range meetings {
		usedRooms += meeting[1]
		maxRooms = max(maxRooms, usedRooms)
	}

	return maxRooms
}
