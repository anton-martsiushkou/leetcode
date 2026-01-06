# Conditions

You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

 

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
 

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104

# Explanation

This is a classic "single pass" solution for finding the maximum profit from one buy and one sell transaction.
We iterate through the list of prices only once:

Maintain min_price: the lowest price encountered so far (the best day to buy up to now).
At each day, compute the potential profit if we sell today: price - min_price.
Update max_profit if this potential profit is higher than what we've seen before.
If today's price is lower than our current min_price, update min_price because buying later would be better.

This correctly finds the maximum difference where the buy day comes before the sell day.
Algorithms and Data Structures

## Algorithm: 
One-pass greedy scan (also known as Kadane's algorithm variant for maximum difference with order constraint).
Data Structures: Only uses two integer variables (min_price and max_profit). No additional data structures beyond the input list.

## Time and Memory Complexity

### Time Complexity: 
O(n), where n is the length of the prices array. We traverse the array exactly once.
### Space Complexity: 
O(1), constant extra space. We only use two variables regardless of input size.