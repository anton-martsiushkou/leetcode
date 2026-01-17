import pytest
from valid_tree import valid_tree, valid_tree_union_find


def test_example_1():
    """Test case from example 1 - valid tree."""
    n = 5
    edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
    assert valid_tree(n, edges) is True


def test_example_2():
    """Test case from example 2 - has cycle."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
    assert valid_tree(n, edges) is False


def test_single_node():
    """Test with single node."""
    n = 1
    edges = []
    assert valid_tree(n, edges) is True


def test_two_nodes_connected():
    """Test with two nodes connected."""
    n = 2
    edges = [[0, 1]]
    assert valid_tree(n, edges) is True


def test_disconnected_graph():
    """Test with disconnected graph."""
    n = 4
    edges = [[0, 1], [2, 3]]
    assert valid_tree(n, edges) is False


def test_too_many_edges():
    """Test with too many edges."""
    n = 3
    edges = [[0, 1], [0, 2], [1, 2]]
    assert valid_tree(n, edges) is False


def test_too_few_edges():
    """Test with too few edges."""
    n = 4
    edges = [[0, 1], [1, 2]]
    assert valid_tree(n, edges) is False


def test_linear_chain():
    """Test with linear chain."""
    n = 4
    edges = [[0, 1], [1, 2], [2, 3]]
    assert valid_tree(n, edges) is True


def test_union_find_example_1():
    """Test Union Find with example 1."""
    n = 5
    edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
    assert valid_tree_union_find(n, edges) is True


def test_union_find_example_2():
    """Test Union Find with example 2."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
    assert valid_tree_union_find(n, edges) is False


def test_union_find_single_node():
    """Test Union Find with single node."""
    n = 1
    edges = []
    assert valid_tree_union_find(n, edges) is True


def test_union_find_disconnected():
    """Test Union Find with disconnected graph."""
    n = 4
    edges = [[0, 1], [2, 3]]
    assert valid_tree_union_find(n, edges) is False


def test_union_find_linear():
    """Test Union Find with linear chain."""
    n = 4
    edges = [[0, 1], [1, 2], [2, 3]]
    assert valid_tree_union_find(n, edges) is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
