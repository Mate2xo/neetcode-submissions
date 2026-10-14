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
	if root.Left == nil && root.Right == nil { 
		return root.Val
	}

	currentMax := root.Val
	maxLeft := maxPathSum(root.Left)
	currentMax = max(currentMax+maxLeft, currentMax)
	maxRight := maxPathSum(root.Right)
	currentMax = max(currentMax+maxRight, currentMax)

	return max(currentMax, maxLeft, maxRight)
}
