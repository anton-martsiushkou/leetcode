import pytest
from reverse_bits import reverse_bits


def test_example_1():
    """Test case from example 1."""
    n = 0b00000010100101000001111010011100
    result = reverse_bits(n)
    assert result == 0b00111001011110000010100101000000


def test_example_2():
    """Test case from example 2."""
    n = 0b11111111111111111111111111111101
    result = reverse_bits(n)
    assert result == 0b10111111111111111111111111111111


def test_all_zeros():
    """Test with all zeros."""
    n = 0b00000000000000000000000000000000
    result = reverse_bits(n)
    assert result == 0b00000000000000000000000000000000


def test_all_ones():
    """Test with all ones."""
    n = 0b11111111111111111111111111111111
    result = reverse_bits(n)
    assert result == 0b11111111111111111111111111111111


def test_single_bit_set_at_start():
    """Test with single bit set at the start."""
    n = 0b10000000000000000000000000000000
    result = reverse_bits(n)
    assert result == 0b00000000000000000000000000000001


def test_single_bit_set_at_end():
    """Test with single bit set at the end."""
    n = 0b00000000000000000000000000000001
    result = reverse_bits(n)
    assert result == 0b10000000000000000000000000000000


def test_alternating_bits():
    """Test with alternating bits."""
    n = 0b10101010101010101010101010101010
    result = reverse_bits(n)
    assert result == 0b01010101010101010101010101010101


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
