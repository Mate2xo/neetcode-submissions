func pacificAtlantic(heights [][]int) [][]int {
	board := newOceanBoard(heights)
	result := [][]int{}

	for row := range len(heights) {
		for col, cell := range heights[row] {
			position := [2]int{row, col}
			reaches := map[string]bool{"pacific": false, "atlantic": false}
			visited := map[[2]int]bool{}

			if board.reachesOceansAt(position, cell, reaches, visited) {
				result = append(result, []int{position[0], position[1]})
			}
		}
	}
	return result
}

type OceanBoard struct {
	heights    [][]int
	ROWS, COLS int
}

func (board OceanBoard) reachesOceansAt(position [2]int, fromCell int, reaches map[string]bool, visited map[[2]int]bool) bool {
	row, col := position[0], position[1]
	outOfBounds := row < 0 || col < 0 || row > board.ROWS || col > board.COLS
	if outOfBounds ||
		visited[position] ||
		board.currentCell(position) > fromCell {
		return false
	}

	visited[position] = true
	if row == 0 || col == 0 {
		reaches["pacific"] = true
	}
	if row == board.ROWS || col == board.COLS {
		reaches["atlantic"] = true
	}
	if reaches["pacific"] && reaches["atlantic"] {
		return true
	}

	up, down, left, right := [2]int{row - 1, col}, [2]int{row + 1, col}, [2]int{row, col - 1}, [2]int{row, col + 1}
	directions := [][2]int{up, down, left, right}
	for _, direction := range directions {
		reached := board.reachesOceansAt(direction, board.currentCell(position), reaches, visited)
		if reached {
			return true
		}
	}

	return false
}

func (board OceanBoard) currentCell(position [2]int) int {
	row, col := position[0], position[1]
	return board.heights[row][col]
}

func newOceanBoard(heights [][]int) *OceanBoard {
	board := OceanBoard{ROWS: len(heights) - 1, COLS: len(heights[0]) - 1, heights: heights}
	return &board
}
