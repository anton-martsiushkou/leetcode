import pytest
from lca_bst import TreeNode, lowest_common_ancestor, lowest_common_ancestor_recursive


def build_example_tree():
    """Build the tree from examples 1 and 2.

            6
           / \
          2   8
         / \ / \
        0  4 7  9
          / \
         3   5
    """
    root = TreeNode(6)
    node2 = TreeNode(2)
    node8 = TreeNode(8)
    node0 = TreeNode(0)
    node4 = TreeNode(4)
    node7 = TreeNode(7)
    node9 = TreeNode(9)
    node3 = TreeNode(3)
    node5 = TreeNode(5)

    root.left = node2
    root.right = node8
    node2.left = node0
    node2.right = node4
    node8.left = node7
    node8.right = node9
    node4.left = node3
    node4.right = node5

    return root, node2, node8, node0, node4, node7, node9, node3, node5


def test_example_1_lca_is_root():
    """Test case from example 1 - LCA is root."""
    root, node2, node8, *_ = build_example_tree()
    result = lowest_common_ancestor(root, node2, node8)
    assert result.val == 6


def test_example_2_lca_is_one_of_nodes():
    """Test case from example 2 - LCA is one of the nodes."""
    root, node2, _, _, node4, *_ = build_example_tree()
    result = lowest_common_ancestor(root, node2, node4)
    assert result.val == 2


def test_both_nodes_in_left_subtree():
    """Test when both nodes are in left subtree."""
    root, _, _, node0, node4, *_ = build_example_tree()
    result = lowest_common_ancestor(root, node0, node4)
    assert result.val == 2


def test_both_nodes_in_right_subtree():
    """Test when both nodes are in right subtree."""
    root, _, _, _, _, node7, node9, *_ = build_example_tree()
    result = lowest_common_ancestor(root, node7, node9)
    assert result.val == 8


def test_deep_nodes():
    """Test with deep nodes in the tree."""
    root, _, _, _, _, _, _, node3, node5 = build_example_tree()
    result = lowest_common_ancestor(root, node3, node5)
    assert result.val == 4


def test_example_3_simple_tree():
    """Test case from example 3 - simple tree."""
    root = TreeNode(2, TreeNode(1))
    p = root
    q = root.left
    result = lowest_common_ancestor(root, p, q)
    assert result.val == 2


def test_recursive_example_1():
    """Test recursive version with example 1."""
    root, node2, node8, *_ = build_example_tree()
    result = lowest_common_ancestor_recursive(root, node2, node8)
    assert result.val == 6


def test_recursive_example_2():
    """Test recursive version with example 2."""
    root, node2, _, _, node4, *_ = build_example_tree()
    result = lowest_common_ancestor_recursive(root, node2, node4)
    assert result.val == 2


def test_recursive_both_in_left():
    """Test recursive version with both nodes in left subtree."""
    root, _, _, node0, node4, *_ = build_example_tree()
    result = lowest_common_ancestor_recursive(root, node0, node4)
    assert result.val == 2


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
