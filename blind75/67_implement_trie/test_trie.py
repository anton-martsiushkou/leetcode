import pytest
from trie import Trie


def test_example_1():
    """Test case from example 1."""
    trie = Trie()
    trie.insert("apple")
    assert trie.search("apple") is True
    assert trie.search("app") is False
    assert trie.starts_with("app") is True
    trie.insert("app")
    assert trie.search("app") is True


def test_multiple_words():
    """Test with multiple words."""
    trie = Trie()
    words = ["apple", "app", "application", "apply", "banana", "band"]

    for word in words:
        trie.insert(word)

    # Test all inserted words can be found
    for word in words:
        assert trie.search(word) is True, f"Expected to find word '{word}'"

    # Test non-existent words
    non_existent = ["appl", "ban", "banan", "apps", "bandana"]
    for word in non_existent:
        assert trie.search(word) is False, f"Did not expect to find word '{word}'"


def test_prefixes():
    """Test prefix functionality."""
    trie = Trie()
    words = ["apple", "app", "application", "apply", "banana", "band"]

    for word in words:
        trie.insert(word)

    # Test valid prefixes
    prefixes = ["app", "appl", "appli", "ban", "band"]
    for prefix in prefixes:
        assert trie.starts_with(prefix) is True, f"Expected '{prefix}' to be a valid prefix"

    # Test invalid prefixes
    invalid_prefixes = ["orange", "grape", "car"]
    for prefix in invalid_prefixes:
        assert trie.starts_with(prefix) is False, f"Did not expect '{prefix}' to be a valid prefix"


def test_empty_and_single_char():
    """Test with empty string and single character."""
    trie = Trie()

    # Test single character
    trie.insert("a")
    assert trie.search("a") is True
    assert trie.starts_with("a") is True

    # Test that empty search doesn't match
    assert trie.search("") is False


def test_overlapping_words():
    """Test words that are prefixes of each other."""
    trie = Trie()

    trie.insert("test")
    trie.insert("testing")
    trie.insert("tester")

    assert trie.search("test") is True
    assert trie.search("testing") is True
    assert trie.search("tester") is True
    assert trie.search("tes") is False
    assert trie.starts_with("test") is True


def test_case_sensitivity():
    """Test that trie is case-sensitive."""
    trie = Trie()

    trie.insert("hello")
    assert trie.search("hello") is True
    assert trie.search("Hello") is False
    assert trie.starts_with("hel") is True
    assert trie.starts_with("Hel") is False


def test_duplicate_insertions():
    """Test inserting the same word multiple times."""
    trie = Trie()

    trie.insert("word")
    trie.insert("word")
    trie.insert("word")

    assert trie.search("word") is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
