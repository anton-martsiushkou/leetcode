import pytest
from word_search import exist


def test_example_1():
    """Test case from example 1 - word exists."""
    board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    word = "ABCCED"
    assert exist(board, word) is True


def test_example_2():
    """Test case from example 2 - word exists."""
    board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    word = "SEE"
    assert exist(board, word) is True


def test_example_3():
    """Test case from example 3 - word does not exist."""
    board = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    word = "ABCB"
    assert exist(board, word) is False


def test_single_cell_match():
    """Test with single cell that matches."""
    board = [["A"]]
    word = "A"
    assert exist(board, word) is True


def test_single_cell_no_match():
    """Test with single cell that doesn't match."""
    board = [["A"]]
    word = "B"
    assert exist(board, word) is False


def test_word_longer_than_board():
    """Test with word longer than board."""
    board = [["A", "B"]]
    word = "ABCD"
    assert exist(board, word) is False


def test_requires_backtracking():
    """Test that requires backtracking to find solution."""
    board = [["C", "A", "A"], ["A", "A", "A"], ["B", "C", "D"]]
    word = "AAB"
    assert exist(board, word) is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
