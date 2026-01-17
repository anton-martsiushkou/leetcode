package connected_components

// CountComponents counts connected components using DFS
// Time: O(V + E), Space: O(V + E)
func CountComponents(n int, edges [][]int) int {
	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	count := 0

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	// Count components
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}

	return count
}

// CountComponentsBFS counts connected components using BFS
// Time: O(V + E), Space: O(V + E)
func CountComponentsBFS(n int, edges [][]int) int {
	// Build adjacency list
	graph := make([][]int, n)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	visited := make([]bool, n)
	count := 0

	bfs := func(start int) {
		queue := []int{start}
		visited[start] = true

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			for _, neighbor := range graph[node] {
				if !visited[neighbor] {
					visited[neighbor] = true
					queue = append(queue, neighbor)
				}
			}
		}
	}

	// Count components
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			bfs(i)
		}
	}

	return count
}

// UnionFind data structure
type UnionFind struct {
	parent     []int
	rank       []int
	components int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{parent: parent, rank: rank, components: n}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)

	if rootX == rootY {
		return // Already in same component
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

	uf.components--
}

// CountComponentsUnionFind counts connected components using Union Find
// Time: O(V + E × α(V)), Space: O(V)
func CountComponentsUnionFind(n int, edges [][]int) int {
	uf := NewUnionFind(n)

	for _, edge := range edges {
		a, b := edge[0], edge[1]
		uf.Union(a, b)
	}

	return uf.components
}
