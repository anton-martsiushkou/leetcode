from typing import Optional


class TreeNode:
    """Represents a node in a binary tree."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_valid_bst(root: Optional[TreeNode]) -> bool:
    """
    Determines if a binary tree is a valid BST.
    Uses recursive range validation for O(n) time complexity.

    Args:
        root: Root node of the binary tree

    Returns:
        True if the tree is a valid BST, False otherwise
    """

    def validate(node: Optional[TreeNode], min_val: float, max_val: float) -> bool:
        """
        Validates that the node's value is within the valid range
        and recursively validates left and right subtrees.
        """
        if node is None:
            return True

        # Check if current node's value is within valid range
        if node.val <= min_val or node.val >= max_val:
            return False

        # Validate left subtree (all values must be < node.val)
        # and right subtree (all values must be > node.val)
        return validate(node.left, min_val, node.val) and validate(
            node.right, node.val, max_val
        )

    return validate(root, float("-inf"), float("inf"))
