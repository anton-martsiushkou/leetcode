# Graph Valid Tree

## Problem Description

You have a graph of `n` nodes labeled from `0` to `n - 1`. You are given an integer `n` and a list of `edges` where `edges[i] = [ai, bi]` indicates that there is an undirected edge between nodes `ai` and `bi` in the graph.

Return `true` if the edges of the given graph make up a valid tree, and `false` otherwise.

## Examples

**Example 1:**
```
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true
```

**Example 2:**
```
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false
```

## Constraints

- `1 <= n <= 2000`
- `0 <= edges.length <= 5000`
- `edges[i].length == 2`
- `0 <= ai, bi < n`
- `ai != bi`
- There are no self-loops or repeated edges.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS with Cycle Detection | Adjacency List + Visited Set | O(V + E) | O(V + E) |
| Union Find | Disjoint Set Union | O(V + E × α(V)) | O(V) |
| BFS with Cycle Detection | Adjacency List + Queue | O(V + E) | O(V + E) |

### Approach: Graph Properties of a Tree

A valid tree must satisfy two conditions:
1. **Exactly n-1 edges**: A tree with n nodes has exactly n-1 edges
2. **No cycles**: The graph must be acyclic
3. **Connected**: All nodes must be reachable from any starting node

**Key Insight:** For a graph to be a valid tree with n nodes:
- It must have exactly n-1 edges (quick check)
- It must be fully connected (one component)
- It must have no cycles

We can verify this efficiently with DFS/BFS or Union Find.

### Algorithm Steps (DFS Approach)

1. **Quick check**: If edges.length ≠ n-1, return false immediately
2. Build an adjacency list representation
3. Perform DFS from node 0:
   - Keep track of visited nodes
   - Keep track of parent to avoid counting the edge we came from as a cycle
   - If we visit a node that's already visited and it's not our parent → cycle detected
4. After DFS, check if all n nodes were visited (graph is connected)
5. Return true if connected and no cycles

### Algorithm Steps (Union Find)

1. **Quick check**: If edges.length ≠ n-1, return false
2. Initialize Union Find with n nodes
3. For each edge [a, b]:
   - If a and b are already in the same set → cycle detected, return false
   - Otherwise, union them
4. After processing all edges, check if all nodes are in one component
5. Return true if one component and no cycles

### Example Walkthrough

**Example 1:** `n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]`
```
    0
   /|\
  1 2 3
  |
  4
```

1. Edge count: 4 = 5-1 ✓
2. DFS from 0: visits 0 → 1 → 4, then 2, then 3
3. All 5 nodes visited ✓
4. No cycles detected ✓
5. Return true

**Example 2:** `n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]`
```
  0-1-2
    |X|
    3-+
    |
    4
```

1. Edge count: 5 ≠ 5-1 ✗
2. Return false (also has cycle: 1-2-3-1)

### Why This is Optimal

- **Time Complexity O(V + E)**:
  - DFS/BFS: Visit each vertex once, examine each edge once
  - Union Find: Nearly O(V + E) with path compression and union by rank

- **Space Complexity O(V + E)**:
  - Adjacency list: O(V + E)
  - Visited set/recursion stack: O(V)

- This is optimal because we must examine all edges at least once to verify tree properties

### Tree Properties Summary

A valid tree with n nodes must have:
- **Exactly n-1 edges** (necessary but not sufficient)
- **No cycles**
- **Fully connected** (single component)
