import pytest
from binary_tree_level_order_traversal import TreeNode, level_order


def test_example_1_balanced_tree():
    """Test case from example 1."""
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)

    result = level_order(root)
    assert result == [[3], [9, 20], [15, 7]]


def test_example_2_single_node():
    """Test case from example 2."""
    root = TreeNode(1)
    result = level_order(root)
    assert result == [[1]]


def test_example_3_empty_tree():
    """Test case from example 3."""
    result = level_order(None)
    assert result == []


def test_left_skewed_tree():
    """Test with left skewed tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(3)

    result = level_order(root)
    assert result == [[1], [2], [3]]


def test_right_skewed_tree():
    """Test with right skewed tree."""
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.right = TreeNode(3)

    result = level_order(root)
    assert result == [[1], [2], [3]]


def test_complete_binary_tree():
    """Test with complete binary tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.left = TreeNode(4)
    root.left.right = TreeNode(5)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(7)

    result = level_order(root)
    assert result == [[1], [2, 3], [4, 5, 6, 7]]


def test_unbalanced_tree():
    """Test with unbalanced tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.left = TreeNode(4)

    result = level_order(root)
    assert result == [[1], [2, 3], [4]]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
