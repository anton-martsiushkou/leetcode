package valid_tree

// ValidTree checks if graph is a valid tree using DFS
// Time: O(V + E), Space: O(V + E)
func ValidTree(n int, edges [][]int) bool {
	// A tree with n nodes must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)

	var dfs func(node, parent int) bool
	dfs = func(node, parent int) bool {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if neighbor == parent {
				// Skip the edge we came from
				continue
			}
			if visited[neighbor] {
				// Found a cycle
				return false
			}
			if !dfs(neighbor, node) {
				return false
			}
		}
		return true
	}

	// Start DFS from node 0
	if !dfs(0, -1) {
		return false
	}

	// Check if all nodes were visited (graph is connected)
	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}

// UnionFind data structure
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in same set - cycle detected
	}

	// Union by rank
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

// ValidTreeUnionFind checks if graph is a valid tree using Union Find
// Time: O(V + E × α(V)), Space: O(V)
func ValidTreeUnionFind(n int, edges [][]int) bool {
	// A tree with n nodes must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	uf := NewUnionFind(n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		if !uf.Union(a, b) {
			// Cycle detected
			return false
		}
	}

	// With n-1 edges and no cycles, the graph must be connected
	return true
}
