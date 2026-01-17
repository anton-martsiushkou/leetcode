# Detect Cycle in Linked List

## Problem Description

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. Note that `pos` is not passed as a parameter.

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

## Examples

**Example 1:**
```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

**Example 2:**
```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
```

**Example 3:**
```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```

## Constraints

- The number of nodes in the list is in the range `[0, 10^4]`
- `-10^5 <= Node.val <= 10^5`
- `pos` is `-1` or a valid index in the linked list

## Follow-up

Can you solve it using O(1) memory?

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Floyd's Cycle Detection (Tortoise and Hare) | Linked List | O(n) | O(1) |
| Hash Set | Hash Set | O(n) | O(n) |

### Approach 1: Floyd's Cycle Detection Algorithm (Recommended)

This is also known as the "Tortoise and Hare" algorithm. We use two pointers moving at different speeds. If there's a cycle, the fast pointer will eventually catch up to the slow pointer.

**Key Insight:** If there's a cycle, two pointers moving at different speeds will eventually meet. If there's no cycle, the fast pointer will reach the end of the list.

### Algorithm Steps

1. Initialize two pointers:
   - `slow = head` (moves one step at a time)
   - `fast = head` (moves two steps at a time)
2. While `fast` and `fast.next` are not nil:
   - Move `slow` one step: `slow = slow.next`
   - Move `fast` two steps: `fast = fast.next.next`
   - If `slow == fast`, a cycle exists, return `true`
3. If the loop ends (fast reached the end), return `false`

### Example Walkthrough

For `head = [3, 2, 0, -4]` with cycle at position 1:

1. **Initial**: slow=3, fast=3
2. **Step 1**: slow=2, fast=0
3. **Step 2**: slow=0, fast=2
4. **Step 3**: slow=-4, fast=-4
5. **Cycle detected**: slow == fast, return true

### Why This is Optimal

- **Time Complexity O(n)**: In the worst case, we visit each node at most twice
  - If no cycle: fast pointer reaches the end in n/2 steps
  - If cycle exists: slow pointer enters the cycle, and fast pointer catches up within the cycle length
- **Space Complexity O(1)**: We only use two pointers regardless of list size
- This satisfies the follow-up requirement for O(1) space

### Approach 2: Hash Set

Use a hash set to store visited nodes. If we encounter a node we've already seen, there's a cycle.

**Algorithm:**
1. Create an empty hash set
2. Traverse the list:
   - If current node is in the set, return `true`
   - Add current node to the set
   - Move to next node
3. If we reach the end, return `false`

**Complexity:**
- Time: O(n)
- Space: O(n) - requires storing all nodes

This approach is simpler to understand but doesn't meet the O(1) space requirement.

### Mathematical Proof (Floyd's Algorithm)

If there's a cycle of length C:
- When slow enters the cycle, fast is already somewhere in the cycle
- The distance between them decreases by 1 each step (fast gains 1 position)
- They will meet within C steps after slow enters the cycle
- Total time: O(n) where n is the number of nodes before the cycle + C
