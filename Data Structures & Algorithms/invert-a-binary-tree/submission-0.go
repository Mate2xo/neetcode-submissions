/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	var current *TreeNode

	for len(queue) > 0 {
		current, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if current == nil {
			continue
		}
		queue = append(queue, current.Left, current.Right)
		current.Left, current.Right = current.Right, current.Left
	}
	return root
}
