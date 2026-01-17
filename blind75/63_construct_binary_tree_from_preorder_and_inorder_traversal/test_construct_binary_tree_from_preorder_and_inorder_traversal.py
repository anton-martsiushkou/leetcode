import pytest
from construct_binary_tree_from_preorder_and_inorder_traversal import TreeNode, build_tree


def is_same_tree(p, q):
    """Helper function to check if two trees are identical."""
    if p is None and q is None:
        return True
    if p is None or q is None:
        return False
    return (p.val == q.val and
            is_same_tree(p.left, q.left) and
            is_same_tree(p.right, q.right))


def test_example_1_balanced_tree():
    """Test case from example 1."""
    preorder = [3, 9, 20, 15, 7]
    inorder = [9, 3, 15, 20, 7]

    expected = TreeNode(3)
    expected.left = TreeNode(9)
    expected.right = TreeNode(20)
    expected.right.left = TreeNode(15)
    expected.right.right = TreeNode(7)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_example_2_single_node():
    """Test case from example 2."""
    preorder = [-1]
    inorder = [-1]

    expected = TreeNode(-1)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_left_skewed_tree():
    """Test with left skewed tree."""
    preorder = [1, 2, 3]
    inorder = [3, 2, 1]

    expected = TreeNode(1)
    expected.left = TreeNode(2)
    expected.left.left = TreeNode(3)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_right_skewed_tree():
    """Test with right skewed tree."""
    preorder = [1, 2, 3]
    inorder = [1, 2, 3]

    expected = TreeNode(1)
    expected.right = TreeNode(2)
    expected.right.right = TreeNode(3)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_complete_binary_tree():
    """Test with complete binary tree."""
    preorder = [1, 2, 4, 5, 3, 6, 7]
    inorder = [4, 2, 5, 1, 6, 3, 7]

    expected = TreeNode(1)
    expected.left = TreeNode(2)
    expected.left.left = TreeNode(4)
    expected.left.right = TreeNode(5)
    expected.right = TreeNode(3)
    expected.right.left = TreeNode(6)
    expected.right.right = TreeNode(7)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_unbalanced_tree():
    """Test with unbalanced tree."""
    preorder = [1, 2, 4, 3]
    inorder = [4, 2, 1, 3]

    expected = TreeNode(1)
    expected.left = TreeNode(2)
    expected.left.left = TreeNode(4)
    expected.right = TreeNode(3)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


def test_two_nodes():
    """Test with two nodes."""
    preorder = [1, 2]
    inorder = [2, 1]

    expected = TreeNode(1)
    expected.left = TreeNode(2)

    result = build_tree(preorder, inorder)
    assert is_same_tree(result, expected)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
