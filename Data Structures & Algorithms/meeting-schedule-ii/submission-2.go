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

	starts, ends, usedRooms := []int{}, []int{}, 0
	for _, meeting := range intervals {
		starts, ends = append(starts, meeting.start), append(ends, meeting.end)
		slices.Sort(starts)
		slices.Sort(ends)
	}

	start, end := 0, 0
	for start < len(intervals) {
		meetingStart, meetingEnd := starts[start], ends[end]

		if meetingStart < meetingEnd {
			usedRooms++
			start++
		} else {
			usedRooms--
			end++
		}

		maxRooms = max(usedRooms, maxRooms)
	}

	return maxRooms
}
