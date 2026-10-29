/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */
import "slices"

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}
	slices.SortFunc(intervals, func(a, b Interval) int { return a.start - b.start })


	queue := []Interval{intervals[0]}
	concurrentRooms, minimumRooms := len(queue), len(queue)
	for i := 1; i < len(intervals); i++ {
		currentMeeting := intervals[i]
		for _, previousMeeting := range queue {
			if currentMeeting.start < previousMeeting.end {
				continue
			} else {
				queue = queue[1:]
				concurrentRooms += -1
			}
		}
		concurrentRooms += 1
		minimumRooms = max(concurrentRooms, minimumRooms)
		queue = append(queue, currentMeeting)
	}

	return minimumRooms
}
