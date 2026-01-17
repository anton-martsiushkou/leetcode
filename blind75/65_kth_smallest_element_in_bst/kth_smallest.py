from typing import Optional


class TreeNode:
    """Represents a node in a binary tree."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def kth_smallest(root: Optional[TreeNode], k: int) -> int:
    """
    Returns the kth smallest value in a BST.
    Uses iterative in-order traversal for O(H + k) time complexity.

    Args:
        root: Root node of the binary search tree
        k: The kth position (1-indexed)

    Returns:
        The kth smallest value in the BST
    """
    stack = []
    current = root

    while stack or current:
        # Traverse to the leftmost node
        while current:
            stack.append(current)
            current = current.left

        # Pop the next smallest node
        current = stack.pop()

        k -= 1
        if k == 0:
            return current.val

        # Move to right subtree
        current = current.right

    return -1  # Should never reach here if k is valid


def kth_smallest_recursive(root: Optional[TreeNode], k: int) -> int:
    """
    Returns the kth smallest value using recursive in-order traversal.

    Args:
        root: Root node of the binary search tree
        k: The kth position (1-indexed)

    Returns:
        The kth smallest value in the BST
    """
    result = []

    def inorder(node: Optional[TreeNode]) -> None:
        """Performs in-order traversal and collects values."""
        if node is None or len(result) >= k:
            return

        # Traverse left subtree
        inorder(node.left)

        # Visit current node
        result.append(node.val)
        if len(result) >= k:
            return

        # Traverse right subtree
        inorder(node.right)

    inorder(root)
    return result[k - 1] if len(result) >= k else -1
