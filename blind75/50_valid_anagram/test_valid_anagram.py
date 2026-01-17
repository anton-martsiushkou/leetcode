import pytest
from valid_anagram import is_anagram


def test_example_1():
    """Test case from example 1."""
    s = "anagram"
    t = "nagaram"
    result = is_anagram(s, t)
    assert result is True


def test_example_2():
    """Test case from example 2."""
    s = "rat"
    t = "car"
    result = is_anagram(s, t)
    assert result is False


def test_empty_strings():
    """Test with empty strings."""
    s = ""
    t = ""
    result = is_anagram(s, t)
    assert result is True


def test_single_character_same():
    """Test with single same character."""
    s = "a"
    t = "a"
    result = is_anagram(s, t)
    assert result is True


def test_single_character_different():
    """Test with single different character."""
    s = "a"
    t = "b"
    result = is_anagram(s, t)
    assert result is False


def test_different_lengths():
    """Test with different length strings."""
    s = "abc"
    t = "abcd"
    result = is_anagram(s, t)
    assert result is False


def test_same_characters_different_frequencies():
    """Test with same characters but different frequencies."""
    s = "aab"
    t = "abb"
    result = is_anagram(s, t)
    assert result is False


def test_all_same_character():
    """Test with all same character."""
    s = "aaaa"
    t = "aaaa"
    result = is_anagram(s, t)
    assert result is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
