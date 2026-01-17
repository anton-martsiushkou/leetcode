import pytest
from missing_number import missing_number


def test_example_1():
    """Test case from example 1."""
    nums = [3, 0, 1]
    result = missing_number(nums)
    assert result == 2


def test_example_2():
    """Test case from example 2."""
    nums = [0, 1]
    result = missing_number(nums)
    assert result == 2


def test_example_3():
    """Test case from example 3."""
    nums = [9, 6, 4, 2, 3, 5, 7, 0, 1]
    result = missing_number(nums)
    assert result == 8


def test_missing_0():
    """Test when 0 is missing."""
    nums = [1, 2, 3, 4, 5]
    result = missing_number(nums)
    assert result == 0


def test_missing_last():
    """Test when the last number is missing."""
    nums = [0, 1, 2, 3, 4]
    result = missing_number(nums)
    assert result == 5


def test_single_element_missing_0():
    """Test with single element, missing 0."""
    nums = [1]
    result = missing_number(nums)
    assert result == 0


def test_single_element_missing_1():
    """Test with single element, missing 1."""
    nums = [0]
    result = missing_number(nums)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
