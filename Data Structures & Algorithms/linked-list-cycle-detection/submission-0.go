/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	fastP, slowP := head, head

	for i := 1; slowP != nil && fastP != nil; i++ {
		slowP = slowP.Next

		for j := 0; j < 2*i; j++ {
			fastP = fastP.Next
			if fastP == nil {
				return false
			}
			if fastP.Val < slowP.Val {
				return true
			}
		}
	}
	return false
}
