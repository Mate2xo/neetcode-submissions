/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
type Codec struct {
    
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	serialized, _ := json.Marshal(dfSearch(root))
	return string(serialized)
}

func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	var values []any
	json.Unmarshal([]byte(data), &values)

	newRoot := &TreeNode{}
	dfSet(newRoot, &values)
	return newRoot
}

func dfSearch(node *TreeNode) []any {
	visited := []any{node.Val}
	if node.Left != nil {
		visited = append(visited, dfSearch(node.Left)...)
	} else {
		visited = append(visited, nil)
	}
	if node.Right != nil {
		visited = append(visited, dfSearch(node.Right)...)
	} else {
		visited = append(visited, nil)
	}

	return visited
}

func dfSet(node *TreeNode, values *[]any) *TreeNode {
	node.Val = int((*values)[0].(float64))
	*values = (*values)[1:]

	if (*values)[0] != nil {
		node.Left = &TreeNode{}
		dfSet(node.Left, values)
	} else {
		*values = (*values)[1:]
	}

	if (*values)[0] != nil {
		node.Right = &TreeNode{}
		dfSet(node.Right, values)
	} else {
		*values = (*values)[1:]
	}

	return node
}
