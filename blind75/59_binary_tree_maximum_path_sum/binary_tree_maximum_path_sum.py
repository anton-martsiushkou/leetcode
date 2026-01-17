from typing import Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def max_path_sum(root: Optional[TreeNode]) -> int:
    """
    Returns the maximum path sum in a binary tree.
    Uses recursive DFS with global maximum tracking for O(n) time complexity.

    Args:
        root: Root node of the binary tree

    Returns:
        Maximum path sum of any non-empty path in the tree
    """
    max_sum = float('-inf')

    def max_gain(node: Optional[TreeNode]) -> int:
        nonlocal max_sum

        if node is None:
            return 0

        # Get maximum gain from left and right subtrees (ignore negative gains)
        left_gain = max(max_gain(node.left), 0)
        right_gain = max(max_gain(node.right), 0)

        # Calculate path sum through current node
        path_sum = node.val + left_gain + right_gain

        # Update global maximum
        max_sum = max(max_sum, path_sum)

        # Return maximum single path (can only use one child)
        return node.val + max(left_gain, right_gain)

    max_gain(root)
    return int(max_sum)
