from typing import List


def coin_change(coins: List[int], amount: int) -> int:
    """
    Returns the minimum number of coins needed to make up the amount.
    Returns -1 if the amount cannot be made up.

    Args:
        coins: List of coin denominations
        amount: Target amount to make

    Returns:
        Minimum number of coins needed, or -1 if impossible
    """
    # dp[i] represents minimum coins needed for amount i
    dp = [amount + 1] * (amount + 1)

    # Base case: 0 coins needed for amount 0
    dp[0] = 0

    # Build up solutions for each amount
    for i in range(1, amount + 1):
        for coin in coins:
            if coin <= i:
                dp[i] = min(dp[i], dp[i - coin] + 1)

    # If dp[amount] is still infinity, amount cannot be made
    return dp[amount] if dp[amount] <= amount else -1
