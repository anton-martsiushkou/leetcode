import pytest
from house_robber import rob


def test_example_1():
    """Test case from example 1."""
    nums = [1, 2, 3, 1]
    result = rob(nums)
    assert result == 4


def test_example_2():
    """Test case from example 2."""
    nums = [2, 7, 9, 3, 1]
    result = rob(nums)
    assert result == 12


def test_single_house():
    """Test with a single house."""
    nums = [5]
    result = rob(nums)
    assert result == 5


def test_two_houses():
    """Test with two houses."""
    nums = [1, 2]
    result = rob(nums)
    assert result == 2


def test_all_zeros():
    """Test with all zeros."""
    nums = [0, 0, 0, 0]
    result = rob(nums)
    assert result == 0


def test_alternating_high_values():
    """Test with alternating high values."""
    nums = [2, 1, 1, 2]
    result = rob(nums)
    assert result == 4


def test_increasing_values():
    """Test with increasing values."""
    nums = [1, 2, 3, 4, 5]
    result = rob(nums)
    assert result == 9


def test_decreasing_values():
    """Test with decreasing values."""
    nums = [5, 4, 3, 2, 1]
    result = rob(nums)
    assert result == 9


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
