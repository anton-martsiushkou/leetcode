import pytest
from container_with_most_water import max_area


def test_example_1():
    """Test case from example 1."""
    height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
    result = max_area(height)
    assert result == 49


def test_example_2():
    """Test case from example 2."""
    height = [1, 1]
    result = max_area(height)
    assert result == 1


def test_ascending_order():
    """Test with heights in ascending order."""
    height = [1, 2, 3, 4, 5]
    result = max_area(height)
    assert result == 6


def test_descending_order():
    """Test with heights in descending order."""
    height = [5, 4, 3, 2, 1]
    result = max_area(height)
    assert result == 6


def test_all_same_height():
    """Test with all same heights."""
    height = [5, 5, 5, 5]
    result = max_area(height)
    assert result == 15


def test_two_tall_ends():
    """Test with tall heights at both ends."""
    height = [10, 1, 1, 1, 10]
    result = max_area(height)
    assert result == 40


def test_tall_in_middle():
    """Test with tall height in the middle."""
    height = [1, 10, 1]
    result = max_area(height)
    assert result == 2


def test_alternating_heights():
    """Test with alternating heights."""
    height = [1, 3, 2, 4, 3, 5]
    result = max_area(height)
    assert result == 12


def test_with_zeros():
    """Test with zeros in the array."""
    height = [0, 2, 0, 3, 0, 4]
    result = max_area(height)
    assert result == 6


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
