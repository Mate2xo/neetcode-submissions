/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	lists = filterNils(lists)
	if len(lists) == 0 {
		return nil
	}

	var smallestIdxs []int
	smallestVal := math.MaxInt
	for idx, list := range lists {
		if list.Val < smallestVal {
			smallestVal = list.Val
			smallestIdxs = []int{}
			smallestIdxs = append(smallestIdxs, idx)
		} else if list.Val == smallestVal {
			smallestIdxs = append(smallestIdxs, idx)
		}
	}

	smallestHead := lists[smallestIdxs[0]]
	lists[smallestIdxs[0]] = lists[smallestIdxs[0]].Next
	current := smallestHead
	for i := 1; i < len(smallestIdxs); i++ {
		current.Next = lists[smallestIdxs[i]]
		current = current.Next
		lists[smallestIdxs[i]] = lists[smallestIdxs[i]].Next
	}

	current.Next = mergeKLists(lists)

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
