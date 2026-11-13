func canFinish(numCourses int, prerequisites [][]int) bool {
	adjacency := map[int][]int{}
	for num := range numCourses {
		adjacency[num] = []int{}
	}
	for _, prereq := range prerequisites {
		course, required := prereq[0], prereq[1]
		adjacency[course] = append(adjacency[course], required)
	}

	var dfs func(num int, visited map[int]bool) bool
	dfs = func(num int, visited map[int]bool) bool {
		if visited[num] {
			return false
		}

		for _, required := range adjacency[num] {
			visited[num] = true
			if !dfs(required, visited) {
				return false
			}
			visited[num] = false
		}
		return true
	}

	for num := range numCourses {
		if _, found := adjacency[num]; !found {
			continue
		}

		visited := map[int]bool{}
		if !dfs(num, visited) {
			return false
		}
	}
	return true
}
