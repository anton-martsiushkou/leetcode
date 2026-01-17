import pytest
from longest_common_subsequence import longest_common_subsequence


def test_example_1():
    """Test case from example 1."""
    text1 = "abcde"
    text2 = "ace"
    result = longest_common_subsequence(text1, text2)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    text1 = "abc"
    text2 = "abc"
    result = longest_common_subsequence(text1, text2)
    assert result == 3


def test_example_3():
    """Test case from example 3."""
    text1 = "abc"
    text2 = "def"
    result = longest_common_subsequence(text1, text2)
    assert result == 0


def test_one_character_match():
    """Test with one character match."""
    text1 = "abc"
    text2 = "axc"
    result = longest_common_subsequence(text1, text2)
    assert result == 2


def test_single_character_strings():
    """Test with single character strings that match."""
    text1 = "a"
    text2 = "a"
    result = longest_common_subsequence(text1, text2)
    assert result == 1


def test_single_character_no_match():
    """Test with single character strings that don't match."""
    text1 = "a"
    text2 = "b"
    result = longest_common_subsequence(text1, text2)
    assert result == 0


def test_longer_strings():
    """Test with longer strings."""
    text1 = "abcdefghij"
    text2 = "ecdgi"
    result = longest_common_subsequence(text1, text2)
    assert result == 4


def test_reversed_strings():
    """Test with reversed strings."""
    text1 = "abc"
    text2 = "cba"
    result = longest_common_subsequence(text1, text2)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
