import pytest
from min_window import min_window


def test_example_1():
    """Test case from example 1."""
    s = "ADOBECODEBANC"
    t = "ABC"
    result = min_window(s, t)
    assert result == "BANC"


def test_example_2():
    """Test case from example 2."""
    s = "a"
    t = "a"
    result = min_window(s, t)
    assert result == "a"


def test_example_3():
    """Test case from example 3."""
    s = "a"
    t = "aa"
    result = min_window(s, t)
    assert result == ""


def test_no_valid_window():
    """Test when no valid window exists."""
    s = "abc"
    t = "xyz"
    result = min_window(s, t)
    assert result == ""


def test_entire_string_is_minimum():
    """Test when entire string is the minimum window."""
    s = "abc"
    t = "abc"
    result = min_window(s, t)
    assert result == "abc"


def test_duplicate_characters_in_t():
    """Test with duplicate characters in t."""
    s = "aaabbbc"
    t = "abc"
    result = min_window(s, t)
    assert result == "abbc"


def test_multiple_valid_windows():
    """Test with multiple valid windows."""
    s = "abcabdebac"
    t = "abc"
    result = min_window(s, t)
    assert result == "bac"


def test_s_shorter_than_t():
    """Test when s is shorter than t."""
    s = "ab"
    t = "abc"
    result = min_window(s, t)
    assert result == ""


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
