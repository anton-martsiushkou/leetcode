import pytest
from set_matrix_zeroes import set_zeroes


def test_example_1():
    """Test case from example 1."""
    matrix = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    set_zeroes(matrix)
    assert matrix == [[1, 0, 1], [0, 0, 0], [1, 0, 1]]


def test_example_2():
    """Test case from example 2."""
    matrix = [[0, 1, 2, 0], [3, 4, 5, 2], [1, 3, 1, 5]]
    set_zeroes(matrix)
    assert matrix == [[0, 0, 0, 0], [0, 4, 5, 0], [0, 3, 1, 0]]


def test_single_element_zero():
    """Test with single element that is zero."""
    matrix = [[0]]
    set_zeroes(matrix)
    assert matrix == [[0]]


def test_single_element_non_zero():
    """Test with single element that is non-zero."""
    matrix = [[1]]
    set_zeroes(matrix)
    assert matrix == [[1]]


def test_all_zeros():
    """Test with all zeros."""
    matrix = [[0, 0], [0, 0]]
    set_zeroes(matrix)
    assert matrix == [[0, 0], [0, 0]]


def test_no_zeros():
    """Test with no zeros."""
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    set_zeroes(matrix)
    assert matrix == [[1, 2, 3], [4, 5, 6], [7, 8, 9]]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
