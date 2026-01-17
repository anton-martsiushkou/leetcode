# Kth Smallest Element in a BST

## Problem Description

Given the root of a binary search tree, and an integer `k`, return the `kth` smallest value (1-indexed) of all the values of the nodes in the tree.

## Examples

**Example 1:**
```
    3
   / \
  1   4
   \
    2

Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Example 2:**
```
        5
       / \
      3   6
     / \
    2   4
   /
  1

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

## Constraints

- The number of nodes in the tree is `n`
- `1 <= k <= n <= 10^4`
- `0 <= Node.val <= 10^4`

## Follow-up

If the BST is modified often (i.e., we can do insert and delete operations) and you need to find the kth smallest frequently, how would you optimize?

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| In-order Traversal (Iterative) | Stack | O(H + k) | O(H) where H is height |
| In-order Traversal (Recursive) | Binary Search Tree | O(n) | O(H) where H is height |

### Approach: In-order Traversal

The key insight is that in-order traversal of a BST visits nodes in ascending sorted order. Therefore, the kth node visited during in-order traversal is the kth smallest element.

**Key Insight:** In-order traversal (left → root → right) of a BST produces values in strictly increasing order, so we can stop once we've visited k nodes.

### Algorithm Steps (Iterative)

1. Initialize an empty stack and set `current` to root
2. While stack is not empty or current is not null:
   - Traverse to the leftmost node, pushing all nodes onto stack
   - Pop a node from stack (this is the next smallest element)
   - Decrement k
   - If k == 0, return current node's value
   - Move to the right child
3. The kth popped node contains the kth smallest value

### Example Walkthrough

For the tree `[3,1,4,null,2]` with k = 1:

```
    3
   / \
  1   4
   \
    2
```

**In-order traversal visits:** 1 → 2 → 3 → 4

1. Push 3, then push 1 (leftmost)
2. Pop 1 (1st smallest), k = 0
3. Return 1

For k = 3:
1. Visit 1 (k=2), then 2 (k=1), then 3 (k=0)
2. Return 3

### Why This is Optimal

- **Time Complexity O(H + k)**: We traverse down to the leftmost node (H steps) and then visit k nodes. In the worst case when k = n, this is O(n)
- **Space Complexity O(H)**: Stack depth equals tree height (O(log n) for balanced tree, O(n) for skewed tree)
- This is optimal because we must visit at least k nodes to find the kth smallest

### Follow-up Optimization

For frequent kth smallest queries with modifications:
1. **Augment BST nodes** with a `size` field tracking the number of nodes in each subtree
2. For kth smallest query:
   - If `leftSubtreeSize == k-1`, return current node
   - If `leftSubtreeSize >= k`, search left subtree
   - Otherwise, search right subtree for `(k - leftSubtreeSize - 1)`th element
3. This reduces query time to O(H) even with modifications
4. Insert/delete operations update size fields: O(H) additional work
