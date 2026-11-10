/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	breadthQueue := []*Node{node}
	clone := Node{node.Val, nil}
	clones := map[*Node]*Node{node: &clone}
	for len(breadthQueue) > 0 {
		currentNode := breadthQueue[0]
		breadthQueue = breadthQueue[1:]

		var currentClone *Node
		if _, found := clones[currentNode]; !found {
			currentClone = &Node{currentNode.Val, nil}
		} else {
			currentClone = clones[currentNode]
		}

		for _, neighbor := range currentNode.Neighbors {
			if _, exists := clones[neighbor]; exists {
				clones[currentNode].Neighbors = append(clones[currentNode].Neighbors, clones[neighbor])
				continue
			}
			clones[neighbor] = &Node{neighbor.Val, nil}
			currentClone.Neighbors = append(currentClone.Neighbors, clones[neighbor])
			breadthQueue = append(breadthQueue, neighbor)
		}
		fmt.Printf("current queue is %+v, currentClone %+v\n", breadthQueue, currentClone)
	}

	return &clone
}
