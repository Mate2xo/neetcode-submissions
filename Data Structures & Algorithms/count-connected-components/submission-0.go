func countComponents(n int, edges [][]int) int {
	count := 0
	adjacency := map[int][]int{}
	for _, neighbors := range edges {
		adjacency[neighbors[0]] = append(adjacency[neighbors[0]], neighbors[1])
		adjacency[neighbors[1]] = append(adjacency[neighbors[1]], neighbors[0])
	}

	visited := map[int]bool{}
	var dfs func(node, previous int)
	dfs = func(node, previous int) {
		if visited[node] {
			return
		}
		visited[node] = true

		for _, neighbor := range adjacency[node] {
			if node == previous {
				continue
			}
			dfs(neighbor, node)
		}
	}

	for node := range n {
		if !visited[node] {
			count++
		}
	dfs(node, -1)
	}
	return count

}
