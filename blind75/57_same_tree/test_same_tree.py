import pytest
from same_tree import TreeNode, is_same_tree


def test_example_1_identical_trees():
    """Test case from example 1."""
    p = TreeNode(1)
    p.left = TreeNode(2)
    p.right = TreeNode(3)

    q = TreeNode(1)
    q.left = TreeNode(2)
    q.right = TreeNode(3)

    result = is_same_tree(p, q)
    assert result is True


def test_example_2_different_structure():
    """Test case from example 2."""
    p = TreeNode(1)
    p.left = TreeNode(2)

    q = TreeNode(1)
    q.right = TreeNode(2)

    result = is_same_tree(p, q)
    assert result is False


def test_example_3_different_values():
    """Test case from example 3."""
    p = TreeNode(1)
    p.left = TreeNode(2)
    p.right = TreeNode(1)

    q = TreeNode(1)
    q.left = TreeNode(1)
    q.right = TreeNode(2)

    result = is_same_tree(p, q)
    assert result is False


def test_both_empty():
    """Test with both trees empty."""
    result = is_same_tree(None, None)
    assert result is True


def test_one_empty():
    """Test with one tree empty."""
    p = TreeNode(1)
    result = is_same_tree(p, None)
    assert result is False


def test_single_node_same():
    """Test with single node, same value."""
    p = TreeNode(5)
    q = TreeNode(5)
    result = is_same_tree(p, q)
    assert result is True


def test_single_node_different():
    """Test with single node, different values."""
    p = TreeNode(5)
    q = TreeNode(10)
    result = is_same_tree(p, q)
    assert result is False


def test_deep_identical_trees():
    """Test with deeper identical trees."""
    p = TreeNode(1)
    p.left = TreeNode(2)
    p.left.left = TreeNode(4)
    p.left.right = TreeNode(5)
    p.right = TreeNode(3)
    p.right.right = TreeNode(6)

    q = TreeNode(1)
    q.left = TreeNode(2)
    q.left.left = TreeNode(4)
    q.left.right = TreeNode(5)
    q.right = TreeNode(3)
    q.right.right = TreeNode(6)

    result = is_same_tree(p, q)
    assert result is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
