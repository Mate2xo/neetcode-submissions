/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(a, b int) bool { return h[a].Val < h[b].Val }
func (h MinHeap) Swap(a, b int)      { h[a], h[b] = h[b], h[a] }
func (h *MinHeap) Push(node any)     { *h = append(*h, node.(*ListNode)) }
func (h *MinHeap) Pop() any {
	old := *h
	lastIdx := len(old) - 1
	last := old[lastIdx]
	*h = old[:lastIdx]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, list := range lists {
		if list == nil {
			continue
		}

		for list != nil {
			heap.Push(minHeap, list)
			list = list.Next
		}
	}
	if minHeap.Len() == 0 {
		return nil
	}

	placeholder := &ListNode{}
	current := placeholder
	for minHeap.Len() > 0 {
		current.Next = heap.Pop(minHeap).(*ListNode)
		current = current.Next
	}

	return placeholder.Next
}


// Iterative Divide & Conquer(O(nlogk))
//
// func mergeKLists(lists []*ListNode) *ListNode {
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	if len(lists) == 1 {
// 		return lists[0]
// 	}

// 	for len(lists) > 1 {
// 		mergedLists := []*ListNode{}
// 		for i := 0; i+1 < len(lists); i += 2 {
// 			mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
// 		}
// 		if len(lists)%2 == 1 {
// 			mergedLists = append(mergedLists, lists[(len(lists)-1)])
// 		}
// 		lists = mergedLists
// 	}
// 	return lists[0]
// }

// func mergeTwoLists(list1, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val <= list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	}

// 	list2.Next = mergeTwoLists(list1, list2.Next)
// 	return list2
// }

// Brute force recursive (O(nlogk))
// func mergeKLists(lists []*ListNode) *ListNode {
// 	lists = filterNils(lists)
// 	if len(lists) == 0 {
// 		return nil
// 	}
// 	var smallestIdxs []int
// 	smallestVal := math.MaxInt
// 	for idx, list := range lists {
// 		if list.Val < smallestVal {
// 			smallestVal = list.Val
// 			smallestIdxs = []int{}
// 			smallestIdxs = append(smallestIdxs, idx)
// 		} else if list.Val == smallestVal {
// 			smallestIdxs = append(smallestIdxs, idx)
// 		}
// 	}
// 	smallestHead := lists[smallestIdxs[0]]
// 	lists[smallestIdxs[0]] = lists[smallestIdxs[0]].Next
// 	current := smallestHead
// 	for i := 1; i < len(smallestIdxs); i++ {
// 		current.Next = lists[smallestIdxs[i]]
// 		current = current.Next
// 		lists[smallestIdxs[i]] = lists[smallestIdxs[i]].Next
// 	}
// 	current.Next = mergeKLists(lists)
// 	return smallestHead
// }
// func filterNils(lists []*ListNode) []*ListNode {
// 	end := 0
// 	for _, list := range lists {
// 		if list != nil {
// 			lists[end] = list
// 			end++
// 		}
// 	}
// 	lists = lists[:end]
// 	return lists
// }