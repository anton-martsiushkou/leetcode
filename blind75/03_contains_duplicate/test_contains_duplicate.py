import pytest
from contains_duplicate import contains_duplicate


def test_example_1_has_duplicate():
    """Test case from example 1."""
    nums = [1, 2, 3, 1]
    assert contains_duplicate(nums) is True


def test_example_2_no_duplicate():
    """Test case from example 2."""
    nums = [1, 2, 3, 4]
    assert contains_duplicate(nums) is False


def test_example_3_multiple_duplicates():
    """Test case from example 3."""
    nums = [1, 1, 1, 3, 3, 4, 3, 2, 4, 2]
    assert contains_duplicate(nums) is True


def test_single_element():
    """Test with single element - cannot have duplicates."""
    nums = [1]
    assert contains_duplicate(nums) is False


def test_two_same_elements():
    """Test with two identical elements."""
    nums = [1, 1]
    assert contains_duplicate(nums) is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
