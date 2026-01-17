import pytest
from search_in_rotated_sorted_array import search


def test_example_1():
    """Test case from example 1."""
    nums = [4, 5, 6, 7, 0, 1, 2]
    target = 0
    result = search(nums, target)
    assert result == 4


def test_example_2():
    """Test case from example 2."""
    nums = [4, 5, 6, 7, 0, 1, 2]
    target = 3
    result = search(nums, target)
    assert result == -1


def test_example_3():
    """Test case from example 3."""
    nums = [1]
    target = 0
    result = search(nums, target)
    assert result == -1


def test_single_element_found():
    """Test with single element that matches target."""
    nums = [1]
    target = 1
    result = search(nums, target)
    assert result == 0


def test_target_at_beginning():
    """Test with target at the beginning."""
    nums = [4, 5, 6, 7, 0, 1, 2]
    target = 4
    result = search(nums, target)
    assert result == 0


def test_target_at_end():
    """Test with target at the end."""
    nums = [4, 5, 6, 7, 0, 1, 2]
    target = 2
    result = search(nums, target)
    assert result == 6


def test_no_rotation():
    """Test with no rotation."""
    nums = [1, 2, 3, 4, 5]
    target = 3
    result = search(nums, target)
    assert result == 2


def test_rotated_once():
    """Test with array rotated once."""
    nums = [5, 1, 2, 3, 4]
    target = 1
    result = search(nums, target)
    assert result == 1


def test_two_elements_rotated():
    """Test with two elements rotated."""
    nums = [3, 1]
    target = 1
    result = search(nums, target)
    assert result == 1


def test_two_elements_not_found():
    """Test with two elements, target not found."""
    nums = [1, 3]
    target = 2
    result = search(nums, target)
    assert result == -1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
