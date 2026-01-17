from typing import List


def max_profit(prices: List[int]) -> int:
    """
    Returns the maximum profit from buying and selling stock once.
    Uses single pass with O(n) time and O(1) space.

    Args:
        prices: List of stock prices where prices[i] is the price on day i

    Returns:
        Maximum profit achievable, or 0 if no profit is possible
    """
    if not prices:
        return 0

    min_price = prices[0]
    max_profit = 0

    for price in prices:
        if price < min_price:
            min_price = price
        elif price - min_price > max_profit:
            max_profit = price - min_price

    return max_profit
