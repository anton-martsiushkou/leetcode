# Combination Sum

## Problem Description

Given an array of distinct integers `candidates` and a target integer `target`, return a list of all unique combinations of `candidates` where the chosen numbers sum to `target`. You may return the combinations in any order.

The same number may be chosen from `candidates` an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to `target` is less than `150` combinations for the given input.

## Examples

**Example 1:**
```
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
```

**Example 2:**
```
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
```

**Example 3:**
```
Input: candidates = [2], target = 1
Output: []
```

## Constraints

- `1 <= candidates.length <= 30`
- `2 <= candidates[i] <= 40`
- All elements of `candidates` are distinct.
- `1 <= target <= 40`

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Backtracking with Pruning | Recursion Stack + Array | O(n^(target/min)) | O(target/min) |

where n = number of candidates, min = smallest candidate value

### Approach: Backtracking with Index Tracking

The optimal solution uses backtracking to explore all possible combinations. To avoid duplicates, we maintain an index that ensures we only use candidates at or after the current position. This prevents generating the same combination in different orders (e.g., [2,3] and [3,2]).

**Key Insight:** Instead of trying all candidates at each step (which creates duplicates), we maintain a start index. At each recursive call, we only consider candidates from this index onwards. This ensures combinations are built in non-decreasing order, eliminating duplicates.

### Algorithm Steps

1. **Sort Candidates (Optional):**
   - Sorting helps with early termination but not required for correctness

2. **Backtracking Function:**
   - Parameters: `start` (index to start from), `target` (remaining sum), `path` (current combination)
   - **Base Cases:**
     - If `target == 0`: add current path to results (valid combination found)
     - If `target < 0`: return (invalid path)
   - **Recursive Case:**
     - For each candidate from `start` to end:
       - Add candidate to path
       - Recurse with same start index (can reuse same number), reduced target
       - Remove candidate from path (backtrack)

3. **Return Results:**
   - Return all valid combinations found

### Example Walkthrough

For `candidates = [2, 3, 6, 7]` and `target = 7`:

```
Start with empty path, target=7, start=0

├─ Try 2: path=[2], target=5
│  ├─ Try 2: path=[2,2], target=3
│  │  ├─ Try 2: path=[2,2,2], target=1
│  │  │  └─ Try 2: path=[2,2,2,2], target=-1 ✗
│  │  ├─ Try 3: path=[2,2,3], target=0 ✓ ADD [2,2,3]
│  │  └─ Try 6,7: target < 0 ✗
│  ├─ Try 3: path=[2,3], target=2
│  │  └─ All tries fail
│  └─ Try 6,7: target < 0 ✗
│
├─ Try 3: path=[3], target=4
│  └─ All tries don't reach 0
│
├─ Try 6: path=[6], target=1
│  └─ All tries fail
│
└─ Try 7: path=[7], target=0 ✓ ADD [7]

Result: [[2,2,3], [7]]
```

### Why This is Optimal

- **Time Complexity O(n^(target/min))**: In the worst case, we explore all possible combinations. The depth of recursion is at most `target/min` and at each level we have up to `n` choices.
- **Space Complexity O(target/min)**: The recursion stack depth is at most `target/min`, and we store the current path which has the same maximum length.
- This is optimal because we must explore all valid combinations, and backtracking with pruning is the most efficient way to do this.

### Optimizations

1. **Sort Candidates:** Allows early termination when a candidate exceeds remaining target
2. **Start Index:** Prevents generating duplicate combinations in different orders
3. **Early Pruning:** Skip candidates larger than remaining target

### Alternative Approaches (Not Optimal)

1. **Generate All Subsets:** O(2^n × n) - Generates many invalid combinations
2. **Dynamic Programming:** Can find if target is reachable but harder to construct all actual combinations
3. **BFS:** Similar time complexity but more memory overhead for queue
