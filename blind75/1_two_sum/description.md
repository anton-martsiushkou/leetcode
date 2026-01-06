# Conditions

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


# Explanation:

To solve the Two Sum problem, we iterate through the list of numbers once. For each number at index i, we calculate its complement (target - num) and check if it exists in a dictionary that maps numbers to their indices. If the complement is found, we return the indices of the complement and the current number. If not, we add the current number and its index to the dictionary. This ensures we find the pair in a single pass without revisiting elements.
Algorithms and Data Structures

## Algorithm: 
One-pass hash table approach. We use hashing to achieve constant-time lookups for complements.
Data Structures: Input is a list (nums). We use a dictionary (num_to_index) as a hash map to store number-index pairs for quick access.

## Time and Memory Complexity

### Time Complexity: 
O(n), where n is the length of the input list. We perform a single loop with constant-time operations (dictionary lookups and insertions) inside.
### Space Complexity: 
O(n), in the worst case, as we may store all elements in the dictionary before finding the pair.