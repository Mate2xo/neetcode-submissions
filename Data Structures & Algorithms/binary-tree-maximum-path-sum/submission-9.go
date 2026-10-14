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
		return 0
	}

	maxLeft := max(memoizeMax(node.Left, maxPath), 0)
	maxRight := max(memoizeMax(node.Right, maxPath), 0)

	currentMax := node.Val + maxLeft + maxRight
	*maxPath = max(*maxPath, currentMax)
	return node.Val + max(maxLeft, maxRight)
}