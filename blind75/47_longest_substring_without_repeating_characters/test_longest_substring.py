import pytest
from longest_substring import length_of_longest_substring


def test_example_1():
    """Test case from example 1."""
    s = "abcabcbb"
    result = length_of_longest_substring(s)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    s = "bbbbb"
    result = length_of_longest_substring(s)
    assert result == 1


def test_example_3():
    """Test case from example 3."""
    s = "pwwkew"
    result = length_of_longest_substring(s)
    assert result == 3


def test_empty_string():
    """Test with empty string."""
    s = ""
    result = length_of_longest_substring(s)
    assert result == 0


def test_single_character():
    """Test with single character."""
    s = "a"
    result = length_of_longest_substring(s)
    assert result == 1


def test_all_unique_characters():
    """Test with all unique characters."""
    s = "abcdef"
    result = length_of_longest_substring(s)
    assert result == 6


def test_with_spaces_and_symbols():
    """Test with spaces and symbols."""
    s = "a b c a"
    result = length_of_longest_substring(s)
    assert result == 3


def test_long_string_with_repeat_at_end():
    """Test with long string with repeat at end."""
    s = "abcdefga"
    result = length_of_longest_substring(s)
    assert result == 7


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
