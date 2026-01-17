# Course Schedule

## Problem Description

There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course `bi` first if you want to take course `ai`.

For example, the pair `[0, 1]`, indicates that to take course `0` you have to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

## Examples

**Example 1:**
```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
```

**Example 2:**
```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.
```

## Constraints

- `1 <= numCourses <= 2000`
- `0 <= prerequisites.length <= 5000`
- `prerequisites[i].length == 2`
- `0 <= ai, bi < numCourses`
- All the pairs `prerequisites[i]` are unique.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Topological Sort (DFS) | Adjacency List + Visited Set | O(V + E) | O(V + E) |
| Topological Sort (Kahn's - BFS) | Adjacency List + Queue | O(V + E) | O(V + E) |

### Approach: Cycle Detection in Directed Graph

The key insight is that this problem is equivalent to detecting if there's a cycle in a directed graph. If there's a cycle, it's impossible to finish all courses. We can use topological sort to detect cycles.

**Key Insight:** A valid course schedule exists if and only if the prerequisite graph is a Directed Acyclic Graph (DAG). We need to detect if there's a cycle in the graph.

### Algorithm Steps (DFS Approach)

1. Build an adjacency list representation of the graph
2. Create a visited state tracker with three states:
   - 0 = unvisited
   - 1 = visiting (currently in DFS path)
   - 2 = visited (completely processed)
3. For each course, perform DFS:
   - If the course is being visited (state 1), we found a cycle → return false
   - If the course is already visited (state 2), skip it
   - Mark as visiting (1), explore all prerequisites
   - Mark as visited (2)
4. If no cycle is detected, return true

### Algorithm Steps (Kahn's Algorithm - BFS)

1. Build an adjacency list and calculate in-degree for each node
2. Add all nodes with in-degree 0 to a queue
3. While queue is not empty:
   - Remove a node from queue
   - For each neighbor, decrease its in-degree
   - If in-degree becomes 0, add to queue
4. If all nodes were processed, return true; otherwise false

### Example Walkthrough

For `numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]`:

**Graph representation:**
```
0 → 1 → 3
  → 2 →
```

**DFS Traversal:**
1. Start with course 0 (no prerequisites)
2. Visit course 1 (requires 0)
3. Visit course 2 (requires 0)
4. Visit course 3 (requires 1 and 2)
5. No cycles detected → return true

**For cycle case:** `prerequisites = [[1,0],[0,1]]`:
```
0 ⇄ 1 (cycle)
```
- Start with 0, mark as visiting
- Try to visit 1, mark as visiting
- Try to visit 0, but it's already visiting → cycle detected → return false

### Why This is Optimal

- **Time Complexity O(V + E)**: We visit each vertex once and each edge once
- **Space Complexity O(V + E)**: Adjacency list stores all edges, recursion stack or queue can store all vertices
- This is optimal because we must examine all courses and prerequisites at least once to determine if a valid ordering exists
