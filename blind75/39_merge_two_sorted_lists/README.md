# Merge Two Sorted Lists

## Problem Description

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

## Examples

**Example 1:**
```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```

**Example 2:**
```
Input: list1 = [], list2 = []
Output: []
```

**Example 3:**
```
Input: list1 = [], list2 = [0]
Output: [0]
```

## Constraints

- The number of nodes in both lists is in the range `[0, 50]`
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in non-decreasing order

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Iterative (Two Pointers) | Linked List | O(n + m) | O(1) |
| Recursive | Linked List | O(n + m) | O(n + m) |

### Approach 1: Iterative (Recommended)

Use a dummy node to simplify the logic and build the merged list by comparing nodes from both lists.

**Key Insight:** Since both lists are already sorted, we can build the merged list by always choosing the smaller node from the two lists. A dummy node helps us avoid special cases for the head.

### Algorithm Steps

1. Create a dummy node and a tail pointer pointing to it
2. While both `list1` and `list2` are not nil:
   - Compare the current nodes of both lists
   - Attach the smaller node to the tail
   - Move the pointer of the list from which we took the node
   - Move the tail pointer forward
3. After the loop, one list might still have remaining nodes
4. Attach the remaining nodes (if any) to the tail
5. Return `dummy.next` (the actual head of the merged list)

### Example Walkthrough

For `list1 = [1, 2, 4]` and `list2 = [1, 3, 4]`:

1. **Initial**: dummy -> null, tail = dummy
   - list1: 1->2->4
   - list2: 1->3->4

2. **Step 1**: Compare 1 and 1, choose list1
   - dummy -> 1, tail = 1
   - list1: 2->4, list2: 1->3->4

3. **Step 2**: Compare 2 and 1, choose list2
   - dummy -> 1 -> 1, tail = 1
   - list1: 2->4, list2: 3->4

4. **Step 3**: Compare 2 and 3, choose list1
   - dummy -> 1 -> 1 -> 2, tail = 2
   - list1: 4, list2: 3->4

5. **Step 4**: Compare 4 and 3, choose list2
   - dummy -> 1 -> 1 -> 2 -> 3, tail = 3
   - list1: 4, list2: 4

6. **Step 5**: Compare 4 and 4, choose list1
   - dummy -> 1 -> 1 -> 2 -> 3 -> 4, tail = 4
   - list1: nil, list2: 4

7. **Step 6**: list1 is nil, attach remaining list2
   - dummy -> 1 -> 1 -> 2 -> 3 -> 4 -> 4

8. **Return**: dummy.next = 1 (head of merged list)

### Why This is Optimal

- **Time Complexity O(n + m)**: We visit each node exactly once, where n and m are the lengths of the two lists
- **Space Complexity O(1)**: We only use a constant amount of extra space (dummy node and pointers)
- This is optimal because we must examine every node at least once to build the merged list

### Approach 2: Recursive

The recursive approach solves the problem by making a choice at each step: which node should be the next in the merged list?

**Algorithm:**
1. Base cases:
   - If `list1` is nil, return `list2`
   - If `list2` is nil, return `list1`
2. Recursive case:
   - If `list1.val <= list2.val`:
     - `list1.next = mergeTwoLists(list1.next, list2)`
     - Return `list1`
   - Else:
     - `list2.next = mergeTwoLists(list1, list2.next)`
     - Return `list2`

**Complexity:**
- Time: O(n + m)
- Space: O(n + m) due to call stack

The recursive solution is elegant but uses more space due to the call stack.
