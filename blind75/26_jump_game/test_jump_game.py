import pytest
from jump_game import can_jump


def test_example_1():
    """Test case from example 1."""
    nums = [2, 3, 1, 1, 4]
    result = can_jump(nums)
    assert result is True


def test_example_2():
    """Test case from example 2."""
    nums = [3, 2, 1, 0, 4]
    result = can_jump(nums)
    assert result is False


def test_single_element():
    """Test with a single element."""
    nums = [0]
    result = can_jump(nums)
    assert result is True


def test_two_elements_can_reach():
    """Test with two elements where we can reach."""
    nums = [1, 0]
    result = can_jump(nums)
    assert result is True


def test_two_elements_cannot_reach():
    """Test with two elements where we cannot reach."""
    nums = [0, 1]
    result = can_jump(nums)
    assert result is False


def test_all_zeros_except_first():
    """Test with all zeros except first element."""
    nums = [5, 0, 0, 0, 0, 0]
    result = can_jump(nums)
    assert result is True


def test_large_jump_at_start():
    """Test with large jump at start."""
    nums = [10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
    result = can_jump(nums)
    assert result is True


def test_can_barely_reach():
    """Test where we can barely reach the end."""
    nums = [1, 1, 1, 1, 1]
    result = can_jump(nums)
    assert result is True


def test_stuck_in_middle():
    """Test where we get stuck in the middle."""
    nums = [1, 1, 0, 1]
    result = can_jump(nums)
    assert result is False


def test_all_ones():
    """Test with all ones."""
    nums = [1, 1, 1, 1]
    result = can_jump(nums)
    assert result is True


def test_decreasing_jumps():
    """Test with decreasing jump values."""
    nums = [5, 4, 3, 2, 1, 0]
    result = can_jump(nums)
    assert result is True


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
