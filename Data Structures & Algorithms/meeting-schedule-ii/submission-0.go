/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func minMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

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
