import pytest
from word_dictionary import WordDictionary


def test_example_1():
    """Test case from example 1."""
    wd = WordDictionary()
    wd.add_word("bad")
    wd.add_word("dad")
    wd.add_word("mad")

    assert wd.search("pad") is False
    assert wd.search("bad") is True
    assert wd.search(".ad") is True
    assert wd.search("b..") is True


def test_exact_matches():
    """Test exact word matches."""
    wd = WordDictionary()
    wd.add_word("bad")
    wd.add_word("dad")
    wd.add_word("mad")

    assert wd.search("bad") is True
    assert wd.search("dad") is True
    assert wd.search("mad") is True
    assert wd.search("pad") is False
    assert wd.search("bat") is False


def test_wildcard_at_start():
    """Test wildcard at the start of the word."""
    wd = WordDictionary()
    wd.add_word("bad")
    wd.add_word("dad")
    wd.add_word("mad")

    assert wd.search(".ad") is True
    assert wd.search(".at") is False


def test_wildcard_at_end():
    """Test wildcard at the end of the word."""
    wd = WordDictionary()
    wd.add_word("bad")
    wd.add_word("bat")
    wd.add_word("ban")

    assert wd.search("ba.") is True
    assert wd.search("be.") is False


def test_multiple_wildcards():
    """Test multiple wildcards in the word."""
    wd = WordDictionary()
    wd.add_word("bad")
    wd.add_word("dad")

    assert wd.search("b..") is True
    assert wd.search("d..") is True
    assert wd.search("...") is True
    assert wd.search("....") is False


def test_various_lengths():
    """Test words of various lengths."""
    wd = WordDictionary()
    words = ["a", "ab", "abc", "abcd", "abcde"]

    for word in words:
        wd.add_word(word)

    assert wd.search("a") is True
    assert wd.search(".") is True
    assert wd.search("ab") is True
    assert wd.search(".b") is True
    assert wd.search("..") is True
    assert wd.search("abcde") is True
    assert wd.search("a.c.e") is True
    assert wd.search("a.d") is False
    assert wd.search("abcdef") is False


def test_no_branch():
    """Test when there's no matching branch."""
    wd = WordDictionary()
    wd.add_word("test")

    assert wd.search("test") is True
    assert wd.search("t..t") is True
    assert wd.search("....") is True
    assert wd.search("best") is False
    assert wd.search("t..") is False
    assert wd.search("tests") is False


def test_empty_dictionary():
    """Test searching in an empty dictionary."""
    wd = WordDictionary()

    assert wd.search("anything") is False
    assert wd.search(".") is False
    assert wd.search("...") is False


def test_overlapping_words():
    """Test with overlapping words."""
    wd = WordDictionary()
    wd.add_word("test")
    wd.add_word("testing")
    wd.add_word("tester")

    assert wd.search("test") is True
    assert wd.search("testing") is True
    assert wd.search("tester") is True
    assert wd.search("test..") is False  # "test" is 4 chars, "test.." is 6
    assert wd.search("test...") is True  # matches "testing" and "tester"


def test_duplicate_additions():
    """Test adding the same word multiple times."""
    wd = WordDictionary()
    wd.add_word("word")
    wd.add_word("word")
    wd.add_word("word")

    assert wd.search("word") is True
    assert wd.search("w...") is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
