/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	currentMax := root.Val
	if root.Left == nil && root.Right == nil { 
		return currentMax
	}

	maxLeft := max(maxPathSum(root.Left), 0)
	currentMax = currentMax + maxLeft
	maxRight := max(maxPathSum(root.Right), 0)
	currentMax = currentMax + maxRight

	return max(currentMax, maxLeft, maxRight)
}
