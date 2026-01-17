# Invert Binary Tree

## Problem Description

Given the `root` of a binary tree, invert the tree, and return its root.

Inverting a binary tree means swapping the left and right children of all nodes in the tree.

## Examples

**Example 1:**
```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
Explanation:
Before:              After:
      4                 4
     / \               / \
    2   7             7   2
   / \ / \           / \ / \
  1  3 6  9         9  6 3  1
```

**Example 2:**
```
Input: root = [2,1,3]
Output: [2,3,1]
Explanation:
Before:    After:
    2         2
   / \       / \
  1   3     3   1
```

**Example 3:**
```
Input: root = []
Output: []
```

## Constraints

- The number of nodes in the tree is in the range `[0, 100]`
- `-100 <= Node.val <= 100`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree | O(n) | O(h) |
| BFS Iterative | Binary Tree + Queue | O(n) | O(w) |

### Approach: DFS Recursive (Recommended)

The recursive approach naturally handles the tree inversion by:
1. Swapping the left and right children of each node
2. Recursively inverting the left subtree
3. Recursively inverting the right subtree

**Key Insight:** The inversion operation is naturally recursive - to invert a tree, swap its children and recursively invert each subtree.

### Algorithm Steps

1. Base case: If the node is `null`, return `null`
2. Swap the left and right children of the current node
3. Recursively invert the left subtree (which is now the original right subtree)
4. Recursively invert the right subtree (which is now the original left subtree)
5. Return the current node

### Example Walkthrough

For the tree `[4,2,7,1,3,6,9]`:

```
invert(4):
  ├─ swap children: left=7, right=2
  ├─ invert(7):
  │    ├─ swap children: left=9, right=6
  │    ├─ invert(9) = 9 (leaf)
  │    └─ invert(6) = 6 (leaf)
  │    return 7
  └─ invert(2):
       ├─ swap children: left=3, right=1
       ├─ invert(3) = 3 (leaf)
       └─ invert(1) = 1 (leaf)
       return 2
  return 4
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node exactly once to swap its children
- **Space Complexity O(h)**: The recursion stack depth equals the height of the tree
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)
- The recursive solution is simple, elegant, and modifies the tree in-place

### Alternative Approaches

1. **BFS Iterative**: Use a queue to process nodes level by level
   - Same time complexity O(n)
   - Space complexity O(w) where w is maximum width
   - Avoids recursion stack

2. **DFS Iterative**: Use a stack instead of recursion
   - Same time and space complexity as recursive approach
   - More code but explicit stack management
