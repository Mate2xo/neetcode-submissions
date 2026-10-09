/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isSubtree(root, subRoot *TreeNode) bool {
if isSameTree(root, subRoot) {
		return true
	}
	leftHasSubtree := isSameTree(root.Left, subRoot)
	rightHasSubtree := isSameTree(root.Right, subRoot)

	return leftHasSubtree || rightHasSubtree
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	sameLevelHere := p.Val == q.Val

	return sameLevelHere && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
