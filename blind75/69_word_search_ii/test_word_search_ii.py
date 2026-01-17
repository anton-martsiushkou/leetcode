import pytest
from word_search_ii import find_words


def test_example_1():
    """Test case from example 1."""
    board = [
        ["o", "a", "a", "n"],
        ["e", "t", "a", "e"],
        ["i", "h", "k", "r"],
        ["i", "f", "l", "v"],
    ]
    words = ["oath", "pea", "eat", "rain"]
    result = find_words(board, words)
    assert sorted(result) == sorted(["eat", "oath"])


def test_example_2():
    """Test case from example 2."""
    board = [["a", "b"], ["c", "d"]]
    words = ["abcb"]
    result = find_words(board, words)
    assert result == []


def test_single_cell_board():
    """Test with single cell board."""
    board = [["a"]]
    words = ["a", "b"]
    result = find_words(board, words)
    assert result == ["a"]


def test_no_matches():
    """Test when no words match."""
    board = [["a", "b"], ["c", "d"]]
    words = ["xyz", "pqr"]
    result = find_words(board, words)
    assert result == []


def test_all_words_match():
    """Test when all words match."""
    board = [["a", "b"], ["c", "d"]]
    words = ["ab", "cd", "ac", "bd"]
    result = find_words(board, words)
    assert sorted(result) == sorted(["ab", "cd", "ac", "bd"])


def test_overlapping_paths():
    """Test with overlapping paths."""
    board = [["a", "a"], ["a", "a"]]
    words = ["aa", "aaa", "aaaa"]
    result = find_words(board, words)
    assert sorted(result) == sorted(["aa", "aaa", "aaaa"])


def test_longer_word():
    """Test with longer words."""
    board = [["a", "b", "c"], ["d", "e", "f"], ["g", "h", "i"]]
    words = ["abef", "abeh", "abehi"]
    result = find_words(board, words)
    assert sorted(result) == sorted(["abef", "abehi"])


def test_empty_board():
    """Test with empty board."""
    board = []
    words = ["test"]
    result = find_words(board, words)
    assert result == []


def test_single_word_found():
    """Test finding a single word."""
    board = [["a", "b"], ["c", "d"]]
    words = ["abdc"]
    result = find_words(board, words)
    assert result == ["abdc"]


def test_diagonal_not_allowed():
    """Test that diagonal moves are not allowed."""
    board = [["a", "b"], ["c", "d"]]
    words = ["ad", "bc"]
    result = find_words(board, words)
    # "ad" and "bc" would require diagonal moves, so should not be found
    assert result == []


def test_complex_board():
    """Test with a more complex board."""
    board = [
        ["a", "b", "c", "e"],
        ["s", "f", "c", "s"],
        ["a", "d", "e", "e"],
    ]
    words = ["abcced", "see", "abcb", "abfsadee"]
    result = find_words(board, words)
    assert sorted(result) == sorted(["abcced", "see", "abfsadee"])


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
