/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	min := min(p.Val, q.Val)
	max := max(p.Val, q.Val)

	if min < root.Val && root.Val < max {
		return root
	}
	if min < root.Val && max < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if min > root.Val && max > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if (root.Val == min && max > root.Val) || (min < root.Val && root.Val == max) {
		return root
	}

	return root

}
