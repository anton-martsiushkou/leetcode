from typing import List, Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def build_tree(preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
    """
    Constructs a binary tree from preorder and inorder traversals.
    Uses recursive DFS with hashmap for O(n) time complexity.

    Args:
        preorder: Preorder traversal of the tree
        inorder: Inorder traversal of the tree

    Returns:
        Root node of the constructed binary tree
    """
    # Create hashmap for O(1) lookup of inorder indices
    inorder_map = {val: i for i, val in enumerate(inorder)}
    preorder_index = 0

    def build(left: int, right: int) -> Optional[TreeNode]:
        nonlocal preorder_index

        # Base case: no elements to construct the tree
        if left > right:
            return None

        # Select the preorder index element as root and increment
        root_val = preorder[preorder_index]
        preorder_index += 1

        root = TreeNode(root_val)

        # Build left and right subtree
        # excluding inorder_map[root_val] element because it's the root
        root.left = build(left, inorder_map[root_val] - 1)
        root.right = build(inorder_map[root_val] + 1, right)

        return root

    return build(0, len(inorder) - 1)
