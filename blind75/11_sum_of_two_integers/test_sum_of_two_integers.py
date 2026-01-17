import pytest
from sum_of_two_integers import get_sum


def test_example_1():
    """Test case from example 1."""
    a = 1
    b = 2
    result = get_sum(a, b)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    a = 2
    b = 3
    result = get_sum(a, b)
    assert result == 5


def test_example_3():
    """Test case from example 3 - negative and positive."""
    a = -1
    b = 1
    result = get_sum(a, b)
    assert result == 0


def test_both_positive():
    """Test with both positive numbers."""
    a = 5
    b = 3
    result = get_sum(a, b)
    assert result == 8


def test_both_negative():
    """Test with both negative numbers."""
    a = -5
    b = -3
    result = get_sum(a, b)
    assert result == -8


def test_zero_and_positive():
    """Test with zero and positive number."""
    a = 0
    b = 5
    result = get_sum(a, b)
    assert result == 5


def test_zero_and_negative():
    """Test with zero and negative number."""
    a = 0
    b = -5
    result = get_sum(a, b)
    assert result == -5


def test_both_zero():
    """Test with both numbers as zero."""
    a = 0
    b = 0
    result = get_sum(a, b)
    assert result == 0


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
