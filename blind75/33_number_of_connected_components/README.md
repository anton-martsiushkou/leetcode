# Number of Connected Components in Undirected Graph

## Problem Description

You have a graph of `n` nodes. You are given an integer `n` and an array `edges` where `edges[i] = [ai, bi]` indicates that there is an edge between `ai` and `bi` in the graph.

Return the number of connected components in the graph.

## Examples

**Example 1:**
```
Input: n = 5, edges = [[0,1],[1,2],[3,4]]
Output: 2
Explanation:
  0 --- 1       3
        |       |
        2       4

There are two connected components: {0, 1, 2} and {3, 4}
```

**Example 2:**
```
Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
Output: 1
Explanation:
  0 --- 1 --- 2 --- 3 --- 4

All nodes are connected in one component: {0, 1, 2, 3, 4}
```

## Constraints

- `1 <= n <= 2000`
- `1 <= edges.length <= 5000`
- `edges[i].length == 2`
- `0 <= ai <= bi < n`
- `ai != bi`
- There are no repeated edges.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS | Adjacency List + Visited Set | O(V + E) | O(V + E) |
| BFS | Adjacency List + Queue | O(V + E) | O(V + E) |
| Union Find | Disjoint Set Union | O(V + E × α(V)) | O(V) |

### Approach: Count Connected Components

We need to find the number of disjoint sets (connected components) in an undirected graph. We can use DFS, BFS, or Union Find to solve this problem.

**Key Insight:** Each connected component can be found by starting a DFS/BFS from an unvisited node and marking all reachable nodes as visited. The number of times we start a new DFS/BFS equals the number of connected components.

### Algorithm Steps (DFS Approach)

1. Build an adjacency list representation of the graph
2. Create a visited set to track visited nodes
3. Initialize component count to 0
4. For each node from 0 to n-1:
   - If the node is not visited:
     - Increment component count
     - Perform DFS to mark all nodes in this component as visited
5. Return the component count

### Algorithm Steps (Union Find)

1. Initialize Union Find with n nodes (each node is its own component initially)
2. For each edge [a, b]:
   - Union nodes a and b
3. Count the number of distinct root parents
4. Return the count

### Example Walkthrough

**Example 1:** `n = 5, edges = [[0,1],[1,2],[3,4]]`

**DFS Approach:**
1. Start with visited = {}
2. Node 0 not visited → component count = 1
   - DFS visits: 0 → 1 → 2
   - visited = {0, 1, 2}
3. Node 1 visited, skip
4. Node 2 visited, skip
5. Node 3 not visited → component count = 2
   - DFS visits: 3 → 4
   - visited = {0, 1, 2, 3, 4}
6. Node 4 visited, skip
7. Return 2

**Union Find Approach:**
1. Initialize: parent = [0, 1, 2, 3, 4], components = 5
2. Union(0, 1): parent = [0, 0, 2, 3, 4], components = 4
3. Union(1, 2): parent = [0, 0, 0, 3, 4], components = 3
4. Union(3, 4): parent = [0, 0, 0, 3, 3], components = 2
5. Return 2

### Why This is Optimal

- **Time Complexity**:
  - DFS/BFS: O(V + E) - Visit each vertex and edge once
  - Union Find: O(V + E × α(V)) where α is the inverse Ackermann function (nearly constant)

- **Space Complexity**:
  - DFS/BFS: O(V + E) for adjacency list and visited set/recursion stack
  - Union Find: O(V) for parent and rank arrays

- This is optimal because we must examine all edges at least once to determine connectivity

### Union Find Optimization

The Union Find approach is particularly elegant for this problem because:
- It naturally tracks the number of components
- Each successful union decreases the component count by 1
- No need to explicitly count roots at the end
