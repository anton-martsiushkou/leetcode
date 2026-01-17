import pytest
from character_replacement import character_replacement


def test_example_1():
    """Test case from example 1."""
    s = "ABAB"
    k = 2
    result = character_replacement(s, k)
    assert result == 4


def test_example_2():
    """Test case from example 2."""
    s = "AABABBA"
    k = 1
    result = character_replacement(s, k)
    assert result == 4


def test_all_same_characters():
    """Test with all same characters."""
    s = "AAAA"
    k = 0
    result = character_replacement(s, k)
    assert result == 4


def test_k_equals_length():
    """Test when k equals string length."""
    s = "ABCD"
    k = 4
    result = character_replacement(s, k)
    assert result == 4


def test_single_character():
    """Test with single character."""
    s = "A"
    k = 0
    result = character_replacement(s, k)
    assert result == 1


def test_no_replacements():
    """Test with no replacements allowed."""
    s = "ABCDE"
    k = 0
    result = character_replacement(s, k)
    assert result == 1


def test_alternating_characters():
    """Test with alternating characters."""
    s = "ABABAB"
    k = 2
    result = character_replacement(s, k)
    assert result == 5


def test_large_k():
    """Test with large k value."""
    s = "AABA"
    k = 10
    result = character_replacement(s, k)
    assert result == 4


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
