func foreignDictionary(words []string) string {
	orderedLetters := ""
	if len(words) == 1 {
		return words[0]
	}

	adjacency := map[byte][]byte{}
	for i := 0; i < len(words)-1; i++ {
		if !deduceDirectedEdges(&adjacency, words[i], words[i+1]) {
			return orderedLetters
		}
	}
	fmt.Printf("adjacency: %+v\n", adjacency)

	sortedVertices := kahnTopologicalSort(adjacency)
	fmt.Printf("sortedVertices: %+v\n", sortedVertices)
	for _, letter := range sortedVertices {
		orderedLetters += string(letter)
	}
	return orderedLetters
}

func deduceDirectedEdges(list *map[byte][]byte, w1, w2 string) bool {
	for _, letter := range w2 {
		if _, found := (*list)[byte(letter)]; !found {
			(*list)[byte(letter)] = []byte{}
		}
	}
	i := 0
	for w1[i] == w2[i] && i < min(len(w1), len(w2))-1 {
		i++
	}
	if len(w1) > len(w2) && strings.HasPrefix(w1, w2) {
		fmt.Printf("w1 %q has prefix %q\n", w1, w2)
		return false
	}

	var previous, following byte
	if w1[i] != w2[i] {
		previous, following = w1[i], w2[i]
	} else if i+1 < len(w1) && i+1 < len(w2) {
		previous, following = w1[i+1], w2[i+1]
	} else if i+1 == len(w1) && i+1 < len(w2) {
		previous, following = w1[i], w2[i+1]
	} else {
		fmt.Printf("no case match, w1 %q w2 %q\n", w1, w2)
		return true
	}

	if previous != following {
		(*list)[w1[i]] = append((*list)[previous], following)
	}
	return true
}


func kahnTopologicalSort(list map[byte][]byte) []byte {
	verticesCount := len(list)
	inDegree := make(map[byte]int, verticesCount)
	for previousletter, followingLetters := range list {
		if _, found := inDegree[previousletter]; !found {
			inDegree[previousletter] = 0
		}
		for _, letter := range followingLetters {
			inDegree[letter]++
		}
	}
	fmt.Printf("inDegree: %+v\n", inDegree)

	queue := []byte{}
	for letter, degrees := range inDegree {
		if degrees == 0 {
			queue = append(queue, letter)
		}
	}
	topologicalOrder := []byte{}
	fmt.Printf("queue: %+v\n", queue)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Printf("current %v added\n", current)
		topologicalOrder = append(topologicalOrder, current)

		for _, neighbor := range list[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}

	}

	fmt.Printf("topologicalOrder %+v, list %+v\n", topologicalOrder, list)
	cycleDetected := len(topologicalOrder) != len(list)
	if cycleDetected {
		fmt.Printf("/!\\ cycleDetected\n")
		return nil
	}

	return topologicalOrder
}
