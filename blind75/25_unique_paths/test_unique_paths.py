import pytest
from unique_paths import unique_paths


def test_example_1():
    """Test case from example 1."""
    m = 3
    n = 7
    result = unique_paths(m, n)
    assert result == 28


def test_example_2():
    """Test case from example 2."""
    m = 3
    n = 2
    result = unique_paths(m, n)
    assert result == 3


def test_1x1_grid():
    """Test with 1x1 grid."""
    m = 1
    n = 1
    result = unique_paths(m, n)
    assert result == 1


def test_1x10_grid():
    """Test with 1x10 grid."""
    m = 1
    n = 10
    result = unique_paths(m, n)
    assert result == 1


def test_10x1_grid():
    """Test with 10x1 grid."""
    m = 10
    n = 1
    result = unique_paths(m, n)
    assert result == 1


def test_square_grid_3x3():
    """Test with square 3x3 grid."""
    m = 3
    n = 3
    result = unique_paths(m, n)
    assert result == 6


def test_square_grid_4x4():
    """Test with square 4x4 grid."""
    m = 4
    n = 4
    result = unique_paths(m, n)
    assert result == 20


def test_2x2_grid():
    """Test with 2x2 grid."""
    m = 2
    n = 2
    result = unique_paths(m, n)
    assert result == 2


def test_large_grid():
    """Test with large grid."""
    m = 10
    n = 10
    result = unique_paths(m, n)
    assert result == 48620


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
