import pytest
from pacific_atlantic import pacific_atlantic, pacific_atlantic_bfs


def sort_result(result):
    """Sort result for comparison."""
    return sorted(result, key=lambda x: (x[0], x[1]))


def test_example_1():
    """Test case from example 1."""
    heights = [
        [1, 2, 2, 3, 5],
        [3, 2, 3, 4, 4],
        [2, 4, 5, 3, 1],
        [6, 7, 1, 4, 5],
        [5, 1, 1, 2, 4]
    ]
    expected = [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]
    result = pacific_atlantic(heights)
    assert sort_result(result) == sort_result(expected)


def test_example_2():
    """Test case from example 2 - single cell."""
    heights = [[1]]
    expected = [[0, 0]]
    result = pacific_atlantic(heights)
    assert result == expected


def test_2x2_grid():
    """Test with 2x2 grid."""
    heights = [[1, 2], [3, 4]]
    expected = [[0, 1], [1, 0], [1, 1]]
    result = pacific_atlantic(heights)
    assert sort_result(result) == sort_result(expected)


def test_all_same_height():
    """Test with all cells at same height."""
    heights = [[1, 1], [1, 1]]
    expected = [[0, 0], [0, 1], [1, 0], [1, 1]]
    result = pacific_atlantic(heights)
    assert sort_result(result) == sort_result(expected)


def test_empty_grid():
    """Test with empty grid."""
    heights = []
    expected = []
    result = pacific_atlantic(heights)
    assert result == expected


def test_bfs_example_1():
    """Test BFS implementation with example 1."""
    heights = [
        [1, 2, 2, 3, 5],
        [3, 2, 3, 4, 4],
        [2, 4, 5, 3, 1],
        [6, 7, 1, 4, 5],
        [5, 1, 1, 2, 4]
    ]
    expected = [[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]]
    result = pacific_atlantic_bfs(heights)
    assert sort_result(result) == sort_result(expected)


def test_bfs_single_cell():
    """Test BFS with single cell."""
    heights = [[1]]
    expected = [[0, 0]]
    result = pacific_atlantic_bfs(heights)
    assert result == expected


def test_bfs_all_same():
    """Test BFS with all same height."""
    heights = [[1, 1], [1, 1]]
    expected = [[0, 0], [0, 1], [1, 0], [1, 1]]
    result = pacific_atlantic_bfs(heights)
    assert sort_result(result) == sort_result(expected)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
