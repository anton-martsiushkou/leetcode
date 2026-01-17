# Graph Valid Tree

## Problem Description

Given `n` nodes labeled from `0` to `n - 1` and a list of undirected edges (each edge is a pair of nodes), write a function to check whether these edges make up a valid tree.

A valid tree must satisfy two conditions:
1. The graph must be fully connected (all nodes are reachable from any node)
2. The graph must have no cycles

## Examples

**Example 1:**
```
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true
Explanation: The graph forms a valid tree structure.
```

**Example 2:**
```
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false
Explanation: The graph contains a cycle: 1 -> 2 -> 3 -> 1
```

## Constraints

- `1 <= n <= 2000`
- `0 <= edges.length <= 5000`
- `edges[i].length == 2`
- `0 <= ai, bi < n`
- `ai != bi`
- There are no duplicate edges

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Union-Find | Disjoint Set Union | O(E * α(n)) ≈ O(E) | O(n) |
| DFS | Graph + Visited Set | O(E + n) | O(n) |
| BFS | Graph + Queue | O(E + n) | O(n) |

### Approach 1: Graph Theory Properties

A valid tree with `n` nodes must have exactly `n - 1` edges and be fully connected.

**Key Insight:** A tree with n nodes has exactly n-1 edges. If we have n-1 edges and the graph is connected, then by definition it cannot have cycles.

### Algorithm Steps

1. Check if number of edges equals `n - 1`
   - If not, return false (either has cycles or is disconnected)
2. Build adjacency list from edges
3. Use DFS/BFS to check if all nodes are reachable from node 0
4. Return true if all nodes visited, false otherwise

### Approach 2: Union-Find (Disjoint Set Union)

Use Union-Find to detect cycles while building the graph.

**Key Insight:** When adding an edge, if both nodes already belong to the same set, we've found a cycle.

### Algorithm Steps

1. Initialize Union-Find with `n` disjoint sets
2. For each edge `[u, v]`:
   - If `u` and `v` are already in the same set, return false (cycle detected)
   - Otherwise, union the sets containing `u` and `v`
3. After processing all edges, check if there's exactly one connected component
4. Return true if one component, false otherwise

### Example Walkthrough

For `n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]`:

**Check edge count:**
- Edges: 4, Expected: n-1 = 4 ✓

**DFS from node 0:**
- Visit 0 → neighbors: [1, 2, 3]
- Visit 1 → neighbors: [0, 4]
- Visit 4 → neighbors: [1]
- Visit 2 → neighbors: [0]
- Visit 3 → neighbors: [0]
- All 5 nodes visited ✓

**Result:** true

For `n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]`:

**Check edge count:**
- Edges: 5, Expected: n-1 = 4 ✗

**Result:** false (too many edges means there must be a cycle)

### Why This is Optimal

- **Time Complexity O(E)**: We process each edge once, and DFS/BFS visits each node once
- **Space Complexity O(n)**: For the adjacency list and visited set
- This is optimal because we must examine all edges to verify the graph structure

### Tree Properties

A graph is a valid tree if and only if:
1. It has exactly `n - 1` edges (where n is the number of nodes)
2. It is connected (all nodes are reachable)
3. These two conditions together guarantee no cycles

### Alternative Approaches

1. **Cycle Detection with DFS**: Track parent nodes to detect cycles - O(E + n) time
2. **Topological Sort**: Try to perform topological sort - only works for trees
