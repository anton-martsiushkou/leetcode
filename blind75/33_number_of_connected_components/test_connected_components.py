import pytest
from connected_components import (
    count_components,
    count_components_bfs,
    count_components_union_find
)


def test_example_1():
    """Test case from example 1 - two components."""
    n = 5
    edges = [[0, 1], [1, 2], [3, 4]]
    assert count_components(n, edges) == 2


def test_example_2():
    """Test case from example 2 - one component."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [3, 4]]
    assert count_components(n, edges) == 1


def test_no_edges():
    """Test with no edges - all nodes isolated."""
    n = 4
    edges = []
    assert count_components(n, edges) == 4


def test_single_node():
    """Test with single node."""
    n = 1
    edges = []
    assert count_components(n, edges) == 1


def test_complete_graph():
    """Test with complete graph."""
    n = 4
    edges = [[0, 1], [0, 2], [0, 3], [1, 2], [1, 3], [2, 3]]
    assert count_components(n, edges) == 1


def test_three_components():
    """Test with three separate components."""
    n = 6
    edges = [[0, 1], [2, 3], [4, 5]]
    assert count_components(n, edges) == 3


def test_star_graph():
    """Test with star graph."""
    n = 5
    edges = [[0, 1], [0, 2], [0, 3], [0, 4]]
    assert count_components(n, edges) == 1


def test_bfs_example_1():
    """Test BFS with example 1."""
    n = 5
    edges = [[0, 1], [1, 2], [3, 4]]
    assert count_components_bfs(n, edges) == 2


def test_bfs_example_2():
    """Test BFS with example 2."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [3, 4]]
    assert count_components_bfs(n, edges) == 1


def test_bfs_no_edges():
    """Test BFS with no edges."""
    n = 4
    edges = []
    assert count_components_bfs(n, edges) == 4


def test_bfs_complete_graph():
    """Test BFS with complete graph."""
    n = 4
    edges = [[0, 1], [0, 2], [0, 3], [1, 2], [1, 3], [2, 3]]
    assert count_components_bfs(n, edges) == 1


def test_union_find_example_1():
    """Test Union Find with example 1."""
    n = 5
    edges = [[0, 1], [1, 2], [3, 4]]
    assert count_components_union_find(n, edges) == 2


def test_union_find_example_2():
    """Test Union Find with example 2."""
    n = 5
    edges = [[0, 1], [1, 2], [2, 3], [3, 4]]
    assert count_components_union_find(n, edges) == 1


def test_union_find_no_edges():
    """Test Union Find with no edges."""
    n = 4
    edges = []
    assert count_components_union_find(n, edges) == 4


def test_union_find_single_node():
    """Test Union Find with single node."""
    n = 1
    edges = []
    assert count_components_union_find(n, edges) == 1


def test_union_find_three_components():
    """Test Union Find with three components."""
    n = 6
    edges = [[0, 1], [2, 3], [4, 5]]
    assert count_components_union_find(n, edges) == 3


def test_union_find_star():
    """Test Union Find with star graph."""
    n = 5
    edges = [[0, 1], [0, 2], [0, 3], [0, 4]]
    assert count_components_union_find(n, edges) == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
