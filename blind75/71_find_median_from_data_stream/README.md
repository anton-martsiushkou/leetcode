# Find Median from Data Stream

## Problem Description

The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value and the median is the mean of the two middle values.

Implement the MedianFinder class:

- `MedianFinder()` initializes the MedianFinder object.
- `void addNum(int num)` adds the integer `num` from the data stream to the data structure.
- `double findMedian()` returns the median of all elements so far.

## Examples

**Example 1:**
```
Input:
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]

Output:
[null, null, null, 1.5, null, 2.0]

Explanation:
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr = [1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

## Constraints

- `-10^5 <= num <= 10^5`
- There will be at least one element in the data structure before calling `findMedian`.
- At most `5 * 10^4` calls will be made to `addNum` and `findMedian`.

## Solution

| Algorithm | Data Structure | Time Complexity (add/find) | Space Complexity |
|-----------|---------------|----------------------------|------------------|
| Two Heaps | Max Heap + Min Heap | O(log n) / O(1) | O(n) |

### Approach: Two Heaps

The optimal solution uses two heaps to maintain the median efficiently:
- A max heap for the smaller half of numbers
- A min heap for the larger half of numbers

**Key Insight:** By maintaining two balanced heaps, the median is always at the top of one or both heaps, allowing O(1) median retrieval.

### Algorithm Steps

1. **Initialization**: Create a max heap (for smaller half) and a min heap (for larger half)

2. **Add Number**:
   - Always add to max heap first
   - Move the largest from max heap to min heap
   - If max heap has fewer elements than min heap, move the smallest from min heap back to max heap
   - This ensures max heap has equal or one more element than min heap

3. **Find Median**:
   - If max heap has more elements: return its top
   - Otherwise: return average of both heap tops

### Example Walkthrough

For sequence `[1, 2, 3]`:

1. **addNum(1)**:
   - Max heap: [1], Min heap: []

2. **addNum(2)**:
   - Add to max: [1, 2]
   - Balance: Max [1], Min [2]
   - Median: (1 + 2) / 2 = 1.5

3. **addNum(3)**:
   - Add to max: [1, 3], Min [2]
   - Balance: Max [1, 2], Min [3]
   - Median: 2.0 (top of max heap)

### Why This is Optimal

- **Time Complexity**: O(log n) for addNum (heap operations), O(1) for findMedian
- **Space Complexity**: O(n) to store all numbers
- This is optimal because we need to maintain order for median calculation, and heaps provide the best balance between insertion and retrieval

### Alternative Approaches (Not Optimal)

1. **Sort on Each Query**: O(n log n) per findMedian - Too slow
2. **Insertion Sort**: O(n) per addNum - Too slow for many insertions
3. **Balanced BST**: O(log n) for both operations but more complex to implement
