/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(filterNils(lists)) == 0 {
		return nil
	}

	var smallestIdx int
	smallestVal := math.MaxInt
	for idx, list := range lists {
		if list == nil {
			continue
		}
		if list.Val < smallestVal {
			smallestVal = list.Val
			smallestIdx = idx
		}
	}
	smallestHead := lists[smallestIdx]
	lists[smallestIdx] = lists[smallestIdx].Next

	smallestHead.Next = mergeKLists(filterNils(lists))

	return smallestHead
}

func filterNils(lists []*ListNode) []*ListNode {
	end := 0
	for _, list := range lists {
		if list != nil {
			lists[end] = list
			end++
		}
	}
	lists = lists[:end]
	return lists
}
