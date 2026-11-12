func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights), len(heights[0])
	var flowUphill func([2]int, int, map[[2]int]bool)
	flowUphill = func(position [2]int, previousHeight int, visited map[[2]int]bool) {
		row, col := position[0], position[1]
		outOfBounds := row < 0 || col < 0 || row >= ROWS || col >= COLS
		if outOfBounds || visited[position] || heights[row][col] < previousHeight {
			return
		}

		visited[position] = true
		up, down, left, right := [2]int{row + 1, col}, [2]int{row - 1, col}, [2]int{row, col + 1}, [2]int{row, col - 1}
		flowUphill(up, heights[row][col], visited)
		flowUphill(down, heights[row][col], visited)
		flowUphill(left, heights[row][col], visited)
		flowUphill(right, heights[row][col], visited)
	}

	reaches := map[string]map[[2]int]bool{"pacific": {}, "atlantic": {}}
	for col := range COLS {
		flowUphill([2]int{0, col}, heights[0][col], reaches["pacific"])
		flowUphill([2]int{ROWS - 1, col}, heights[ROWS-1][col], reaches["atlantic"])
	}
	for row := range ROWS {
		flowUphill([2]int{row, 0}, heights[row][0], reaches["pacific"])
		flowUphill([2]int{row, COLS - 1}, heights[row][COLS-1], reaches["atlantic"])
	}

	result := [][]int{}
	for row := range ROWS {
		for col := range COLS {
			position := [2]int{row, col}
			if reaches["pacific"][position] && reaches["atlantic"][position] {
				result = append(result, []int{row, col})
			}
		}
	}
	return result
}
