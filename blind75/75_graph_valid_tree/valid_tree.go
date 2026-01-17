package valid_tree

// ValidTree checks if the edges form a valid tree using DFS.
// Time: O(E + n), Space: O(n)
func ValidTree(n int, edges [][]int) bool {
	// A tree with n nodes must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	// Build adjacency list
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Check if all nodes are connected using DFS
	visited := make(map[int]bool)
	dfs(0, graph, visited)

	return len(visited) == n
}

func dfs(node int, graph map[int][]int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}

// UnionFind structure for cycle detection
type UnionFind struct {
	parent []int
	rank   []int
}

// NewUnionFind creates a new Union-Find structure
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &UnionFind{parent: parent, rank: rank}
}

// Find returns the root of the set containing x
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing x and y
// Returns false if they're already in the same set (cycle detected)
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return false // Already in same set - cycle detected
	}

	// Union by rank
	if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}

	return true
}

// ValidTreeUnionFind checks if the edges form a valid tree using Union-Find.
// Time: O(E * α(n)) ≈ O(E), Space: O(n)
func ValidTreeUnionFind(n int, edges [][]int) bool {
	// A tree with n nodes must have exactly n-1 edges
	if len(edges) != n-1 {
		return false
	}

	uf := NewUnionFind(n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		// If union returns false, we found a cycle
		if !uf.Union(u, v) {
			return false
		}
	}

	// If we have n-1 edges and no cycles, it's a valid tree
	return true
}
