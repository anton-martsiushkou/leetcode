import pytest
from longest_palindromic_substring import longest_palindrome


def is_palindrome(s: str) -> bool:
    """Check if a string is a palindrome."""
    return s == s[::-1]


def test_example_1():
    """Test case from example 1."""
    s = "babad"
    result = longest_palindrome(s)
    assert len(result) == 3
    assert is_palindrome(result)
    assert result in ["bab", "aba"]


def test_example_2():
    """Test case from example 2."""
    s = "cbbd"
    result = longest_palindrome(s)
    assert result == "bb"
    assert is_palindrome(result)


def test_example_3():
    """Test case from example 3 - single character."""
    s = "a"
    result = longest_palindrome(s)
    assert result == "a"


def test_example_4():
    """Test case from example 4 - multiple single chars."""
    s = "ac"
    result = longest_palindrome(s)
    assert len(result) == 1
    assert result in ["a", "c"]


def test_all_same_characters():
    """Test with all same characters."""
    s = "aaaa"
    result = longest_palindrome(s)
    assert result == "aaaa"
    assert is_palindrome(result)


def test_no_palindrome_except_singles():
    """Test with no palindrome except single characters."""
    s = "abcd"
    result = longest_palindrome(s)
    assert len(result) == 1
    assert is_palindrome(result)


def test_palindrome_at_start():
    """Test with palindrome at the start."""
    s = "abacabad"
    result = longest_palindrome(s)
    assert result == "abacaba"
    assert is_palindrome(result)


def test_palindrome_at_end():
    """Test with palindrome at the end."""
    s = "xyracecar"
    result = longest_palindrome(s)
    assert result == "racecar"
    assert is_palindrome(result)


def test_even_length_palindrome():
    """Test with even length palindrome."""
    s = "abba"
    result = longest_palindrome(s)
    assert result == "abba"
    assert is_palindrome(result)


def test_odd_length_palindrome():
    """Test with odd length palindrome."""
    s = "racecar"
    result = longest_palindrome(s)
    assert result == "racecar"
    assert is_palindrome(result)


def test_multiple_palindromes_same_length():
    """Test with multiple palindromes of same length."""
    s = "abacdfgdcaba"
    result = longest_palindrome(s)
    assert len(result) == 3
    assert is_palindrome(result)


def test_two_character_palindrome():
    """Test with two character palindrome."""
    s = "bb"
    result = longest_palindrome(s)
    assert result == "bb"
    assert is_palindrome(result)


def test_two_different_characters():
    """Test with two different characters."""
    s = "ab"
    result = longest_palindrome(s)
    assert len(result) == 1
    assert is_palindrome(result)


def test_complex_string():
    """Test with a complex string."""
    s = "bananas"
    result = longest_palindrome(s)
    assert len(result) == 5
    assert result == "anana"
    assert is_palindrome(result)


def test_result_is_always_palindrome():
    """Verify that all results are valid palindromes."""
    test_cases = ["babad", "cbbd", "a", "ac", "aaaa", "abcd", "racecar"]
    for s in test_cases:
        result = longest_palindrome(s)
        assert is_palindrome(result), f"Result '{result}' for input '{s}' is not a palindrome"
        assert len(result) > 0, f"Result should not be empty for input '{s}'"


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
