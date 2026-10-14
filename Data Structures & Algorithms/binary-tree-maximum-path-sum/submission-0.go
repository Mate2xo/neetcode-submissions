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
	maxLeft := maxPathSum(root.Left)
	if currentMax+maxLeft > currentMax {
		currentMax = currentMax + maxLeft
	}
	maxRight := maxPathSum(root.Right)
	if currentMax+maxRight > currentMax {
		currentMax = currentMax + maxRight
	}

	return max(currentMax, maxLeft, maxRight)
}
