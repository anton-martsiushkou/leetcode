import pytest
from decode_ways import num_decodings


def test_example_1():
    """Test case from example 1."""
    s = "12"
    result = num_decodings(s)
    assert result == 2


def test_example_2():
    """Test case from example 2."""
    s = "226"
    result = num_decodings(s)
    assert result == 3


def test_example_3():
    """Test case from example 3."""
    s = "06"
    result = num_decodings(s)
    assert result == 0


def test_single_digit():
    """Test with a single digit."""
    s = "1"
    result = num_decodings(s)
    assert result == 1


def test_leading_zero():
    """Test with leading zero."""
    s = "0"
    result = num_decodings(s)
    assert result == 0


def test_valid_two_digit():
    """Test with valid two digit number."""
    s = "10"
    result = num_decodings(s)
    assert result == 1


def test_invalid_zero_in_middle():
    """Test with invalid zero in middle."""
    s = "100"
    result = num_decodings(s)
    assert result == 0


def test_multiple_valid_decodings():
    """Test with multiple valid decodings."""
    s = "111111"
    result = num_decodings(s)
    assert result == 13


def test_boundary_case():
    """Test boundary case."""
    s = "2611055971756562"
    result = num_decodings(s)
    assert result == 4


def test_single_character_valid():
    """Test single character valid."""
    s = "5"
    result = num_decodings(s)
    assert result == 1


def test_two_digits_out_of_range():
    """Test two digits out of range."""
    s = "27"
    result = num_decodings(s)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
