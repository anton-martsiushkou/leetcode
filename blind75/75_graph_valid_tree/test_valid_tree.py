import pytest
from valid_tree import valid_tree, valid_tree_union_find


def test_example_1_valid_tree():
    """Test case from example 1 - valid tree."""
    n = 5
    edges = [[0, 1], [0, 2], [0, 3], [1, 4]]
    assert valid_tree(n, edges) == True
    assert valid_tree_union_find(n, edges) == True


def test_example_2_contains_cycle():
    """Test case from example 2 - contains cycle."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
    assert valid_tree(n, edges) == False
    assert valid_tree_union_find(n, edges) == False


def test_single_node():
    """Test with single node."""
    n = 1
    edges = []
    assert valid_tree(n, edges) == True
    assert valid_tree_union_find(n, edges) == True


def test_two_nodes_connected():
    """Test with two nodes connected."""
    n = 2
    edges = [[0, 1]]
    assert valid_tree(n, edges) == True
    assert valid_tree_union_find(n, edges) == True


def test_disconnected_graph():
    """Test disconnected graph."""
    n = 4
    edges = [[0, 1], [2, 3]]
    assert valid_tree(n, edges) == False
    assert valid_tree_union_find(n, edges) == False


def test_linear_chain():
    """Test linear chain of nodes."""
    n = 4
    edges = [[0, 1], [1, 2], [2, 3]]
    assert valid_tree(n, edges) == True
    assert valid_tree_union_find(n, edges) == True


def test_simple_cycle():
    """Test simple cycle."""
    n = 3
    edges = [[0, 1], [1, 2], [2, 0]]
    assert valid_tree(n, edges) == False
    assert valid_tree_union_find(n, edges) == False


def test_too_many_edges():
    """Test with too many edges."""
    n = 3
    edges = [[0, 1], [0, 2], [1, 2]]
    assert valid_tree(n, edges) == False
    assert valid_tree_union_find(n, edges) == False


def test_too_few_edges():
    """Test with too few edges."""
    n = 4
    edges = [[0, 1], [1, 2]]
    assert valid_tree(n, edges) == False
    assert valid_tree_union_find(n, edges) == False


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
