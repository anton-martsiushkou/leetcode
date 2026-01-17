import pytest
from serialize_and_deserialize_binary_tree import TreeNode, Codec


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
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    root.right.left = TreeNode(4)
    root.right.right = TreeNode(5)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_example_2_empty_tree():
    """Test case from example 2."""
    codec = Codec()
    serialized = codec.serialize(None)
    deserialized = codec.deserialize(serialized)

    assert deserialized is None


def test_example_3_single_node():
    """Test case from example 3."""
    root = TreeNode(1)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_left_skewed_tree():
    """Test with left skewed tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(3)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_right_skewed_tree():
    """Test with right skewed tree."""
    root = TreeNode(1)
    root.right = TreeNode(2)
    root.right.right = TreeNode(3)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_complex_tree():
    """Test with complex tree."""
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(4)
    root.left.right = TreeNode(5)
    root.right = TreeNode(3)
    root.right.right = TreeNode(6)
    root.right.right.left = TreeNode(7)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_tree_with_negative_values():
    """Test with negative values."""
    root = TreeNode(-1)
    root.left = TreeNode(-2)
    root.right = TreeNode(-3)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_large_values():
    """Test with large values."""
    root = TreeNode(1000)
    root.left = TreeNode(-1000)
    root.right = TreeNode(500)

    codec = Codec()
    serialized = codec.serialize(root)
    deserialized = codec.deserialize(serialized)

    assert is_same_tree(root, deserialized)


def test_serialize_format():
    """Test serialization format."""
    codec = Codec()

    # Single node
    root = TreeNode(1)
    assert codec.serialize(root) == "1,null,null"

    # Empty tree
    assert codec.serialize(None) == "null"

    # Simple tree
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.right = TreeNode(3)
    assert codec.serialize(root) == "1,2,null,null,3,null,null"


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
