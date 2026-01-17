# Best Time to Buy and Sell Stock

## Problem Description

You are given an array `prices` where `prices[i]` is the price of a given stock on the `i`th day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return `0`.

## Examples

**Example 1:**
```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

**Example 2:**
```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

## Constraints

- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Single Pass (Kadane's variant) | Variables | O(n) | O(1) |

### Approach: Track Minimum Price and Maximum Profit

The optimal solution uses a single pass through the array, keeping track of the minimum price seen so far and the maximum profit achievable.

**Key Insight:** For each price, the maximum profit we can achieve by selling at that price is `current_price - minimum_price_so_far`. We track the minimum price and update the maximum profit as we iterate.

### Algorithm Steps

1. Initialize `min_price` to infinity (or first element) and `max_profit` to 0
2. For each price in the array:
   - If `price < min_price`, update `min_price = price`
   - Calculate potential profit: `profit = price - min_price`
   - If `profit > max_profit`, update `max_profit = profit`
3. Return `max_profit`

### Example Walkthrough

For `prices = [7, 1, 5, 3, 6, 4]`:

1. **price=7**: min_price=7, profit=0, max_profit=0
2. **price=1**: min_price=1, profit=0, max_profit=0
3. **price=5**: min_price=1, profit=4, max_profit=4
4. **price=3**: min_price=1, profit=2, max_profit=4
5. **price=6**: min_price=1, profit=5, max_profit=5
6. **price=4**: min_price=1, profit=3, max_profit=5

Result: 5

### Why This is Optimal

- **Time Complexity O(n)**: Single pass through the array
- **Space Complexity O(1)**: Only two variables regardless of input size
- This is optimal because we must examine every price at least once to find the answer
