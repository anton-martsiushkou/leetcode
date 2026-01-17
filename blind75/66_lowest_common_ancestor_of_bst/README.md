# Lowest Common Ancestor of a Binary Search Tree

## Problem Description

Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: "The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow a node to be a descendant of itself)."

## Examples

**Example 1:**
```
        6
       / \
      2   8
     / \ / \
    0  4 7  9
      / \
     3   5

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: The LCA of nodes 2 and 8 is 6.
```

**Example 2:**
```
        6
       / \
      2   8
     / \ / \
    0  4 7  9
      / \
     3   5

Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself.
```

**Example 3:**
```
Input: root = [2,1], p = 2, q = 1
Output: 2
```

## Constraints

- The number of nodes in the tree is in the range `[2, 10^5]`
- `-10^9 <= Node.val <= 10^9`
- All `Node.val` are unique
- `p != q`
- `p` and `q` will exist in the BST

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| BST Property (Iterative) | Binary Search Tree | O(H) | O(1) |
| BST Property (Recursive) | Binary Search Tree | O(H) | O(H) |

### Approach: BST Property

The key insight is to leverage the BST property: all values in the left subtree are less than the root, and all values in the right subtree are greater than the root.

**Key Insight:** The LCA is the first node where the paths to `p` and `q` diverge. In a BST, this occurs when one value is less than or equal to the node and the other is greater than or equal to the node.

### Algorithm Steps

1. Start at the root node
2. For the current node:
   - If both `p` and `q` are less than current node's value, LCA must be in left subtree
   - If both `p` and `q` are greater than current node's value, LCA must be in right subtree
   - Otherwise, we've found the split point - current node is the LCA
3. Return the LCA node

### Example Walkthrough

For the tree with root = 6, p = 2, q = 8:

```
        6
       / \
      2   8
     / \ / \
    0  4 7  9
```

1. **Node 6**: 2 < 6 < 8
   - Values are on different sides, so 6 is the LCA
2. Return 6

For p = 2, q = 4:

1. **Node 6**: Both 2 < 6 and 4 < 6
   - Go left to node 2
2. **Node 2**: 2 == 2 and 4 > 2
   - Split point found, 2 is the LCA
3. Return 2

### Why This is Optimal

- **Time Complexity O(H)**: We traverse at most one path from root to a leaf, where H is the height (O(log n) for balanced tree, O(n) for skewed tree)
- **Space Complexity O(1)** for iterative, **O(H)** for recursive due to call stack
- This is optimal because in the worst case, the LCA could be a leaf node, requiring traversal of the entire height

### Why BST is Special

For a general binary tree (not BST), we would need:
- O(n) time to search for both nodes
- Track ancestors or use parent pointers

The BST property allows us to determine the search direction in O(1) time at each node.
