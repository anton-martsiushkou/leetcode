import pytest
from climbing_stairs import climb_stairs


def test_example_1():
    """Test case from example 1."""
    n = 2
    result = climb_stairs(n)
    assert result == 2


def test_example_2():
    """Test case from example 2."""
    n = 3
    result = climb_stairs(n)
    assert result == 3


def test_example_3():
    """Test case from example 3."""
    n = 4
    result = climb_stairs(n)
    assert result == 5


def test_base_case_n1():
    """Test base case n=1."""
    n = 1
    result = climb_stairs(n)
    assert result == 1


def test_n5():
    """Test with n=5."""
    n = 5
    result = climb_stairs(n)
    assert result == 8


def test_n10():
    """Test with n=10."""
    n = 10
    result = climb_stairs(n)
    assert result == 89


def test_large_n20():
    """Test with large n=20."""
    n = 20
    result = climb_stairs(n)
    assert result == 10946


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
