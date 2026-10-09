/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


import "slices"

func isSubtree(root, subRoot *TreeNode) bool {
	commonRoots := searchNodes(root, subRoot.Val)
	if len(commonRoots) == 0 {
		return false
	}
	for _, commonRoot := range commonRoots {
		if isSameTree(commonRoot, subRoot) {
			return true
		}
	}
	return false
}

func searchNodes(root *TreeNode, val int) []*TreeNode {
	matches := []*TreeNode{}
	if root == nil {
		return matches
	}
	if root.Val == val {
		matches = append(matches, root)
	}
	left := searchNodes(root.Left, val)
	right := searchNodes(root.Right, val)
	matches = slices.Concat(matches, left, right)
	return matches
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
