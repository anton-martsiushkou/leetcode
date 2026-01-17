import pytest
from course_schedule import can_finish, can_finish_bfs


def test_example_1():
    """Test case from example 1 - possible."""
    num_courses = 2
    prerequisites = [[1, 0]]
    assert can_finish(num_courses, prerequisites) is True


def test_example_2():
    """Test case from example 2 - impossible cycle."""
    num_courses = 2
    prerequisites = [[1, 0], [0, 1]]
    assert can_finish(num_courses, prerequisites) is False


def test_no_prerequisites():
    """Test with no prerequisites."""
    num_courses = 3
    prerequisites = []
    assert can_finish(num_courses, prerequisites) is True


def test_linear_chain():
    """Test with linear chain of prerequisites."""
    num_courses = 4
    prerequisites = [[1, 0], [2, 1], [3, 2]]
    assert can_finish(num_courses, prerequisites) is True


def test_complex_valid_graph():
    """Test with complex valid graph."""
    num_courses = 4
    prerequisites = [[1, 0], [2, 0], [3, 1], [3, 2]]
    assert can_finish(num_courses, prerequisites) is True


def test_cycle_in_middle():
    """Test with cycle in middle of graph."""
    num_courses = 3
    prerequisites = [[0, 1], [1, 2], [2, 1]]
    assert can_finish(num_courses, prerequisites) is False


def test_self_loop():
    """Test with self loop."""
    num_courses = 1
    prerequisites = [[0, 0]]
    assert can_finish(num_courses, prerequisites) is False


def test_bfs_example_1():
    """Test BFS implementation with example 1."""
    num_courses = 2
    prerequisites = [[1, 0]]
    assert can_finish_bfs(num_courses, prerequisites) is True


def test_bfs_example_2():
    """Test BFS implementation with example 2."""
    num_courses = 2
    prerequisites = [[1, 0], [0, 1]]
    assert can_finish_bfs(num_courses, prerequisites) is False


def test_bfs_complex_valid():
    """Test BFS with complex valid graph."""
    num_courses = 4
    prerequisites = [[1, 0], [2, 0], [3, 1], [3, 2]]
    assert can_finish_bfs(num_courses, prerequisites) is True


def test_bfs_cycle():
    """Test BFS with cycle."""
    num_courses = 3
    prerequisites = [[0, 1], [1, 2], [2, 1]]
    assert can_finish_bfs(num_courses, prerequisites) is False


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
