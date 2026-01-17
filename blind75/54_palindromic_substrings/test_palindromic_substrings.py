import pytest
from palindromic_substrings import count_substrings


def test_example_1():
    """Test case from example 1 - no palindromes except single chars."""
    s = "abc"
    result = count_substrings(s)
    assert result == 3


def test_example_2():
    """Test case from example 2 - all same characters."""
    s = "aaa"
    result = count_substrings(s)
    assert result == 6


def test_example_3():
    """Test case from example 3 - word palindrome."""
    s = "racecar"
    result = count_substrings(s)
    assert result == 10


def test_single_character():
    """Test with a single character."""
    s = "a"
    result = count_substrings(s)
    assert result == 1


def test_two_same_characters():
    """Test with two same characters."""
    s = "aa"
    result = count_substrings(s)
    assert result == 3


def test_two_different_characters():
    """Test with two different characters."""
    s = "ab"
    result = count_substrings(s)
    assert result == 2


def test_palindrome_at_start():
    """Test with palindrome at the start."""
    s = "aab"
    result = count_substrings(s)
    assert result == 4


def test_palindrome_at_end():
    """Test with palindrome at the end."""
    s = "baa"
    result = count_substrings(s)
    assert result == 4


def test_palindrome_in_middle():
    """Test with palindrome in the middle."""
    s = "baaab"
    result = count_substrings(s)
    assert result == 9


def test_multiple_palindromes():
    """Test with multiple palindromes."""
    s = "abba"
    result = count_substrings(s)
    assert result == 6


def test_long_repeated_characters():
    """Test with long string of repeated characters."""
    s = "aaaa"
    result = count_substrings(s)
    assert result == 10


def test_complex_palindrome():
    """Test with complex palindrome."""
    s = "abcba"
    result = count_substrings(s)
    assert result == 7


def test_empty_string():
    """Test with empty string."""
    s = ""
    result = count_substrings(s)
    assert result == 0


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
