/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValid(root, math.MinInt, math.MaxInt)
}

func isValid(node *TreeNode, minParent, maxParent int) bool {
	if node == nil {
		return true
	}

	if !(minParent < node.Val && node.Val < maxParent) {
		return false
	}
	return isValid(node.Left, minParent, node.Val) && isValid(node.Right, node.Val, maxParent)
	
}
