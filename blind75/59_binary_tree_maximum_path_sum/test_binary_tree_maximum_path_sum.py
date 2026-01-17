import pytest
from binary_tree_maximum_path_sum import TreeNode, max_path_sum


def test_example_1_simple_tree():
    """Test case from example 1."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    result = max_path_sum(root)
    assert result == 6


def test_example_2_tree_with_negative_root():
    """Test case from example 2."""
    root = TreeNode(-10)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)

    result = max_path_sum(root)
    assert result == 42


def test_example_3_single_negative_node():
    """Test case from example 3."""
    root = TreeNode(-3)
    result = max_path_sum(root)
    assert result == -3


def test_single_positive_node():
    """Test with single positive node."""
    root = TreeNode(5)
    result = max_path_sum(root)
    assert result == 5


def test_all_negative_values():
    """Test with all negative values."""
    root = TreeNode(-2)
    root.left = TreeNode(-1)
    root.right = TreeNode(-3)

    result = max_path_sum(root)
    assert result == -1


def test_left_skewed_tree():
    """Test with left skewed tree."""
    root = TreeNode(5)
    root.left = TreeNode(4)
    root.left.left = TreeNode(3)

    result = max_path_sum(root)
    assert result == 12


def test_tree_with_negative_branch():
    """Test tree with negative branch."""
    root = TreeNode(2)
    root.left = TreeNode(-1)
    root.left.left = TreeNode(3)
    root.left.right = TreeNode(4)
    root.right = TreeNode(-2)

    result = max_path_sum(root)
    assert result == 6


def test_complex_tree():
    """Test with complex tree."""
    root = TreeNode(5)
    root.left = TreeNode(4)
    root.left.left = TreeNode(11)
    root.left.right = TreeNode(2)
    root.right = TreeNode(8)
    root.right.left = TreeNode(13)
    root.right.right = TreeNode(4)

    result = max_path_sum(root)
    assert result == 48


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
