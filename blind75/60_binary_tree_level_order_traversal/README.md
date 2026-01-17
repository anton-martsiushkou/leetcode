# Binary Tree Level Order Traversal

## Problem Description

Given the `root` of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Examples

**Example 1:**
```
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
Explanation:
      3
     / \
    9  20
      /  \
     15   7
Level 0: [3]
Level 1: [9, 20]
Level 2: [15, 7]
```

**Example 2:**
```
Input: root = [1]
Output: [[1]]
```

**Example 3:**
```
Input: root = []
Output: []
```

## Constraints

- The number of nodes in the tree is in the range `[0, 2000]`
- `-1000 <= Node.val <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| BFS Iterative | Binary Tree + Queue | O(n) | O(w) |
| DFS Recursive | Binary Tree + Array | O(n) | O(h) |

### Approach 1: BFS Iterative (Recommended)

The BFS (Breadth-First Search) approach uses a queue to process nodes level by level. This naturally matches the problem's requirement of level-order traversal.

**Key Insight:** A queue ensures we process all nodes at level k before processing any node at level k+1.

### Algorithm Steps

1. Handle edge case: If root is `null`, return empty list
2. Initialize a queue with the root node
3. Initialize result list to store levels
4. While queue is not empty:
   - Get the current level size (number of nodes in queue)
   - Create a list for current level values
   - For each node in the current level:
     - Remove node from queue
     - Add its value to current level list
     - Add its left child to queue (if exists)
     - Add its right child to queue (if exists)
   - Add current level list to result
5. Return result

### Example Walkthrough

For the tree `[3,9,20,null,null,15,7]`:

```
Initial: queue = [3], result = []

Level 0:
  queue size = 1
  process 3 → current_level = [3]
  add children: queue = [9, 20]
  result = [[3]]

Level 1:
  queue size = 2
  process 9 → current_level = [9]
  process 20 → current_level = [9, 20]
  add children: queue = [15, 7]
  result = [[3], [9, 20]]

Level 2:
  queue size = 2
  process 15 → current_level = [15]
  process 7 → current_level = [15, 7]
  result = [[3], [9, 20], [15, 7]]

Return: [[3], [9, 20], [15, 7]]
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node exactly once
- **Space Complexity O(w)**: The queue holds at most w nodes, where w is the maximum width of the tree
  - For a balanced tree: O(n/2) ≈ O(n) at the last level
  - For a skewed tree: O(1)
- BFS is the natural approach for level-order traversal and is more intuitive than DFS for this problem

### Approach 2: DFS Recursive (Alternative)

Use DFS with a level parameter to track the current depth, and add values to the appropriate level list.

**Steps:**
1. Pass current level as parameter in recursion
2. Ensure result list has a sublist for current level
3. Add current node's value to that level's list
4. Recursively visit left and right children with level + 1

**Advantages:**
- No explicit queue needed
- Can be more memory-efficient for very wide trees

**Disadvantages:**
- Less intuitive than BFS for level-order traversal
- Still requires O(h) recursion stack space
