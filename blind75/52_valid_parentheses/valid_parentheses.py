def is_valid(s: str) -> bool:
    """
    Determines if the input string has valid parentheses.
    Uses stack for O(n) time complexity.

    Args:
        s: String containing only parentheses characters

    Returns:
        True if parentheses are valid, False otherwise
    """
    if len(s) % 2 != 0:
        return False

    stack = []
    pairs = {
        ')': '(',
        '}': '{',
        ']': '['
    }

    for char in s:
        if char in '({[':
            stack.append(char)
        else:
            if not stack:
                return False
            top = stack.pop()
            if pairs[char] != top:
                return False

    return len(stack) == 0
