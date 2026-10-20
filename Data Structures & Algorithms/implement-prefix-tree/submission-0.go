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
