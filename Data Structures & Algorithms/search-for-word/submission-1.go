func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return false
	}
	visited := map[[2]int]bool{}
	rows, cols := len(board), len(board[0])

	var dfsMatch func(word string, row, col int) bool
	dfsMatch = func(word string, row, col int) bool {
		if len(word) == 0 {
			return true
		}

		position := [2]int{row, col}
		outOfBoard := row < 0 || col < 0 || row >= rows || col >= cols
		if outOfBoard || visited[position] || board[row][col] != word[0] {
			return false
		}

		up, down, left, right := row-1, row+1, col+1, col-1
		if dfsMatch(word[1:], up, col) ||
			dfsMatch(word[1:], row, right) ||
			dfsMatch(word[1:], down, col) ||
			dfsMatch(word[1:], row, left) {
			return true
		}
		visited[position] = false
		return false
	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if dfsMatch(word, row, col) {
				return true
			}
		}
	}
	return false
}