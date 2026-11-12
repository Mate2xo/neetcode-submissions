func pacificAtlantic(heights [][]int) [][]int {
	ROWS, COLS := len(heights)-1, len(heights[0])-1
	result := [][]int{}

	for row := range len(heights) {
		for col, currentCell := range heights[row] {
			position := [2]int{row, col}
			flowsTo := map[string]bool{"pacific": false, "atlantic": false}
			visited := map[[2]int]bool{}

			var flow func([2]int, int, map[string]bool) bool
			flow = func(position [2]int, fromCell int, flowsTo map[string]bool) bool {
				row, col := position[0], position[1]
				outOfBounds := row < 0 || col < 0 || row > ROWS || col > COLS
				if outOfBounds || visited[position] || heights[row][col] > fromCell {
					return false
				}
				visited[position] = true

				if row == 0 || col == 0 {
					flowsTo["pacific"] = true
				}
				if row == ROWS || col == COLS {
					flowsTo["atlantic"] = true
				}
				if flowsTo["pacific"] && flowsTo["atlantic"] {
					return true
				}

				up, down, left, right := [2]int{row - 1, col}, [2]int{row + 1, col}, [2]int{row, col - 1}, [2]int{row, col + 1}
				directions := [][2]int{up, down, left, right}
				for _, direction := range directions {
					reached := flow(direction, heights[row][col], flowsTo)
					if reached {
						return true
					}
				}

				return false
			}

			if flow(position, currentCell, flowsTo) {
				result = append(result, []int{position[0], position[1]})
			}
		}
	}
	return result
}
