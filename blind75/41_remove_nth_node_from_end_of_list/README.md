# Remove Nth Node From End of List

## Problem Description

Given the `head` of a linked list, remove the `nth` node from the end of the list and return its head.

## Examples

**Example 1:**
```
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Explanation: Remove the 2nd node from the end (node with value 4)
```

**Example 2:**
```
Input: head = [1], n = 1
Output: []
Explanation: Remove the only node
```

**Example 3:**
```
Input: head = [1,2], n = 1
Output: [1]
Explanation: Remove the last node
```

## Constraints

- The number of nodes in the list is `sz`
- `1 <= sz <= 30`
- `0 <= Node.val <= 100`
- `1 <= n <= sz`

## Follow-up

Could you do this in one pass?

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Two Pointers (One Pass) | Linked List | O(n) | O(1) |
| Two Pass | Linked List | O(n) | O(1) |

### Approach 1: Two Pointers - One Pass (Recommended)

Use two pointers with a gap of n nodes between them. When the fast pointer reaches the end, the slow pointer will be at the node before the one to be removed.

**Key Insight:** If we maintain two pointers that are exactly n nodes apart, when the fast pointer reaches the end, the slow pointer will be pointing to the node just before the one we need to remove.

### Algorithm Steps

1. Create a dummy node pointing to head (handles edge case of removing the first node)
2. Initialize two pointers: `fast` and `slow`, both pointing to dummy
3. Move `fast` n+1 steps ahead (to create a gap of n nodes)
4. Move both `fast` and `slow` one step at a time until `fast` reaches the end
5. Now `slow` is pointing to the node before the one to remove
6. Remove the nth node: `slow.next = slow.next.next`
7. Return `dummy.next`

### Example Walkthrough

For `head = [1, 2, 3, 4, 5]` and `n = 2`:

1. **Initial**: Create dummy -> 1 -> 2 -> 3 -> 4 -> 5
   - fast = dummy, slow = dummy

2. **Move fast n+1 times (3 times)**:
   - fast moves: dummy -> 1 -> 2 -> 3
   - slow stays at dummy

3. **Move both until fast reaches end**:
   - Step 1: fast = 4, slow = 1
   - Step 2: fast = 5, slow = 2
   - Step 3: fast = nil, slow = 3

4. **Remove node**: slow.next = slow.next.next
   - 3.next = 5 (removing node 4)

5. **Result**: [1, 2, 3, 5]

### Edge Cases

1. **Removing the first node** (n = length):
   - The dummy node handles this gracefully
   - After moving fast n+1 times, slow stays at dummy
   - slow.next = slow.next.next removes the first node

2. **Removing the only node**:
   - dummy.next becomes nil

3. **Removing the last node** (n = 1):
   - Standard case, works as described above

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the list only once
- **Space Complexity O(1)**: We only use two pointers regardless of list size
- Satisfies the follow-up requirement of one pass
- The dummy node eliminates the need for special case handling

### Approach 2: Two Pass

First pass: count the total number of nodes. Second pass: remove the (length - n)th node from the beginning.

**Algorithm:**
1. Count the total nodes in the list: `length`
2. Create a dummy node
3. Move to the `(length - n)`th node
4. Remove the next node
5. Return dummy.next

**Complexity:**
- Time: O(n) - Two passes through the list
- Space: O(1)

This approach is simpler to understand but requires two passes, which doesn't satisfy the follow-up.

### Comparison

| Approach | Passes | Complexity | Code Complexity |
|----------|--------|------------|-----------------|
| Two Pointers | 1 | O(n) time, O(1) space | Moderate |
| Two Pass | 2 | O(n) time, O(1) space | Simple |

The two-pointer approach is preferred as it satisfies the one-pass requirement and is still reasonably simple to implement.
