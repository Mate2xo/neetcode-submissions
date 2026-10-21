type TrieNode struct {
	Children  map[rune]*TrieNode
	EndOfWord bool
}

type PrefixTree struct {
	Root *TrieNode
}

func Constructor() PrefixTree {
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

func (trie *PrefixTree) Search(word string) bool {
	currentLevel := trie.Root
	for _, letter := range word {
		if _, found := currentLevel.Children[letter]; !found {
			return false
		}
		currentLevel = currentLevel.Children[letter]
	}

	return currentLevel.EndOfWord
}

func (trie *PrefixTree) StartsWith(prefix string) bool {
	currentLevel := trie.Root
	for _, letter := range prefix {
		if _, found := currentLevel.Children[letter]; !found {
			return false
		}
		currentLevel = currentLevel.Children[letter]
	}
	return true
}


func findWords(board [][]byte, words []string) []string {
	found := []string{}
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			visited := map[[2]int]bool{}
			position := [2]int{row, col}
			isInTrie := checkCell(board, position, "", &visited, trie, &found)
			if !isInTrie {
				continue
			}
		}
	}
	return found
}

func checkCell(board [][]byte, position [2]int, accumulatedWord string, visited *map[[2]int]bool, trie PrefixTree, found *[]string) bool {
	if (*visited)[position] {
		return false
	}
	currentLetter := board[position[0]][position[1]]
	currentWord := accumulatedWord + string(currentLetter)
	if !trie.StartsWith(currentWord) {
		return false
	}

	(*visited)[position] = true

	if trie.Search(currentWord) {
		*found = append(*found, currentWord)
	}

	if position[0] > 0 {
		up := [2]int{position[0] - 1, position[1]}
		checkCell(board, up, currentWord, visited, trie, found)
	}
	if position[1] < len(board)-1 {
		right := [2]int{position[0], position[1] + 1}
		checkCell(board, right, currentWord, visited, trie, found)
	}
	if position[0] < len(board)-1 {
		down := [2]int{position[0] + 1, position[1]}
		checkCell(board, down, currentWord, visited, trie, found)
	}
	if position[1] > 0 {
		left := [2]int{position[0], position[1] - 1}
		checkCell(board, left, currentWord, visited, trie, found)
	}
	return true
}
