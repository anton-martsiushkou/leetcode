# Sum of Two Integers

## Problem Description

Given two integers `a` and `b`, return the sum of the two integers without using the operators `+` and `-`.

## Examples

**Example 1:**
```
Input: a = 1, b = 2
Output: 3
```

**Example 2:**
```
Input: a = 2, b = 3
Output: 5
```

**Example 3:**
```
Input: a = -1, b = 1
Output: 0
```

## Constraints

- `-1000 <= a, b <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Bit Manipulation | N/A | O(1) | O(1) |

### Approach: Bit Manipulation

The optimal solution uses bitwise operations to simulate addition. The key insight is that addition can be broken down into two operations:
1. XOR gives us the sum without considering carry
2. AND followed by left shift gives us the carry

**Key Insight:** We repeatedly calculate the sum (using XOR) and the carry (using AND and left shift) until there's no carry left.

### Algorithm Steps

1. While `b` is not zero:
   - Calculate sum without carry: `sum = a ^ b`
   - Calculate carry: `carry = (a & b) << 1`
   - Update `a = sum`
   - Update `b = carry`
2. Return `a` (which now contains the final sum)

### How Bit Addition Works

- **XOR (^)**: Adds bits without carry
  - `1 ^ 0 = 1`, `0 ^ 1 = 1`, `0 ^ 0 = 0`, `1 ^ 1 = 0`
- **AND (&)**: Finds where carry occurs
  - `1 & 1 = 1` (carry needed), all others = 0
- **Left Shift (<<)**: Moves carry to the next position

### Example Walkthrough

For `a = 5` (101 in binary) and `b = 3` (011 in binary):

1. **Iteration 1**:
   - Sum without carry: `101 ^ 011 = 110` (6)
   - Carry: `(101 & 011) << 1 = 001 << 1 = 010` (2)
   - a = 6, b = 2

2. **Iteration 2**:
   - Sum without carry: `110 ^ 010 = 100` (4)
   - Carry: `(110 & 010) << 1 = 010 << 1 = 100` (4)
   - a = 4, b = 4

3. **Iteration 3**:
   - Sum without carry: `100 ^ 100 = 000` (0)
   - Carry: `(100 & 100) << 1 = 100 << 1 = 1000` (8)
   - a = 0, b = 8

4. **Iteration 4**:
   - Sum without carry: `000 ^ 1000 = 1000` (8)
   - Carry: `(000 & 1000) << 1 = 000 << 1 = 0000` (0)
   - a = 8, b = 0

5. **Result**: Return 8

### Why This is Optimal

- **Time Complexity O(1)**: The number of iterations is bounded by the number of bits in the integer (32 bits for a standard int). Since this is constant, the time complexity is O(1).
- **Space Complexity O(1)**: We only use a constant amount of extra space for variables.
- This is optimal because we're using the fundamental way computers perform addition at the hardware level.

### Edge Cases

- **Negative numbers**: The algorithm works correctly with negative numbers due to two's complement representation
- **Zero**: If either number is zero, the XOR will return the other number and carry will be zero
- **Overflow**: Not a concern within the given constraints, but in practice, bit operations handle overflow naturally
