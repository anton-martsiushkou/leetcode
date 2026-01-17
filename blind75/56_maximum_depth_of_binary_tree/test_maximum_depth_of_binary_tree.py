import pytest
from maximum_depth_of_binary_tree import TreeNode, max_depth


def test_example_1_balanced_tree():
    """Test case from example 1."""
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)

    result = max_depth(root)
    assert result == 3


def test_example_2_right_skewed():
    """Test case from example 2."""
    root = TreeNode(1)
    root.right = TreeNode(2)

    result = max_depth(root)
    assert result == 2


def test_example_3_empty_tree():
    """Test case from example 3."""
    result = max_depth(None)
    assert result == 0


def test_example_4_single_node():
    """Test case from example 4."""
    root = TreeNode(0)
    result = max_depth(root)
    assert result == 1


def test_left_skewed_tree():
    """Test with left skewed tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(3)
    root.left.left.left = TreeNode(4)

    result = max_depth(root)
    assert result == 4


def test_complete_binary_tree():
    """Test with complete binary tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.left.left = TreeNode(4)
    root.left.right = TreeNode(5)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(7)

    result = max_depth(root)
    assert result == 3


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
