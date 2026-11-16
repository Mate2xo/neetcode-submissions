func validTree(n int, edges [][]int) bool {
	if len(edges) == 0 {
		return true
	}
	adjacency := map[int][]int{}
	for _, nodes := range edges {
		adjacency[nodes[0]] = append(adjacency[nodes[0]], nodes[1])
		adjacency[nodes[1]] = append(adjacency[nodes[1]], nodes[0])
	}

	visited := map[int]bool{}
	var dfsHasCycle func(node, previous int) bool
	dfsHasCycle = func(node, previous int) bool {
		if visited[node] {
			return true
		}

		visited[node] = true
		for _, neighbor := range adjacency[node] {
			if neighbor == previous {
				continue
			}
			if dfsHasCycle(neighbor, node) {
				return true
			}
		}
		visited[node] = false

		return false
	}

	if dfsHasCycle(0, -1) {
		return false
	}

	allNodesConnected := len(visited) == n
	return allNodesConnected
}
