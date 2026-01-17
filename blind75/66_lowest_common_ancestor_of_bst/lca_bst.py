from typing import Optional


class TreeNode:
    """Represents a node in a binary tree."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def lowest_common_ancestor(
    root: Optional[TreeNode], p: TreeNode, q: TreeNode
) -> Optional[TreeNode]:
    """
    Finds the LCA of two nodes in a BST.
    Uses iterative approach with O(H) time and O(1) space complexity.

    Args:
        root: Root node of the binary search tree
        p: First node
        q: Second node

    Returns:
        The lowest common ancestor of p and q
    """
    current = root

    while current:
        # Both nodes are in left subtree
        if p.val < current.val and q.val < current.val:
            current = current.left
        # Both nodes are in right subtree
        elif p.val > current.val and q.val > current.val:
            current = current.right
        else:
            # Split point found - one node is on left, other on right
            # (or current node equals one of them)
            return current

    return None


def lowest_common_ancestor_recursive(
    root: Optional[TreeNode], p: TreeNode, q: TreeNode
) -> Optional[TreeNode]:
    """
    Finds the LCA using recursive approach.

    Args:
        root: Root node of the binary search tree
        p: First node
        q: Second node

    Returns:
        The lowest common ancestor of p and q
    """
    if root is None:
        return None

    # Both nodes are in left subtree
    if p.val < root.val and q.val < root.val:
        return lowest_common_ancestor_recursive(root.left, p, q)

    # Both nodes are in right subtree
    if p.val > root.val and q.val > root.val:
        return lowest_common_ancestor_recursive(root.right, p, q)

    # Split point found
    return root
