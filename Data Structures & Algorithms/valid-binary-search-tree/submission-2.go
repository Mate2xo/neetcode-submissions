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

	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}

	return isValidLeft(root.Left, root) && isValidRight(root.Right, root)
}

func isValidLeft(node, parentNode *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil && (node.Left.Val >= node.Val || node.Left.Val >= parentNode.Val) {
		return false
	}
	if node.Right != nil && (node.Right.Val <= node.Val || node.Right.Val >= parentNode.Val) {
		return false
	}
	return isValidLeft(node.Left, node) && isValidRight(node.Right, node)
}

func isValidRight(node, parentNode *TreeNode) bool {
	if node == nil {
		return true
	}
	if node.Left != nil && (node.Left.Val >= node.Val || node.Left.Val <= parentNode.Val) {
		return false
	}
	if node.Right != nil && (node.Right.Val <= node.Val || node.Right.Val <= parentNode.Val) {
		return false
	}
	return isValidLeft(node.Left, node) && isValidRight(node.Right, node)
}
