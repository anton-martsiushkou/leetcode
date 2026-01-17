import pytest
from subtree_of_another_tree import TreeNode, is_subtree


def test_example_1_subtree_exists():
    """Test case from example 1."""
    root = TreeNode(3)
    root.left = TreeNode(4)
    root.left.left = TreeNode(1)
    root.left.right = TreeNode(2)
    root.right = TreeNode(5)

    sub_root = TreeNode(4)
    sub_root.left = TreeNode(1)
    sub_root.right = TreeNode(2)

    result = is_subtree(root, sub_root)
    assert result is True


def test_example_2_subtree_has_extra_node():
    """Test case from example 2."""
    root = TreeNode(3)
    root.left = TreeNode(4)
    root.left.left = TreeNode(1)
    root.left.left.left = TreeNode(0)
    root.left.right = TreeNode(2)
    root.right = TreeNode(5)

    sub_root = TreeNode(4)
    sub_root.left = TreeNode(1)
    sub_root.right = TreeNode(2)

    result = is_subtree(root, sub_root)
    assert result is False


def test_example_3_simple_match():
    """Test case from example 3."""
    root = TreeNode(1)
    root.left = TreeNode(1)

    sub_root = TreeNode(1)

    result = is_subtree(root, sub_root)
    assert result is True


def test_identical_trees():
    """Test with identical trees."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    sub_root = TreeNode(1)
    sub_root.left = TreeNode(2)
    sub_root.right = TreeNode(3)

    result = is_subtree(root, sub_root)
    assert result is True


def test_subroot_is_leaf():
    """Test where subRoot is a leaf of root."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(4)
    root.right = TreeNode(3)

    sub_root = TreeNode(4)

    result = is_subtree(root, sub_root)
    assert result is True


def test_no_match_different_values():
    """Test with no match - different values."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    sub_root = TreeNode(1)
    sub_root.left = TreeNode(2)
    sub_root.right = TreeNode(4)

    result = is_subtree(root, sub_root)
    assert result is False


def test_no_match_different_structure():
    """Test with no match - different structure."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)

    sub_root = TreeNode(1)
    sub_root.left = TreeNode(2)

    result = is_subtree(root, sub_root)
    assert result is False


def test_empty_root():
    """Test with empty root."""
    sub_root = TreeNode(1)
    result = is_subtree(None, sub_root)
    assert result is False


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
