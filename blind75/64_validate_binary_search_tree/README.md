# Validate Binary Search Tree

## Problem Description

Given the root of a binary tree, determine if it is a valid binary search tree (BST).

A valid BST is defined as follows:
- The left subtree of a node contains only nodes with keys less than the node's key.
- The right subtree of a node contains only nodes with keys greater than the node's key.
- Both the left and right subtrees must also be binary search trees.

## Examples

**Example 1:**
```
    2
   / \
  1   3

Input: root = [2,1,3]
Output: true
```

**Example 2:**
```
    5
   / \
  1   4
     / \
    3   6

Input: root = [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
```

**Example 3:**
```
Input: root = []
Output: true
Explanation: An empty tree is a valid BST.
```

## Constraints

- The number of nodes in the tree is in the range `[0, 10^4]`
- `-2^31 <= Node.val <= 2^31 - 1`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Recursive with Range Validation | Binary Tree | O(n) | O(h) where h is height |
| In-order Traversal | Binary Tree | O(n) | O(h) where h is height |

### Approach 1: Recursive with Range Validation

The key insight is that for each node, all values in its left subtree must be less than the node's value, and all values in its right subtree must be greater than the node's value. We can enforce this by maintaining valid ranges for each subtree.

**Key Insight:** Instead of just comparing a node with its immediate children, we need to validate that all descendants satisfy the BST property relative to their ancestors.

### Algorithm Steps

1. Define a helper function that takes a node and valid range `[min, max]`
2. For each node:
   - Check if the node's value is within the valid range `(min, max)`
   - Recursively validate left subtree with updated max value (node's value)
   - Recursively validate right subtree with updated min value (node's value)
3. Base case: `null` nodes are valid
4. Return `true` only if all nodes satisfy their range constraints

### Example Walkthrough

For the invalid tree `[5,1,4,null,null,3,6]`:

1. **Node 5**: Range `(-∞, +∞)` - Valid
2. **Node 1** (left of 5): Range `(-∞, 5)` - Valid (1 < 5)
3. **Node 4** (right of 5): Range `(5, +∞)` - **Invalid!** (4 < 5)

The algorithm catches that node 4 violates its range constraint.

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node exactly once
- **Space Complexity O(h)**: Recursion stack depth equals tree height (O(log n) for balanced tree, O(n) for skewed tree)
- This is optimal because we must examine every node to validate the entire tree

### Approach 2: In-order Traversal

An alternative approach uses the property that in-order traversal of a BST produces values in strictly increasing order.

**Algorithm:**
1. Perform in-order traversal (left → root → right)
2. Track the previous node's value
3. Ensure each visited value is greater than the previous
4. If any value violates this ordering, return `false`

This approach has the same time and space complexity as the range validation method.
