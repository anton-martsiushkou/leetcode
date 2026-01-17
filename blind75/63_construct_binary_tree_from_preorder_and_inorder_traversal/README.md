# Construct Binary Tree from Preorder and Inorder Traversal

## Problem Description

Given two integer arrays `preorder` and `inorder` where:
- `preorder` is the preorder traversal of a binary tree
- `inorder` is the inorder traversal of the same tree

Construct and return the binary tree.

## Examples

**Example 1:**
```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
Explanation:
      3
     / \
    9  20
      /  \
     15   7
```

**Example 2:**
```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

## Constraints

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` and `inorder` consist of unique values
- Each value of `inorder` also appears in `preorder`
- `preorder` is guaranteed to be the preorder traversal of the tree
- `inorder` is guaranteed to be the inorder traversal of the tree

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Recursive | Binary Tree + HashMap | O(n) | O(n) |

### Approach: DFS Recursive with HashMap

The key insight is understanding how preorder and inorder traversals work:
- **Preorder**: Root → Left → Right (root is always first)
- **Inorder**: Left → Root → Right (root splits left and right subtrees)

By combining these properties:
1. The first element in preorder is always the root
2. Find this root in inorder to split left and right subtrees
3. Recursively build left and right subtrees

**Key Insight:** Preorder tells us the root, inorder tells us which nodes are in the left vs. right subtree.

### Algorithm Steps

1. Create a hashmap of inorder values to indices for O(1) lookup
2. Define recursive helper function with bounds:
   - Take the next value from preorder as the root
   - Find root's index in inorder using the hashmap
   - All values to the left of this index in inorder are in the left subtree
   - All values to the right are in the right subtree
   - Recursively build left subtree with left portion of inorder
   - Recursively build right subtree with right portion of inorder
3. Return the root

### Example Walkthrough

For preorder = `[3,9,20,15,7]` and inorder = `[9,3,15,20,7]`:

```
HashMap: {9: 0, 3: 1, 15: 2, 20: 3, 7: 4}

build(inorder[0:5]):
  ├─ preorder[0] = 3 (root)
  ├─ find 3 in inorder at index 1
  ├─ left subtree: inorder[0:1] = [9]
  │    build(inorder[0:1]):
  │      ├─ preorder[1] = 9 (root)
  │      ├─ find 9 in inorder at index 0
  │      ├─ left: inorder[0:0] = [] → null
  │      └─ right: inorder[1:1] = [] → null
  │      return node(9)
  └─ right subtree: inorder[2:5] = [15,20,7]
       build(inorder[2:5]):
         ├─ preorder[2] = 20 (root)
         ├─ find 20 in inorder at index 3
         ├─ left: inorder[2:3] = [15]
         │    build(inorder[2:3]):
         │      ├─ preorder[3] = 15
         │      └─ return node(15)
         └─ right: inorder[4:5] = [7]
              build(inorder[4:5]):
                ├─ preorder[4] = 7
                └─ return node(7)
         return node(20)
  return node(3)

Result:
      3
     / \
    9  20
      /  \
     15   7
```

### Why This is Optimal

- **Time Complexity O(n)**: We visit each node once, and the hashmap provides O(1) lookups
- **Space Complexity O(n)**:
  - O(n) for the hashmap
  - O(h) for the recursion stack (where h is tree height)
  - Total: O(n)
- This is optimal because we must examine every element at least once to construct the tree

### Key Points

1. **Unique values**: The problem guarantees all values are unique, which ensures:
   - Each value appears exactly once in both arrays
   - We can uniquely identify the root position in inorder

2. **Preorder index**: We use a global index or iterator to track position in preorder array

3. **Inorder bounds**: We pass start/end indices to define the current subtree's portion of inorder

4. **Base cases**: When start > end in inorder bounds, return null (no nodes in this subtree)

### Why Both Traversals Are Needed

- **Preorder alone**: Not enough to determine structure (multiple trees can have same preorder)
- **Inorder alone**: Not enough to determine structure (multiple trees can have same inorder)
- **Together**: Uniquely determines the tree structure
