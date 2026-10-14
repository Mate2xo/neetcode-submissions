/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt
	memoizeMax(root, &maxPath)
	return maxPath
}

func memoizeMax(node *TreeNode, maxPath *int) int {
	if node == nil {
		return math.MinInt
	}

	currentMax := node.Val
	
	maxLeft := memoizeMax(node.Left, maxPath)
	if maxLeft != math.MinInt {
		currentMax += max(maxLeft, 0)
		// currentMax = max(currentMax+maxLeft, currentMax)
	}
	maxRight := memoizeMax(node.Right, maxPath)
	if maxRight != math.MinInt {
		currentMax += max(maxRight, 0)
		// currentMax = max(currentMax+maxRight, currentMax)
	}

	if currentMax > *maxPath {
		*maxPath = currentMax
	}
	return currentMax
}