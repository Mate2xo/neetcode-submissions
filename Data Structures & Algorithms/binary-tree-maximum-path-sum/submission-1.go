/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	currentMax, maxLeft, maxRight := root.Val, math.MinInt, math.MinInt

	if root.Left != nil {
		maxLeft = maxPathSum(root.Left)
		currentMax = max(currentMax + maxLeft, currentMax)
	}
	if root.Right != nil {
		maxRight = maxPathSum(root.Right)
		currentMax = max(currentMax + maxRight, currentMax)
	}

	return max(currentMax, maxLeft, maxRight)
}
