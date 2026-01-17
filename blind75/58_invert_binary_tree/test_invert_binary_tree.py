import pytest
from invert_binary_tree import TreeNode, invert_tree


def is_same_tree(p, q):
    """Helper function to check if two trees are identical."""
    if p is None and q is None:
        return True
    if p is None or q is None:
        return False
    return (p.val == q.val and
            is_same_tree(p.left, q.left) and
            is_same_tree(p.right, q.right))


def test_example_1_complete_tree():
    """Test case from example 1."""
    root = TreeNode(4)
    root.left = TreeNode(2)
    root.left.left = TreeNode(1)
    root.left.right = TreeNode(3)
    root.right = TreeNode(7)
    root.right.left = TreeNode(6)
    root.right.right = TreeNode(9)

    expected = TreeNode(4)
    expected.left = TreeNode(7)
    expected.left.left = TreeNode(9)
    expected.left.right = TreeNode(6)
    expected.right = TreeNode(2)
    expected.right.left = TreeNode(3)
    expected.right.right = TreeNode(1)

    result = invert_tree(root)
    assert is_same_tree(result, expected)


def test_example_2_simple_tree():
    """Test case from example 2."""
    root = TreeNode(2)
    root.left = TreeNode(1)
    root.right = TreeNode(3)

    expected = TreeNode(2)
    expected.left = TreeNode(3)
    expected.right = TreeNode(1)

    result = invert_tree(root)
    assert is_same_tree(result, expected)


def test_example_3_empty_tree():
    """Test case from example 3."""
    result = invert_tree(None)
    assert result is None


def test_single_node():
    """Test with single node."""
    root = TreeNode(1)
    expected = TreeNode(1)

    result = invert_tree(root)
    assert is_same_tree(result, expected)


def test_left_skewed_tree():
    """Test with left skewed tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(3)

    expected = TreeNode(1)
    expected.right = TreeNode(2)
    expected.right.right = TreeNode(3)

    result = invert_tree(root)
    assert is_same_tree(result, expected)


def test_right_skewed_tree():
    """Test with right skewed tree."""
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.right = TreeNode(3)

    expected = TreeNode(1)
    expected.left = TreeNode(2)
    expected.left.left = TreeNode(3)

    result = invert_tree(root)
    assert is_same_tree(result, expected)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
