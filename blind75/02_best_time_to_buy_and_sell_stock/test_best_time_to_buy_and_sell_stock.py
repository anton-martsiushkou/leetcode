import pytest
from best_time_to_buy_and_sell_stock import max_profit


def test_example_1():
    """Test case from example 1."""
    prices = [7, 1, 5, 3, 6, 4]
    assert max_profit(prices) == 5


def test_example_2_descending_prices():
    """Test with descending prices - no profit possible."""
    prices = [7, 6, 4, 3, 1]
    assert max_profit(prices) == 0


def test_ascending_prices():
    """Test with ascending prices - max profit is last - first."""
    prices = [1, 2, 3, 4, 5]
    assert max_profit(prices) == 4


def test_single_price():
    """Test with single price - no transaction possible."""
    prices = [5]
    assert max_profit(prices) == 0


def test_two_prices_profit():
    """Test with two prices where profit is possible."""
    prices = [1, 5]
    assert max_profit(prices) == 4


def test_two_prices_no_profit():
    """Test with two prices where no profit is possible."""
    prices = [5, 1]
    assert max_profit(prices) == 0


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
