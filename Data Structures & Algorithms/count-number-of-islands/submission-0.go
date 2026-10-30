func numIslands(grid [][]byte) int {
	islandCount, visited := 0, map[[2]int]bool{}
	ROWS, COLS := len(grid)-1, len(grid[0])-1
	var seekLand func([2]int) bool
	seekLand = func(position [2]int) bool {
		outOfBounds := position[0] < 0 || position[1] < 0 || position[0] > ROWS || position[1] > COLS
		if outOfBounds {
			return false
		}
		if visited[position] {
			return false
		}

		visited[position] = true
		if grid[position[0]][position[1]] == '0' {
			return false
		}

		up, down := [2]int{position[0] - 1, position[1]}, [2]int{position[0] + 1, position[1]}
		right, left := [2]int{position[0], position[1] + 1}, [2]int{position[0], position[1] - 1}
		seekLand(up)
		seekLand(right)
		seekLand(down)
		seekLand(left)

		return true
	}

	for rowIdx, row := range grid {
		for colIdx := range row {
			position := [2]int{rowIdx, colIdx}
			if visited[position] {
				continue
			}

			if seekLand(position) {
				islandCount++
			}
		}
	}
	return islandCount
}
