from typing import Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_subtree(root: Optional[TreeNode], sub_root: Optional[TreeNode]) -> bool:
    """
    Checks if subRoot is a subtree of root.
    Uses recursive DFS for O(m * n) time complexity.

    Args:
        root: Root node of the main tree
        sub_root: Root node of the subtree to find

    Returns:
        True if subRoot is a subtree of root, False otherwise
    """
    if root is None:
        return False

    # Check if current tree matches subRoot
    if is_same_tree(root, sub_root):
        return True

    # Recursively check left and right subtrees
    return is_subtree(root.left, sub_root) or is_subtree(root.right, sub_root)


def is_same_tree(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    """Helper function to check if two trees are identical."""
    if p is None and q is None:
        return True
    if p is None or q is None:
        return False
    return (p.val == q.val and
            is_same_tree(p.left, q.left) and
            is_same_tree(p.right, q.right))
