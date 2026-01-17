import pytest
from three_sum import three_sum


def test_example_1():
    """Test case from example 1."""
    nums = [-1, 0, 1, 2, -1, -4]
    result = three_sum(nums)
    expected = [[-1, -1, 2], [-1, 0, 1]]
    assert sorted(result) == sorted(expected)


def test_example_2():
    """Test case from example 2."""
    nums = [0, 1, 1]
    result = three_sum(nums)
    assert result == []


def test_example_3():
    """Test case from example 3."""
    nums = [0, 0, 0]
    result = three_sum(nums)
    assert result == [[0, 0, 0]]


def test_all_negative():
    """Test with all negative numbers."""
    nums = [-1, -2, -3]
    result = three_sum(nums)
    assert result == []


def test_all_positive():
    """Test with all positive numbers."""
    nums = [1, 2, 3]
    result = three_sum(nums)
    assert result == []


def test_multiple_triplets():
    """Test with multiple valid triplets."""
    nums = [-2, 0, 1, 1, 2]
    result = three_sum(nums)
    expected = [[-2, 0, 2], [-2, 1, 1]]
    assert sorted(result) == sorted(expected)


def test_with_duplicates():
    """Test with duplicate elements."""
    nums = [-4, -1, -1, 0, 1, 2]
    result = three_sum(nums)
    expected = [[-1, -1, 2], [-1, 0, 1]]
    assert sorted(result) == sorted(expected)


def test_large_numbers():
    """Test with larger array."""
    nums = [-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5]
    result = three_sum(nums)
    expected = [
        [-5, 0, 5], [-5, 1, 4], [-5, 2, 3],
        [-4, -1, 5], [-4, 0, 4], [-4, 1, 3],
        [-3, -2, 5], [-3, -1, 4], [-3, 0, 3], [-3, 1, 2],
        [-2, -1, 3], [-2, 0, 2],
        [-1, 0, 1]
    ]
    assert sorted(result) == sorted(expected)


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
