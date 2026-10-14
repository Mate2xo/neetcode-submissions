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
		return math.MinInt
	}

	currentMax := root.Val

	maxLeft := maxPathSum(root.Left)
	if maxLeft != math.MinInt {
		currentMax = max(currentMax+maxLeft, currentMax)
	}
	maxRight := maxPathSum(root.Right)
	if maxRight != math.MinInt {
		currentMax = max(currentMax+maxRight, currentMax)
	}

	return max(currentMax, maxLeft, maxRight)
}