type WordDictionaryNode struct {
	Children  map[rune]*WordDictionaryNode
	EndOfWord bool
}
type WordDictionary struct {
	root *WordDictionaryNode
}


func Constructor() WordDictionary {
	return WordDictionary{root: &WordDictionaryNode{Children: map[rune]*WordDictionaryNode{}}}
}

func (trie *WordDictionary) AddWord(word string) {
	currentLevel := trie.root
	for _, letter := range word {
		if _, found := currentLevel.Children[letter]; !found {
			currentLevel.Children[letter] = &WordDictionaryNode{Children: map[rune]*WordDictionaryNode{}}
		}
		currentLevel = currentLevel.Children[letter]
	}
	currentLevel.EndOfWord = true
}

func (trie WordDictionary) Search(word string) bool {
	if word == "" {
		return false
	}

	return trie.root.Search(word)
}

func (node WordDictionaryNode) Search(word string) bool {
	currentLevel := node
	for i, letter := range word {
		if letter == '.' && len(word) > 0 {
			return currentLevel.SearchChildren(word[i+1:])
		}

		if _, found := currentLevel.Children[letter]; !found {
			return false
		}
		currentLevel = *currentLevel.Children[letter]
	}
	return currentLevel.EndOfWord
}

func (node WordDictionaryNode) SearchChildren(word string) bool {
	for _, child := range node.Children {
		match := child.Search(word)
		if match {
			return true
		}
	}
	return false
}
