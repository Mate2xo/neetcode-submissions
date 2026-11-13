import "slices"

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := map[int][]int{}
	for _, prereq := range prerequisites {
		course, required := prereq[0], prereq[1]
		if course == required || slices.Contains(graph[course], required) {
			return false
		}
		graph[required] = append(graph[required], course)
	}

	return true
}
