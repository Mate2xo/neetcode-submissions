import (
	"slices"
	"maps"
)


func findWords(board [][]byte, words []string) []string {
	search := NewBoardWordSearch(board, words)

	search.forEachCell(func(row, col int) bool {
		return search.fromCell(Position{row, col}, "")
	})

	return slices.Collect(maps.Keys(search.foundSet))
}

func (s *BoardWordSearch) fromCell(position Position, accumulatedWord string) bool {
	currentWord := accumulatedWord + string(s.board[position.row][position.col])
	if !s.trie.StartsWith(currentWord) {
		return false
	}

	if s.trie.Search(currentWord) && !s.foundSet[currentWord] {
		s.foundSet[currentWord] = true
	}

	s.visited[position] = true
	for _, neighbor := range s.neighboringPosistions(position) {
		s.fromCell(neighbor, currentWord)
	}
	s.visited[position] = false

	return true
}

func (s *BoardWordSearch) neighboringPosistions(position Position) []Position {
	neighbors := []Position{}
	canGoUp, up := position.row > 0, Position{row: position.row - 1, col: position.col}
	canGoRight, right := position.col < s.colCount-1, Position{position.row, position.col + 1}
	canGoDown, down := position.row < s.rowCount-1, Position{position.row + 1, position.col}
	canGoLeft, left := position.col > 0, Position{position.row, position.col - 1}

	if canGoUp && !s.visited[up] {
		neighbors = append(neighbors, up)
	}
	if canGoRight && !s.visited[right] {
		neighbors = append(neighbors, right)
	}
	if canGoDown && !s.visited[down] {
		neighbors = append(neighbors, down)
	}
	if canGoLeft && !s.visited[left] {
		neighbors = append(neighbors, left)
	}
	return neighbors
}

func (s BoardWordSearch) forEachCell(action func(row, col int) bool) {
	for row := range s.rowCount {
		for col := range s.colCount {
			action(row, col)
		}
	}
}

func NewBoardWordSearch(board [][]byte, words []string) BoardWordSearch {
	search := BoardWordSearch{
		board:    board,
		trie:     NewTrie(),
		visited:  map[Position]bool{},
		foundSet: map[string]bool{},
		rowCount: len(board),
		colCount: len(board[0]),
	}
	for _, word := range words {
		search.trie.Insert(word)
	}
	return search
}


func NewTrie() PrefixTree {
	return PrefixTree{Root: &TrieNode{Children: map[rune]*TrieNode{}}}
}

func (trie *PrefixTree) Insert(word string) {
	currentLevel := trie.Root
	for _, letter := range word {
		if _, found := currentLevel.Children[letter]; !found {
			currentLevel.Children[letter] = &TrieNode{Children: map[rune]*TrieNode{}}
		}
		currentLevel = currentLevel.Children[letter]
	}
	currentLevel.EndOfWord = true
}

func (trie PrefixTree) Search(word string) bool {
	currentLevel := trie.Root
	for _, letter := range word {
		if _, found := currentLevel.Children[letter]; !found {
			return false
		}
		currentLevel = currentLevel.Children[letter]
	}

	return currentLevel.EndOfWord
}

func (trie PrefixTree) StartsWith(prefix string) bool {
	currentLevel := trie.Root
	for _, letter := range prefix {
		if _, found := currentLevel.Children[letter]; !found {
			return false
		}
		currentLevel = currentLevel.Children[letter]
	}
	return true
}


type BoardWordSearch struct {
	board              [][]byte
	colCount, rowCount int
	foundSet           map[string]bool
	trie               PrefixTree
	visited            map[Position]bool
}
type Position struct {
	row int
	col int
}

type PrefixTree struct {
	Root *TrieNode
}
type TrieNode struct {
	Children  map[rune]*TrieNode
	EndOfWord bool
}