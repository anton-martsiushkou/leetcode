import pytest
from coin_change import coin_change


def test_example_1():
    """Test case from example 1."""
    coins = [1, 2, 5]
    amount = 11
    result = coin_change(coins, amount)
    assert result == 3


def test_example_2():
    """Test case from example 2."""
    coins = [2]
    amount = 3
    result = coin_change(coins, amount)
    assert result == -1


def test_example_3():
    """Test case from example 3."""
    coins = [1]
    amount = 0
    result = coin_change(coins, amount)
    assert result == 0


def test_greedy_fails():
    """Test case where greedy algorithm fails."""
    coins = [1, 3, 4]
    amount = 6
    result = coin_change(coins, amount)
    assert result == 2


def test_large_amount():
    """Test with large amount."""
    coins = [1, 5, 10, 25]
    amount = 100
    result = coin_change(coins, amount)
    assert result == 4


def test_impossible_with_large_coins():
    """Test impossible case with large coins."""
    coins = [5, 10]
    amount = 3
    result = coin_change(coins, amount)
    assert result == -1


def test_single_coin_exact_match():
    """Test single coin exact match."""
    coins = [7]
    amount = 7
    result = coin_change(coins, amount)
    assert result == 1


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
