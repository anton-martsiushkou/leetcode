import pytest
from spiral_matrix import spiral_order


def test_example_1():
    """Test case from example 1."""
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    result = spiral_order(matrix)
    assert result == [1, 2, 3, 6, 9, 8, 7, 4, 5]


def test_example_2():
    """Test case from example 2."""
    matrix = [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]
    result = spiral_order(matrix)
    assert result == [1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7]


def test_single_element():
    """Test with single element."""
    matrix = [[1]]
    result = spiral_order(matrix)
    assert result == [1]


def test_single_row():
    """Test with single row."""
    matrix = [[1, 2, 3, 4]]
    result = spiral_order(matrix)
    assert result == [1, 2, 3, 4]


def test_single_column():
    """Test with single column."""
    matrix = [[1], [2], [3], [4]]
    result = spiral_order(matrix)
    assert result == [1, 2, 3, 4]


def test_2x2_matrix():
    """Test with 2x2 matrix."""
    matrix = [[1, 2], [3, 4]]
    result = spiral_order(matrix)
    assert result == [1, 2, 4, 3]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
