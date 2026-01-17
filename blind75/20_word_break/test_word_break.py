import pytest
from word_break import word_break


def test_example_1():
    """Test case from example 1."""
    s = "leetcode"
    word_dict = ["leet", "code"]
    result = word_break(s, word_dict)
    assert result is True


def test_example_2():
    """Test case from example 2."""
    s = "applepenapple"
    word_dict = ["apple", "pen"]
    result = word_break(s, word_dict)
    assert result is True


def test_example_3():
    """Test case from example 3."""
    s = "catsandog"
    word_dict = ["cats", "dog", "sand", "and", "cat"]
    result = word_break(s, word_dict)
    assert result is False


def test_single_word_exact_match():
    """Test with single word exact match."""
    s = "apple"
    word_dict = ["apple"]
    result = word_break(s, word_dict)
    assert result is True


def test_single_word_no_match():
    """Test with single word no match."""
    s = "apple"
    word_dict = ["orange"]
    result = word_break(s, word_dict)
    assert result is False


def test_empty_string():
    """Test with empty string."""
    s = ""
    word_dict = ["a"]
    result = word_break(s, word_dict)
    assert result is True


def test_reuse_word_multiple_times():
    """Test reusing word multiple times."""
    s = "aaaaaaa"
    word_dict = ["aa", "aaa"]
    result = word_break(s, word_dict)
    assert result is True


def test_complex_segmentation():
    """Test with complex segmentation."""
    s = "cars"
    word_dict = ["car", "ca", "rs"]
    result = word_break(s, word_dict)
    assert result is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
