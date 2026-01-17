package clone_graph

// Node represents a node in an undirected graph
type Node struct {
	Val       int
	Neighbors []*Node
}

// CloneGraph creates a deep copy of a graph using DFS
// Time: O(N + E), Space: O(N)
func CloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)
	return dfs(node, cloned)
}

func dfs(node *Node, cloned map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}

	// If already cloned, return the clone
	if clone, exists := cloned[node]; exists {
		return clone
	}

	// Create a new clone
	clone := &Node{Val: node.Val, Neighbors: []*Node{}}
	cloned[node] = clone

	// Clone all neighbors
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, dfs(neighbor, cloned))
	}

	return clone
}

// CloneGraphBFS creates a deep copy using BFS
// Time: O(N + E), Space: O(N)
func CloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloned := make(map[*Node]*Node)
	queue := []*Node{node}

	// Clone the starting node
	cloned[node] = &Node{Val: node.Val, Neighbors: []*Node{}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range current.Neighbors {
			if _, exists := cloned[neighbor]; !exists {
				// Clone the neighbor
				cloned[neighbor] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				queue = append(queue, neighbor)
			}
			// Add the cloned neighbor to current clone's neighbors
			cloned[current].Neighbors = append(cloned[current].Neighbors, cloned[neighbor])
		}
	}

	return cloned[node]
}
