# Product of Array Except Self

## Problem Description

Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operator.

## Examples

**Example 1:**
```
Input: nums = [1,2,3,4]
Output: [24,12,8,6]
```

**Example 2:**
```
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
```

## Constraints

- `2 <= nums.length <= 10^5`
- `-30 <= nums[i] <= 30`
- The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer.

**Follow up:** Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Prefix-Suffix Product | Array | O(n) | O(1)* |

*O(1) if we don't count the output array, otherwise O(n)

### Approach: Prefix-Suffix Product

The optimal solution uses the observation that for each element at index `i`, the product of all other elements equals:
- Product of all elements to the left of `i` (prefix product)
- TIMES
- Product of all elements to the right of `i` (suffix product)

Instead of creating separate prefix and suffix arrays, we can build the result array in-place using two passes:
1. First pass (left to right): Fill the result array with prefix products
2. Second pass (right to left): Multiply each element by the suffix product

**Key Insight:** We can avoid using extra space by storing prefix products in the output array, then multiply them by suffix products in a second pass using a single variable to track the running suffix product.

### Algorithm Steps

1. Create result array of length n, initialized with 1s
2. **First pass (prefix products - left to right):**
   - Keep a running prefix product (initially 1)
   - For each index i from 0 to n-1:
     - Store current prefix product in result[i]
     - Update prefix product by multiplying with nums[i]
3. **Second pass (suffix products - right to left):**
   - Keep a running suffix product (initially 1)
   - For each index i from n-1 to 0:
     - Multiply result[i] by current suffix product
     - Update suffix product by multiplying with nums[i]
4. Return result array

### Example Walkthrough

For `nums = [1, 2, 3, 4]`:

**First pass (prefix products):**
- i=0: result[0] = 1 (no elements before), prefix = 1 * 1 = 1
- i=1: result[1] = 1 (product of elements before: 1), prefix = 1 * 2 = 2
- i=2: result[2] = 2 (product of elements before: 1*2), prefix = 2 * 3 = 6
- i=3: result[3] = 6 (product of elements before: 1*2*3), prefix = 6 * 4 = 24

After first pass: `result = [1, 1, 2, 6]`

**Second pass (suffix products):**
- i=3: result[3] = 6 * 1 = 6 (no elements after), suffix = 1 * 4 = 4
- i=2: result[2] = 2 * 4 = 8 (product of elements after: 4), suffix = 4 * 3 = 12
- i=1: result[1] = 1 * 12 = 12 (product of elements after: 3*4), suffix = 12 * 2 = 24
- i=0: result[0] = 1 * 24 = 24 (product of elements after: 2*3*4), suffix = 24 * 1 = 24

Final result: `[24, 12, 8, 6]`

### Why This is Optimal

- **Time Complexity O(n)**: We make exactly two passes through the array
- **Space Complexity O(1)**: We only use a constant amount of extra space (two variables for prefix and suffix). The output array doesn't count toward space complexity
- This is optimal because we must examine every element at least once, giving us a lower bound of O(n)
- We avoid the division operator and handle zeros naturally
