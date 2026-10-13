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

	greaterLeft := root.Left != nil && root.Left.Val > root.Val
	smallerRight := root.Right != nil && root.Right.Val < root.Val
	if greaterLeft || smallerRight {
		return false
	}
	validLeft := isValidBST(root.Left)
	validRight := isValidBST(root.Right)
	return validLeft && validRight
}
