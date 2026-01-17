import pytest
from top_k_frequent import top_k_frequent


def test_example_1():
    """Test case from example 1."""
    nums = [1, 1, 1, 2, 2, 3]
    k = 2
    result = top_k_frequent(nums, k)
    assert sorted(result) == [1, 2]


def test_example_2():
    """Test case from example 2."""
    nums = [1]
    k = 1
    result = top_k_frequent(nums, k)
    assert result == [1]


def test_all_same_frequency():
    """Test when all elements have same frequency."""
    nums = [1, 2, 3, 4, 5]
    k = 3
    result = top_k_frequent(nums, k)
    assert len(result) == 3
    assert all(num in [1, 2, 3, 4, 5] for num in result)


def test_negative_numbers():
    """Test with negative numbers."""
    nums = [-1, -1, -1, 2, 2, 3]
    k = 2
    result = top_k_frequent(nums, k)
    assert sorted(result) == [-1, 2]


def test_k_equals_unique_elements():
    """Test when k equals number of unique elements."""
    nums = [4, 4, 4, 5, 5, 6]
    k = 3
    result = top_k_frequent(nums, k)
    assert sorted(result) == [4, 5, 6]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
