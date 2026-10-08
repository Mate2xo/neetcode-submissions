/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func isSubtree(root, subRoot *TreeNode) bool {
	commonRoot := searchNode(root, subRoot.Val)
	if commonRoot == nil || commonRoot.Val != subRoot.Val {
		return false
	}
	return isSameTree(commonRoot, subRoot)
}

func searchNode(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	if root == nil || root.Val > val {
		return nil
	}
	left := searchNode(root.Left, val)
	right := searchNode(root.Right, val)

	if left != nil {
		return left
	} else if right != nil {
		return right
	} else {
		return nil
	}
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
