type TreeNode struct {
	Val      int
	Parent   *TreeNode
	Children []*TreeNode
}

func validTree(n int, edges [][]int) bool {
	graph := make([]*TreeNode, n)
	for i := range n {
		graph[i] = &TreeNode{Val: i, Children: []*TreeNode{}}
	}
	for _, nodes := range edges {
		var parentVal, childVal int
		parentVal, childVal = nodes[0], nodes[1]

		graph[parentVal].Children = append(graph[parentVal].Children, graph[childVal])
		if graph[childVal].Parent != nil {
			return false
		}
		graph[childVal].Parent = graph[parentVal]
	}

	return true

}
