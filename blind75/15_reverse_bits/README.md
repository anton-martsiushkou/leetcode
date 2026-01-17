# Reverse Bits

## Problem Description

Reverse bits of a given 32 bits unsigned integer.

## Examples

**Example 1:**
```
Input: n = 00000010100101000001111010011100
Output: 964176192 (00111001011110000010100101000000)
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596,
so return 964176192 which its binary representation is 00111001011110000010100101000000.
```

**Example 2:**
```
Input: n = 11111111111111111111111111111101
Output: 3221225471 (10111111111111111111111111111111)
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293,
so return 3221225471 which its binary representation is 10111111111111111111111111111111.
```

## Constraints

- The input must be a binary string of length 32

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Bit Manipulation | N/A | O(1) | O(1) |

### Approach: Bit-by-Bit Reversal

The optimal solution processes each bit from right to left in the input and builds the result from left to right. We extract each bit from the input, shift the result left to make room, and add the extracted bit.

**Key Insight:** We can reverse bits by iterating through all 32 bits, extracting each bit from the input (rightmost first), and placing it into the result (leftmost first). For each iteration:
1. Extract the rightmost bit of input: `n & 1`
2. Shift result left to make room: `result << 1`
3. Add the extracted bit to result: `result | (n & 1)`
4. Shift input right for next iteration: `n >> 1`

### Algorithm Steps

1. Initialize `result = 0`
2. Loop 32 times (for each bit):
   - Shift result left by 1: `result <<= 1`
   - Extract the rightmost bit of n and add to result: `result |= (n & 1)`
   - Shift n right by 1: `n >>= 1`
3. Return `result`

### Example Walkthrough

For `n = 43261596` (binary: `00000010100101000001111010011100`):

We process bits from right to left:

```
Initial: result = 0, n = 00000010100101000001111010011100

Bit 0: rightmost bit = 0
  result = 0 << 1 = 0
  result = 0 | 0 = 0
  n = n >> 1

Bit 1: rightmost bit = 0
  result = 0 << 1 = 0
  result = 0 | 0 = 0
  n = n >> 1

Bit 2: rightmost bit = 1
  result = 0 << 1 = 0
  result = 0 | 1 = 1
  n = n >> 1

... (continue for all 32 bits) ...

Final result = 964176192 (binary: 00111001011110000010100101000000)
```

### Bit Manipulation Operations Used

- **`n & 1`**: Extracts the rightmost bit (gives 0 or 1)
- **`result << 1`**: Shifts all bits left by one position (makes room for new bit)
- **`result | bit`**: Adds the bit to the rightmost position
- **`n >> 1`**: Shifts all bits right by one position (moves to next bit)

### Why This is Optimal

- **Time Complexity O(1)**: We always iterate exactly 32 times, which is constant
- **Space Complexity O(1)**: We only use a constant amount of extra space
- This is optimal because we must examine every bit to reverse them

### Alternative Approaches

1. **Divide and Conquer (Parallel Bit Reversal)**:
   - Swap adjacent bits, then swap pairs, then nibbles, etc.
   - More complex but can be faster in practice
   - Example: swap all odd/even bits, then swap 2-bit chunks, etc.

2. **Lookup Table**:
   - Pre-compute reversals for smaller chunks (e.g., 8-bit)
   - Reverse byte-by-byte using lookup
   - Time: O(1), Space: O(256) for 8-bit lookup table

The bit-by-bit approach is preferred for clarity and minimal space usage.
