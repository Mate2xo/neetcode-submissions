func countComponents(n int, edges [][]int) int {
	dsu := NewDisjointSetUnion(n)
	componentsCount := n
	for _, neighbors := range edges {
		if dsu.union(neighbors[0], neighbors[1]) {
			componentsCount--
		}
	}
	return componentsCount
}

func (dsu *DSU) union(n1, n2 int) bool {
	parent1 := dsu.findRoot(n1)
	parent2 := dsu.findRoot(n2)

	if parent1 == parent2 {
		return false
	}

	if dsu.size[parent1] > dsu.size[parent2] {
		dsu.parent[parent2] = parent1
		dsu.size[parent1] += dsu.size[parent2]
	} else {
		dsu.parent[parent1] = parent2
		dsu.size[parent2] += dsu.size[parent1]
	}

	return true
}

func (dsu *DSU) findRoot(node int) int {
	current := node
	for current != dsu.parent[current] {
		dsu.parent[current] = dsu.parent[dsu.parent[current]]
		current = dsu.parent[current]
	}
	return current
}

type DSU struct {
	parent []int
	size   []int
}

func NewDisjointSetUnion(n int) *DSU {
	dsu := &DSU{make([]int, n), make([]int, n)}
	for i := range n {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}
