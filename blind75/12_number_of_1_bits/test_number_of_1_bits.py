import pytest
from number_of_1_bits import hamming_weight


def test_example_1():
    """Test case from example 1."""
    n = 11
    result = hamming_weight(n)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    n = 128
    result = hamming_weight(n)
    assert result == 1


def test_example_3():
    """Test case from example 3."""
    n = 2147483645
    result = hamming_weight(n)
    assert result == 30


def test_single_bit_set():
    """Test with a single bit set."""
    n = 1
    result = hamming_weight(n)
    assert result == 1


def test_all_bits_set_in_8_bit():
    """Test with all bits set in an 8-bit number."""
    n = 255
    result = hamming_weight(n)
    assert result == 8


def test_power_of_two():
    """Test with a power of two."""
    n = 1024
    result = hamming_weight(n)
    assert result == 1


def test_alternating_bits():
    """Test with alternating bits."""
    n = 21845  # 0101010101010101
    result = hamming_weight(n)
    assert result == 8


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
