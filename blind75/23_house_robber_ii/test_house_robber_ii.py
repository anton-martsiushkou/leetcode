import pytest
from house_robber_ii import rob


def test_example_1():
    """Test case from example 1."""
    nums = [2, 3, 2]
    result = rob(nums)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    nums = [1, 2, 3, 1]
    result = rob(nums)
    assert result == 4


def test_example_3():
    """Test case from example 3."""
    nums = [1, 2, 3]
    result = rob(nums)
    assert result == 3


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


def test_all_same_values():
    """Test with all same values."""
    nums = [5, 5, 5, 5]
    result = rob(nums)
    assert result == 10


def test_increasing_values():
    """Test with increasing values."""
    nums = [1, 2, 3, 4, 5]
    result = rob(nums)
    assert result == 8


def test_large_first_and_last():
    """Test with large values at first and last positions."""
    nums = [100, 1, 1, 1, 100]
    result = rob(nums)
    assert result == 100


def test_all_zeros():
    """Test with all zeros."""
    nums = [0, 0, 0, 0]
    result = rob(nums)
    assert result == 0


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
