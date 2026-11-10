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
	clones := map[*Node]*Node{node: {node.Val, nil}}
	for len(breadthQueue) > 0 {
		currentNode := breadthQueue[0]
		breadthQueue = breadthQueue[1:]

		for _, neighbor := range currentNode.Neighbors {
			if _, exists := clones[neighbor]; !exists {
				clones[neighbor] = &Node{neighbor.Val, nil}
				breadthQueue = append(breadthQueue, neighbor)
			}
			clones[currentNode].Neighbors = append(clones[currentNode].Neighbors, clones[neighbor])
		}
	}

	return clones[node]
}
