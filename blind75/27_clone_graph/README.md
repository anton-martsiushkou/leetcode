# Clone Graph

## Problem Description

Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph.

Each node in the graph contains a value (`int`) and a list (`List[Node]`) of its neighbors.

```
class Node {
    public int val;
    public List<Node> neighbors;
}
```

## Test Case Format

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with `val = 1`, the second node with `val = 2`, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with `val = 1`. You must return the copy of the given node as a reference to the cloned graph.

## Examples

**Example 1:**
```
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
```

**Example 2:**
```
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.
```

**Example 3:**
```
Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.
```

## Constraints

- The number of nodes in the graph is in the range `[0, 100]`.
- `1 <= Node.val <= 100`
- `Node.val` is unique for each node.
- There are no repeated edges and no self-loops in the graph.
- The Graph is connected and all nodes can be visited starting from the given node.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS with Hash Map | Hash Map + Recursion Stack | O(N + E) | O(N) |
| BFS with Hash Map | Hash Map + Queue | O(N + E) | O(N) |

### Approach 1: DFS with Hash Map

The key insight is to use a hash map to keep track of already cloned nodes to avoid cycles and duplicate work. We traverse the graph using DFS, cloning each node and its neighbors recursively.

**Key Insight:** We need to track cloned nodes in a hash map to handle cycles in the graph. When we encounter a node we've already cloned, we return the clone from the map instead of creating a duplicate.

### Algorithm Steps (DFS)

1. Create a hash map to store `{original_node: cloned_node}` mappings
2. If the input node is null, return null
3. If the current node is already in the hash map, return the cloned node
4. Create a clone of the current node (with the same value but empty neighbors list)
5. Add the clone to the hash map
6. Recursively clone all neighbors and add them to the clone's neighbors list
7. Return the cloned node

### Approach 2: BFS with Hash Map

Similar to DFS, but uses a queue to traverse the graph level by level.

### Algorithm Steps (BFS)

1. If the input node is null, return null
2. Create a hash map and a queue
3. Clone the starting node and add it to the map
4. Add the original starting node to the queue
5. While the queue is not empty:
   - Dequeue a node
   - For each neighbor of the dequeued node:
     - If the neighbor hasn't been cloned, clone it and add to map and queue
     - Add the cloned neighbor to the current cloned node's neighbors list
6. Return the cloned starting node from the map

### Example Walkthrough (DFS)

For the graph `[[2,4],[1,3],[2,4],[1,3]]`:

1. Start with node 1, create clone, map: `{1: Node(1)}`
2. Process neighbor 2, create clone, map: `{1: Node(1), 2: Node(2)}`
3. Process neighbor 1 of node 2 - already cloned, use from map
4. Process neighbor 3, create clone, map: `{1: Node(1), 2: Node(2), 3: Node(3)}`
5. Continue until all nodes and edges are cloned

### Why This is Optimal

- **Time Complexity O(N + E)**: We visit each node once (N nodes) and each edge once (E edges)
- **Space Complexity O(N)**: Hash map stores N nodes, recursion stack can be O(N) in worst case (chain)
- This is optimal because we must visit every node and edge at least once to create a complete clone
