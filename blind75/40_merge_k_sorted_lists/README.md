# Merge K Sorted Lists

## Problem Description

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

## Examples

**Example 1:**
```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```

**Example 2:**
```
Input: lists = []
Output: []
```

**Example 3:**
```
Input: lists = [[]]
Output: []
```

## Constraints

- `k == lists.length`
- `0 <= k <= 10^4`
- `0 <= lists[i].length <= 500`
- `-10^4 <= lists[i][j] <= 10^4`
- `lists[i]` is sorted in ascending order
- The sum of `lists[i].length` will not exceed `10^4`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Min Heap (Priority Queue) | Min Heap | O(N log k) | O(k) |
| Divide and Conquer | Linked List | O(N log k) | O(log k) |
| Brute Force (Merge One by One) | Linked List | O(N * k) | O(1) |

where N is the total number of nodes across all lists and k is the number of lists.

### Approach 1: Min Heap / Priority Queue (Recommended)

Use a min heap to efficiently find the smallest node among the k lists at each step.

**Key Insight:** At each step, we need to find the minimum among k candidates (the current head of each list). A min heap can do this in O(log k) time, which is much better than scanning all k lists linearly.

### Algorithm Steps

1. Create a min heap and insert the first node from each non-empty list
2. Create a dummy node to build the result list
3. While the heap is not empty:
   - Pop the smallest node from the heap
   - Add it to the result list
   - If the popped node has a next node, push it to the heap
4. Return the result list

### Example Walkthrough

For `lists = [[1,4,5], [1,3,4], [2,6]]`:

1. **Initial Heap**: [1(list0), 1(list1), 2(list2)]
   - Result: []

2. **Step 1**: Pop 1(list0), add to result, push 4(list0)
   - Heap: [1(list1), 2(list2), 4(list0)]
   - Result: 1

3. **Step 2**: Pop 1(list1), add to result, push 3(list1)
   - Heap: [2(list2), 3(list1), 4(list0)]
   - Result: 1->1

4. **Step 3**: Pop 2(list2), add to result, push 6(list2)
   - Heap: [3(list1), 4(list0), 6(list2)]
   - Result: 1->1->2

5. **Continue**: Pop 3, 4(list0), 4(list1), 5, 6
   - Final Result: 1->1->2->3->4->4->5->6

### Why This is Optimal

- **Time Complexity O(N log k)**:
  - N total nodes across all lists
  - Each node is inserted and removed from heap once: O(log k) per operation
  - Total: O(N log k)
- **Space Complexity O(k)**: The heap contains at most k nodes (one from each list)
- This is optimal for the comparison-based approach

### Approach 2: Divide and Conquer

Pair up lists and merge them, then repeat the process until only one list remains.

**Algorithm:**
1. If lists is empty, return nil
2. While there's more than one list:
   - Pair up consecutive lists and merge each pair
   - Store the merged lists
   - Repeat with the new list of merged lists
3. Return the final merged list

**Complexity:**
- Time: O(N log k)
  - log k levels of merging
  - Each level processes all N nodes
- Space: O(log k) for recursion stack (or O(1) if iterative)

This approach has the same time complexity as the heap approach but may be simpler to implement.

### Approach 3: Brute Force (Not Optimal)

Merge lists one by one: merge list[0] with list[1], then merge result with list[2], etc.

**Complexity:**
- Time: O(N * k) - Each merge operation gets progressively longer
- Space: O(1)

This is inefficient because we keep reprocessing the same nodes multiple times.

### Comparison

| Approach | Time | Space | Pros | Cons |
|----------|------|-------|------|------|
| Min Heap | O(N log k) | O(k) | Efficient, intuitive | Requires heap implementation |
| Divide & Conquer | O(N log k) | O(log k) or O(1) | Simple, space-efficient | More code |
| Brute Force | O(N * k) | O(1) | Simplest | Too slow for large k |
