/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	size := 1
	scan := head
	for scan.Next != nil {
		size++
		scan = scan.Next
	}

	if size == 1 {
		return nil
	}

	target := size - n
	scan = head
	for i := range target {
		if i < target-1 {
			scan = scan.Next
		} else {
			scan.Next = scan.Next.Next
		}
	}

	return head
}
