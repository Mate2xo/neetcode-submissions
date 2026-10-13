/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}
	if root == nil {
		return levels
	}

	bfsQueue := []*TreeNode{root}
	for currentLevel := 0; len(bfsQueue) > 0; currentLevel++ {
		currentLevelNodes := bfsQueue
		bfsQueue = []*TreeNode{}
		for _, node := range currentLevelNodes {
			if node == nil {
				continue
			}
			
			if len(levels) == currentLevel {
				levels = append(levels, []int{})
			}
			levels[currentLevel] = append(levels[currentLevel], node.Val)
			bfsQueue = append(bfsQueue, node.Left, node.Right)
		}
	}

	return levels

}
