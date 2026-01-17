import pytest
from validate_bst import TreeNode, is_valid_bst


def test_example_1_valid_bst():
    """Test case from example 1 - valid BST."""
    root = TreeNode(2, TreeNode(1), TreeNode(3))
    assert is_valid_bst(root) is True


def test_example_2_invalid_bst():
    """Test case from example 2 - invalid BST."""
    root = TreeNode(
        5, TreeNode(1), TreeNode(4, TreeNode(3), TreeNode(6))
    )
    assert is_valid_bst(root) is False


def test_empty_tree():
    """Test with empty tree."""
    assert is_valid_bst(None) is True


def test_single_node():
    """Test with single node."""
    root = TreeNode(1)
    assert is_valid_bst(root) is True


def test_invalid_left_child_greater():
    """Test invalid BST where left child is greater than root."""
    root = TreeNode(1, TreeNode(2))
    assert is_valid_bst(root) is False


def test_invalid_deep_violation():
    """Test invalid BST with deep violation of BST property."""
    # Tree: 10 with left child 5, which has right child 15 (15 > 10, invalid)
    root = TreeNode(
        10,
        TreeNode(5, None, TreeNode(15)),
        TreeNode(20)
    )
    assert is_valid_bst(root) is False


def test_valid_all_same_side():
    """Test valid BST with all nodes on same side."""
    root = TreeNode(1, None, TreeNode(2, None, TreeNode(3)))
    assert is_valid_bst(root) is True


def test_invalid_duplicate_values():
    """Test invalid BST with duplicate values."""
    root = TreeNode(2, TreeNode(2), TreeNode(3))
    assert is_valid_bst(root) is False


def test_valid_larger_tree():
    """Test valid BST with larger tree."""
    root = TreeNode(
        10,
        TreeNode(5, TreeNode(3), TreeNode(7)),
        TreeNode(15, TreeNode(12), TreeNode(20))
    )
    assert is_valid_bst(root) is True


def test_invalid_right_child_smaller():
    """Test invalid BST where right child is smaller than root."""
    root = TreeNode(10, None, TreeNode(5))
    assert is_valid_bst(root) is False


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
