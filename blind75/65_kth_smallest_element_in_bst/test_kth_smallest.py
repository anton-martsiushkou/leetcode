import pytest
from kth_smallest import TreeNode, kth_smallest, kth_smallest_recursive


def test_example_1():
    """Test case from example 1."""
    root = TreeNode(3, TreeNode(1, None, TreeNode(2)), TreeNode(4))
    assert kth_smallest(root, 1) == 1


def test_example_2():
    """Test case from example 2."""
    root = TreeNode(
        5,
        TreeNode(3, TreeNode(2, TreeNode(1)), TreeNode(4)),
        TreeNode(6)
    )
    assert kth_smallest(root, 3) == 3


def test_single_node():
    """Test with single node."""
    root = TreeNode(1)
    assert kth_smallest(root, 1) == 1


def test_k_equals_tree_size():
    """Test when k equals the total number of nodes."""
    root = TreeNode(2, TreeNode(1), TreeNode(3))
    assert kth_smallest(root, 3) == 3


def test_left_skewed_tree():
    """Test with left-skewed tree."""
    root = TreeNode(3, TreeNode(2, TreeNode(1)))
    assert kth_smallest(root, 2) == 2


def test_right_skewed_tree():
    """Test with right-skewed tree."""
    root = TreeNode(1, None, TreeNode(2, None, TreeNode(3)))
    assert kth_smallest(root, 2) == 2


def test_larger_tree():
    """Test with larger balanced tree."""
    root = TreeNode(
        4,
        TreeNode(2, TreeNode(1), TreeNode(3)),
        TreeNode(6, TreeNode(5), TreeNode(7))
    )
    assert kth_smallest(root, 4) == 4
    assert kth_smallest(root, 1) == 1
    assert kth_smallest(root, 7) == 7


def test_recursive_example_1():
    """Test recursive version with example 1."""
    root = TreeNode(3, TreeNode(1, None, TreeNode(2)), TreeNode(4))
    assert kth_smallest_recursive(root, 1) == 1


def test_recursive_example_2():
    """Test recursive version with example 2."""
    root = TreeNode(
        5,
        TreeNode(3, TreeNode(2, TreeNode(1)), TreeNode(4)),
        TreeNode(6)
    )
    assert kth_smallest_recursive(root, 3) == 3


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
