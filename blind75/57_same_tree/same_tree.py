from typing import Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_same_tree(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    """
    Checks if two binary trees are identical.
    Uses recursive DFS for O(n) time complexity.

    Args:
        p: Root node of first binary tree
        q: Root node of second binary tree

    Returns:
        True if trees are identical, False otherwise
    """
    # Both nodes are None - same structure
    if p is None and q is None:
        return True

    # One node is None, the other isn't - different structure
    if p is None or q is None:
        return False

    # Values differ - not the same
    if p.val != q.val:
        return False

    # Check if left and right subtrees are the same
    return is_same_tree(p.left, q.left) and is_same_tree(p.right, q.right)
