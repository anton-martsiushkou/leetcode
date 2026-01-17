# Serialize and Deserialize Binary Tree

## Problem Description

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

## Examples

**Example 1:**
```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]
Explanation:
      1
     / \
    2   3
       / \
      4   5
Serialize to: "1,2,null,null,3,4,null,null,5,null,null"
Deserialize back to the original tree
```

**Example 2:**
```
Input: root = []
Output: []
```

**Example 3:**
```
Input: root = [1]
Output: [1]
```

## Constraints

- The number of nodes in the tree is in the range `[0, 10^4]`
- `-1000 <= Node.val <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| DFS Preorder | Binary Tree + String | O(n) | O(n) |
| BFS Level Order | Binary Tree + Queue | O(n) | O(n) |

### Approach: DFS Preorder Traversal (Recommended)

The preorder traversal (root → left → right) approach is ideal for serialization because:
1. The root is always first, making it easy to reconstruct
2. We can recursively deserialize left subtree, then right subtree
3. Null markers preserve the tree structure

**Key Insight:** Preorder traversal with null markers contains enough information to uniquely reconstruct the tree.

### Algorithm Steps

**Serialize:**
1. Base case: If node is `null`, append "null" marker
2. Append current node's value
3. Recursively serialize left subtree
4. Recursively serialize right subtree
5. Join values with a delimiter (e.g., comma)

**Deserialize:**
1. Split the string by delimiter to get values
2. Use a helper function with an iterator/index:
   - If current value is "null", return `null`
   - Create node with current value
   - Recursively build left subtree
   - Recursively build right subtree
   - Return the node
3. Return the root

### Example Walkthrough

For the tree `[1,2,3,null,null,4,5]`:

**Serialize:**
```
serialize(1):
  ├─ add "1"
  ├─ serialize(2):
  │    ├─ add "2"
  │    ├─ serialize(null) → add "null"
  │    └─ serialize(null) → add "null"
  └─ serialize(3):
       ├─ add "3"
       ├─ serialize(4):
       │    ├─ add "4"
       │    ├─ serialize(null) → add "null"
       │    └─ serialize(null) → add "null"
       └─ serialize(5):
            ├─ add "5"
            ├─ serialize(null) → add "null"
            └─ serialize(null) → add "null"

Result: "1,2,null,null,3,4,null,null,5,null,null"
```

**Deserialize:**
```
Values: [1, 2, null, null, 3, 4, null, null, 5, null, null]
Index: 0

deserialize():
  ├─ val = 1, create node(1)
  ├─ left = deserialize():
  │    ├─ val = 2, create node(2)
  │    ├─ left = deserialize(): val = null → null
  │    └─ right = deserialize(): val = null → null
  │    return node(2)
  └─ right = deserialize():
       ├─ val = 3, create node(3)
       ├─ left = deserialize():
       │    ├─ val = 4, create node(4)
       │    ├─ left = deserialize(): val = null → null
       │    └─ right = deserialize(): val = null → null
       │    return node(4)
       └─ right = deserialize():
            ├─ val = 5, create node(5)
            ├─ left = deserialize(): val = null → null
            └─ right = deserialize(): val = null → null
            return node(5)
       return node(3)
  return node(1)
```

### Why This is Optimal

- **Time Complexity O(n)**: Both serialize and deserialize visit each node once
- **Space Complexity O(n)**: We store n node values plus null markers
- Preorder traversal is optimal because it naturally matches the recursive structure and reconstruction order

### Alternative Approaches

1. **BFS Level Order**: Serialize level by level
   - More intuitive for some
   - Same complexity
   - Requires more null markers for incomplete levels

2. **Inorder + Preorder**: Use two traversals
   - Redundant and more complex
   - Requires O(2n) space
