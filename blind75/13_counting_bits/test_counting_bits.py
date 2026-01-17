import pytest
from counting_bits import count_bits


def test_example_1():
    """Test case from example 1."""
    n = 2
    result = count_bits(n)
    assert result == [0, 1, 1]


def test_example_2():
    """Test case from example 2."""
    n = 5
    result = count_bits(n)
    assert result == [0, 1, 1, 2, 1, 2]


def test_n_is_0():
    """Test when n is 0."""
    n = 0
    result = count_bits(n)
    assert result == [0]


def test_n_is_1():
    """Test when n is 1."""
    n = 1
    result = count_bits(n)
    assert result == [0, 1]


def test_n_is_8():
    """Test when n is 8."""
    n = 8
    result = count_bits(n)
    assert result == [0, 1, 1, 2, 1, 2, 2, 3, 1]


def test_n_is_15():
    """Test when n is 15."""
    n = 15
    result = count_bits(n)
    assert result == [0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4]


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
