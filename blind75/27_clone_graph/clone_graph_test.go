package clone_graph

import (
	"testing"
)

func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range adjList {
		nodes[i] = &Node{Val: i + 1, Neighbors: []*Node{}}
	}

	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes[0]
}

func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[int]bool)
	adjList := make([][]int, 0)
	var dfs func(*Node)

	dfs = func(n *Node) {
		if visited[n.Val] {
			return
		}
		visited[n.Val] = true

		// Ensure adjList has enough space
		for len(adjList) < n.Val {
			adjList = append(adjList, []int{})
		}

		neighbors := []int{}
		for _, neighbor := range n.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
			if !visited[neighbor.Val] {
				dfs(neighbor)
			}
		}
		adjList[n.Val-1] = neighbors
	}

	dfs(node)
	return adjList
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "example 1 - four nodes",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "example 2 - single node",
			adjList: [][]int{{}},
		},
		{
			name:    "example 3 - empty graph",
			adjList: [][]int{},
		},
		{
			name:    "two connected nodes",
			adjList: [][]int{{2}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := CloneGraph(original)

			// Verify the clone structure matches
			clonedAdjList := graphToAdjList(cloned)
			if len(clonedAdjList) != len(tt.adjList) {
				t.Errorf("CloneGraph() adjacency list length = %v, want %v", len(clonedAdjList), len(tt.adjList))
				return
			}

			for i := range tt.adjList {
				if len(clonedAdjList[i]) != len(tt.adjList[i]) {
					t.Errorf("CloneGraph() node %d neighbors = %v, want %v", i+1, clonedAdjList[i], tt.adjList[i])
				}
			}

			// Verify it's a deep copy (different memory addresses)
			if original != nil && cloned != nil && original == cloned {
				t.Error("CloneGraph() returned same reference, not a deep copy")
			}
		})
	}
}

func TestCloneGraphBFS(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
	}{
		{
			name:    "example 1 - four nodes",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "example 2 - single node",
			adjList: [][]int{{}},
		},
		{
			name:    "example 3 - empty graph",
			adjList: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			cloned := CloneGraphBFS(original)

			clonedAdjList := graphToAdjList(cloned)
			if len(clonedAdjList) != len(tt.adjList) {
				t.Errorf("CloneGraphBFS() adjacency list length = %v, want %v", len(clonedAdjList), len(tt.adjList))
			}

			if original != nil && cloned != nil && original == cloned {
				t.Error("CloneGraphBFS() returned same reference, not a deep copy")
			}
		})
	}
}
