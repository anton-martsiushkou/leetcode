import pytest
from number_of_islands import num_islands, num_islands_bfs
import copy


def test_example_1():
    """Test case from example 1 - single island."""
    grid = [
        ["1", "1", "1", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "0"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 1


def test_example_2():
    """Test case from example 2 - three islands."""
    grid = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 3


def test_no_islands():
    """Test with no islands."""
    grid = [
        ["0", "0", "0"],
        ["0", "0", "0"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 0


def test_all_islands():
    """Test with all cells being land."""
    grid = [
        ["1", "1"],
        ["1", "1"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 1


def test_single_cell_island():
    """Test with single cell island."""
    grid = [["1"]]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 1


def test_single_cell_water():
    """Test with single cell water."""
    grid = [["0"]]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 0


def test_diagonal_islands():
    """Test with diagonal islands (not connected)."""
    grid = [
        ["1", "0", "1"],
        ["0", "1", "0"],
        ["1", "0", "1"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands(grid_copy) == 5


def test_bfs_example_1():
    """Test BFS implementation with example 1."""
    grid = [
        ["1", "1", "1", "1", "0"],
        ["1", "1", "0", "1", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "0", "0", "0"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands_bfs(grid_copy) == 1


def test_bfs_example_2():
    """Test BFS implementation with example 2."""
    grid = [
        ["1", "1", "0", "0", "0"],
        ["1", "1", "0", "0", "0"],
        ["0", "0", "1", "0", "0"],
        ["0", "0", "0", "1", "1"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands_bfs(grid_copy) == 3


def test_bfs_diagonal():
    """Test BFS with diagonal islands."""
    grid = [
        ["1", "0", "1"],
        ["0", "1", "0"],
        ["1", "0", "1"]
    ]
    grid_copy = copy.deepcopy(grid)
    assert num_islands_bfs(grid_copy) == 5


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
