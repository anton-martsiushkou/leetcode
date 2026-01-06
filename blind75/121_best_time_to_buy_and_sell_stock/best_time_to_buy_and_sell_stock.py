from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if not prices:
            return 0

        min_price = prices[0]  # Track the lowest price seen so far
        max_profit = 0  # Track the maximum profit possible

        for price in prices:
            if price < min_price:
                min_price = price  # Update the minimum price if current is lower
            else:
                max_profit = max(max_profit, price - min_price)  # Update profit if selling today

        return max_profit
