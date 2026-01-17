from typing import Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def invert_tree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    """
    Inverts a binary tree (mirrors it).
    Uses recursive DFS for O(n) time complexity.

    Args:
        root: Root node of the binary tree

    Returns:
        Root node of the inverted tree
    """
    if root is None:
        return None

    # Swap left and right children
    root.left, root.right = root.right, root.left

    # Recursively invert subtrees
    invert_tree(root.left)
    invert_tree(root.right)

    return root
