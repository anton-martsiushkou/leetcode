# Valid Parentheses

## Problem Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Examples

**Example 1:**
```
Input: s = "()"
Output: true
```

**Example 2:**
```
Input: s = "()[]{}"
Output: true
```

**Example 3:**
```
Input: s = "(]"
Output: false
```

**Example 4:**
```
Input: s = "([])"
Output: true
```

## Constraints

- `1 <= s.length <= 10^4`
- `s` consists of parentheses only `'()[]{}'`.

## Solution

| Algorithm | Data Structure | Time Complexity | Space Complexity |
|-----------|---------------|-----------------|------------------|
| Stack | Stack | O(n) | O(n) |

### Approach: Stack

The optimal solution uses a stack to track opening brackets. When we encounter a closing bracket, we check if it matches the most recent opening bracket (top of stack).

**Key Insight:** Valid parentheses follow a Last-In-First-Out (LIFO) pattern - the most recently opened bracket must be closed first. This is exactly what a stack provides.

### Algorithm Steps

1. Create an empty stack
2. Create a mapping of closing brackets to opening brackets: `{')': '(', '}': '{', ']': '['}`
3. Iterate through each character in the string:
   - If it's an opening bracket `(`, `{`, or `[`:
     - Push it onto the stack
   - If it's a closing bracket `)`, `}`, or `]`:
     - If stack is empty, return false (no matching opening bracket)
     - Pop from stack and check if it matches the closing bracket
     - If it doesn't match, return false
4. After processing all characters, return true only if stack is empty (all brackets matched)

### Example Walkthrough

For `s = "([])"``:

1. **char='('**: Opening bracket, push to stack. Stack: ['(']
2. **char='['**: Opening bracket, push to stack. Stack: ['(', '[']
3. **char=']'**: Closing bracket, pop '[' from stack, matches. Stack: ['(']
4. **char=')'**: Closing bracket, pop '(' from stack, matches. Stack: []
5. Stack is empty, return true

For `s = "(]"`:

1. **char='('**: Opening bracket, push to stack. Stack: ['(']
2. **char=']'**: Closing bracket, pop '(' from stack, doesn't match ']', return false

### Why This is Optimal

- **Time Complexity O(n)**: We traverse the string once, and stack operations (push/pop) are O(1)
- **Space Complexity O(n)**: In the worst case (all opening brackets), we store all n characters in the stack
- This is optimal because we must examine every character to determine validity

### Alternative Approaches (Not Optimal)

1. **Counting Only**: Just count opening/closing brackets - Doesn't work because it ignores ordering and pairing
2. **Recursive Approach**: Find innermost matching pair and remove - O(n²) time due to string manipulation
