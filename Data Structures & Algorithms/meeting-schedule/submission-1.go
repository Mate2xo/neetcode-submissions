/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

import "slices"

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) == 0 {
		return true
	}

	slices.SortFunc(intervals, func(a, b Interval) int { return a.start - b.start })
	previousMeeting := intervals[0]
	for i := 1; i < len(intervals); i++ {
		currentMeeting := intervals[i]
		if currentMeeting.start < previousMeeting.end {
			return false
		}

		previousMeeting = currentMeeting
	}
	return true
}
