# Number of 1 Bits

## Problem Description

Write a function that takes the binary representation of a positive integer and returns the number of set bits it has (also known as the Hamming weight).

## Examples

**Example 1:**
```
Input: n = 11
Output: 3
Explanation: The input binary string 1011 has three set bits.
```

**Example 2:**
```
Input: n = 128
Output: 1
Explanation: The input binary string 10000000 has one set bit.
```

**Example 3:**
```
Input: n = 2147483645
Output: 30
Explanation: The input binary string 1111111111111111111111111111101 has thirty set bits.
```

## Constraints

- `1 <= n <= 2^31 - 1`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Brian Kernighan's Algorithm | N/A | O(k) where k = number of 1 bits | O(1) |

### Approach: Brian Kernighan's Algorithm

The optimal solution uses Brian Kernighan's algorithm, which repeatedly clears the rightmost set bit. This is more efficient than checking every bit because it only iterates once per set bit, not once per total bit.

**Key Insight:** The operation `n & (n - 1)` always clears the rightmost set bit in `n`. By repeatedly applying this operation until `n` becomes 0, we count exactly how many set bits there were.

### Algorithm Steps

1. Initialize a counter to 0
2. While `n` is not zero:
   - Apply `n = n & (n - 1)` to clear the rightmost set bit
   - Increment the counter
3. Return the counter

### How `n & (n - 1)` Works

When you subtract 1 from a number:
- All bits after the rightmost set bit get flipped
- The rightmost set bit becomes 0
- All bits before remain the same

Example: `n = 12` (binary: 1100)
- `n - 1 = 11` (binary: 1011)
- `n & (n - 1) = 1100 & 1011 = 1000` (8)
- The rightmost 1 bit was cleared!

### Example Walkthrough

For `n = 11` (binary: 1011):

1. **Iteration 1**: n = 1011 (11)
   - `n & (n - 1) = 1011 & 1010 = 1010` (10)
   - count = 1

2. **Iteration 2**: n = 1010 (10)
   - `n & (n - 1) = 1010 & 1001 = 1000` (8)
   - count = 2

3. **Iteration 3**: n = 1000 (8)
   - `n & (n - 1) = 1000 & 0111 = 0000` (0)
   - count = 3

4. **Result**: Return 3

### Why This is Optimal

- **Time Complexity O(k)**: Where k is the number of set bits. In the worst case (all bits set), k = 32 for a 32-bit integer, which is still O(1) constant time.
- **Space Complexity O(1)**: We only use a constant amount of extra space.
- This is better than the naive approach of checking all 32 bits one by one, especially when there are few set bits.

### Alternative Approaches

1. **Loop and Check Each Bit**: O(32) = O(1) time, but always checks all bits even if only one is set
2. **Built-in Population Count**: Some languages have built-in functions (e.g., `__builtin_popcount` in C++), but Brian Kernighan's algorithm is the fundamental implementation
3. **Lookup Table**: O(1) time but requires O(2^n) space for an n-bit lookup table
