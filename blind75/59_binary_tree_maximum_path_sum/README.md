# Binary Tree Maximum Path Sum

## Problem Description

A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the `root` of a binary tree, return the maximum path sum of any non-empty path.

## Examples

**Example 1:**
```
Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
    1
   / \
  2   3
```

**Example 2:**
```
Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
     -10
     / \
    9  20
      /  \
     15   7
```

**Example 3:**
```
Input: root = [-3]
Output: -3
```

## Constraints

- The number of nodes in the tree is in the range `[1, 3 * 10^4]`
- `-1000 <= Node.val <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree | O(n) | O(h) |

### Approach: DFS Recursive with Global Maximum

The key insight is that for each node, we need to consider two different scenarios:
1. **Maximum path sum through this node** (including both left and right paths)
2. **Maximum path sum from this node downward** (only one path, to return to parent)

We use a helper function that returns the maximum path sum from a node downward (single path), while updating a global maximum that tracks the best path sum we've seen (which might go through that node).

**Key Insight:** At each node, the maximum path that goes through it is `node.val + max_left + max_right`, but we can only return `node.val + max(max_left, max_right)` to the parent (single path).

### Algorithm Steps

1. Initialize a global variable to track the maximum path sum (start with negative infinity)
2. Define a recursive helper function `max_gain(node)`:
   - Base case: If node is `null`, return 0
   - Recursively get maximum gain from left subtree (ignore negative gains)
   - Recursively get maximum gain from right subtree (ignore negative gains)
   - Calculate path sum through current node: `node.val + left_gain + right_gain`
   - Update global maximum with this path sum
   - Return maximum single path: `node.val + max(left_gain, right_gain)`
3. Call the helper function on root
4. Return the global maximum

### Example Walkthrough

For the tree `[-10,9,20,null,null,15,7]`:

```
max_gain(-10):
  ├─ max_gain(9) = 9 (leaf, max(0) = 9)
  │    update global_max = 9
  ├─ max_gain(20):
  │    ├─ max_gain(15) = 15
  │    │    update global_max = 15
  │    ├─ max_gain(7) = 7
  │    │    update global_max = 15 (unchanged)
  │    path through 20: 20 + 15 + 7 = 42
  │    update global_max = 42
  │    return 20 + max(15, 7) = 35
  path through -10: -10 + 9 + 35 = 34
  update global_max = 42 (unchanged)
  return max(-10 + 35, 0) = 25

Final answer: 42
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node exactly once
- **Space Complexity O(h)**: The recursion stack depth equals the height of the tree
  - Best case (balanced tree): O(log n)
  - Worst case (skewed tree): O(n)
- This is optimal because we must examine every node to find the maximum path

### Key Points

1. **Negative values**: We use `max(0, gain)` to ignore negative contributions from subtrees
2. **Global maximum**: We track this separately because the path through a node (using both children) is different from the path we return to the parent (using only one child)
3. **Single vs. Split paths**: At each node, we consider the "split" path (through the node) for the global max, but only return a "single" path upward
