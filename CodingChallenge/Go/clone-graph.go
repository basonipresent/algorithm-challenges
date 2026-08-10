package main

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * Clone Graph
 * Input: node = [[2,4],[1,3],[2,4],[1,3]]
 * Output: [[2,4],[1,3],[2,4],[1,3]]
 */
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	return nodeDFS(node, visited)
}

func nodeDFS(node *Node, visited map[*Node]*Node) *Node {
	if clone, ok := visited[node]; ok {
		return clone
	}

	newNode := &Node{Val: node.Val}
	visited[node] = newNode

	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, nodeDFS(neighbor, visited))
	}

	return newNode
}
