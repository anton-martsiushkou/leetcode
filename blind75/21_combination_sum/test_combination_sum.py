import pytest
from combination_sum import combination_sum


def test_example_1():
    """Test case from example 1."""
    candidates = [2, 3, 6, 7]
    target = 7
    result = combination_sum(candidates, target)
    expected = [[2, 2, 3], [7]]
    assert sorted(map(sorted, result)) == sorted(map(sorted, expected))


def test_example_2():
    """Test case from example 2."""
    candidates = [2, 3, 5]
    target = 8
    result = combination_sum(candidates, target)
    expected = [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    assert sorted(map(sorted, result)) == sorted(map(sorted, expected))


def test_example_3():
    """Test case from example 3."""
    candidates = [2]
    target = 1
    result = combination_sum(candidates, target)
    assert result == []


def test_single_candidate_exact_match():
    """Test with single candidate exact match."""
    candidates = [5]
    target = 5
    result = combination_sum(candidates, target)
    assert result == [[5]]


def test_single_candidate_multiple_uses():
    """Test with single candidate used multiple times."""
    candidates = [3]
    target = 9
    result = combination_sum(candidates, target)
    assert result == [[3, 3, 3]]


def test_multiple_combinations():
    """Test with multiple combinations."""
    candidates = [2, 3, 4]
    target = 6
    result = combination_sum(candidates, target)
    expected = [[2, 2, 2], [2, 4], [3, 3]]
    assert sorted(map(sorted, result)) == sorted(map(sorted, expected))


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
