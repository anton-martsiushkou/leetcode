# Maximum Depth of Binary Tree

## Problem Description

Given the `root` of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Examples

**Example 1:**
```
Input: root = [3,9,20,null,null,15,7]
Output: 3
Explanation: The tree has three levels:
    3
   / \
  9  20
    /  \
   15   7
```

**Example 2:**
```
Input: root = [1,null,2]
Output: 2
Explanation: The tree has two levels:
  1
   \
    2
```

**Example 3:**
```
Input: root = []
Output: 0
```

**Example 4:**
```
Input: root = [0]
Output: 1
```

## Constraints

- The number of nodes in the tree is in the range `[0, 10^4]`
- `-100 <= Node.val <= 100`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree | O(n) | O(h) |
| BFS Iterative | Binary Tree + Queue | O(n) | O(w) |

### Approach 1: DFS Recursive (Recommended)

The recursive approach leverages the natural recursive structure of trees. The maximum depth of a tree is 1 (for the current node) plus the maximum of the depths of its left and right subtrees.

**Key Insight:** The problem has optimal substructure - the maximum depth of a tree can be computed from the maximum depths of its subtrees.

### Algorithm Steps

1. Base case: If the node is `null`, return 0
2. Recursively find the maximum depth of the left subtree
3. Recursively find the maximum depth of the right subtree
4. Return 1 + max(left_depth, right_depth)

### Example Walkthrough

For the tree `[3,9,20,null,null,15,7]`:

```
maxDepth(3):
  ├─ maxDepth(9) = 1 (leaf node)
  └─ maxDepth(20):
       ├─ maxDepth(15) = 1 (leaf node)
       └─ maxDepth(7) = 1 (leaf node)
       return 1 + max(1, 1) = 2
  return 1 + max(1, 2) = 3
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node exactly once
- **Space Complexity O(h)**: The recursion stack depth equals the height of the tree
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)
- The recursive solution is simple, elegant, and naturally matches the tree structure

### Approach 2: BFS Iterative (Alternative)

Use a queue to perform level-order traversal, incrementing depth counter for each level.

**Advantages:**
- No recursion stack overflow risk for very deep trees
- Space complexity is O(w) where w is maximum width, not height

**Steps:**
1. Use a queue, starting with root
2. Process nodes level by level
3. Increment depth counter for each complete level
4. Return final depth count
