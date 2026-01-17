# Reorder List

## Problem Description

You are given the head of a singly linked-list. The list can be represented as:

```
L0 → L1 → … → Ln - 1 → Ln
```

Reorder the list to be in the following form:

```
L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
```

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

## Examples

**Example 1:**
```
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```

**Example 2:**
```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

## Constraints

- The number of nodes in the list is in the range `[1, 5 * 10^4]`
- `1 <= Node.val <= 1000`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Find Middle + Reverse + Merge | Linked List | O(n) | O(1) |
| Stack | Stack | O(n) | O(n) |

### Approach 1: Find Middle + Reverse + Merge (Recommended)

This approach combines three classic linked list operations:
1. Find the middle of the list (using slow/fast pointers)
2. Reverse the second half of the list
3. Merge the two halves alternately

**Key Insight:** The reordered list is essentially the first half merged alternately with the reversed second half.

### Algorithm Steps

1. **Find the middle of the list**:
   - Use slow and fast pointers
   - When fast reaches the end, slow will be at the middle

2. **Split and reverse the second half**:
   - Split the list at the middle
   - Reverse the second half

3. **Merge the two halves**:
   - Take nodes alternately from the first half and reversed second half
   - Link them together

### Example Walkthrough

For `head = [1, 2, 3, 4, 5]`:

1. **Find middle**:
   - slow starts at 1, fast at 1
   - After iterations: slow = 3 (middle)
   - First half: 1 -> 2 -> 3
   - Second half: 4 -> 5

2. **Reverse second half**:
   - Original: 4 -> 5
   - Reversed: 5 -> 4

3. **Merge**:
   - Take from first: 1
   - Take from second: 5
   - Take from first: 2
   - Take from second: 4
   - Take from first: 3
   - Result: 1 -> 5 -> 2 -> 4 -> 3

### Detailed Steps

```
Initial: 1 -> 2 -> 3 -> 4 -> 5

Step 1: Find middle
  slow: 1 -> 2 -> 3
  fast: 1 -> 3 -> 5 -> null
  Middle found at 3

Step 2: Split at middle
  first:  1 -> 2 -> 3
  second: 4 -> 5

Step 3: Reverse second half
  second: 5 -> 4

Step 4: Merge alternately
  1 -> 5 -> 2 -> 4 -> 3
```

### Edge Cases

1. **Single node**: Return as is
2. **Two nodes**: Swap them: [1,2] → [1,2] (no change needed)
3. **Odd length**: Middle node stays in the first half
4. **Even length**: Split evenly

### Why This is Optimal

- **Time Complexity O(n)**:
  - Finding middle: O(n)
  - Reversing second half: O(n/2)
  - Merging: O(n/2)
  - Total: O(n)
- **Space Complexity O(1)**: Only using a few pointers
- This is optimal because we must visit every node at least once

### Approach 2: Stack

Use a stack to store all nodes, then rebuild the list by alternately taking from the beginning and popping from the stack.

**Algorithm:**
1. Push all nodes onto a stack
2. Use two pointers: one at the start and one (via stack) at the end
3. Alternately link nodes from start and end
4. Stop when we've processed n/2 nodes

**Complexity:**
- Time: O(n)
- Space: O(n) - requires storing all nodes in the stack

This approach is simpler but uses extra space, so the first approach is preferred.

### Implementation Notes

- The in-place approach modifies the original list structure
- No new nodes are created; we only rearrange pointers
- The middle finding uses Floyd's slow/fast pointer technique
- The reversal uses the standard iterative reversal technique
