/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	stock := []*ListNode{}
	for head != nil {
		stock = append(stock, head)
		head = head.Next
	}

	start, end := 0, len(stock)-1
	newHead := &ListNode{}
	dummy := newHead
	for start < end {
		dummy.Next = stock[start]
		dummy = dummy.Next
		dummy.Next = stock[end]
		dummy = dummy.Next
		start++
		end--
	}

	if start == end {
		dummy.Next = stock[start]
		dummy = dummy.Next
	}
	dummy.Next = nil
}
