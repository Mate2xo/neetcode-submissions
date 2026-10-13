/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	orderedValues := inOrder(root, &[]int{})
	return orderedValues[k-1]
}

func inOrder(node *TreeNode, values *[]int) []int {
	if node.Left != nil {
		inOrder(node.Left, values)
	}
	*values = append(*values, node.Val)
	if node.Right != nil {
		inOrder(node.Right, values)
	}
	return *values
}
