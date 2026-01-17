def num_decodings(s: str) -> int:
    """
    Returns the number of ways to decode the string.
    Uses dynamic programming with O(n) time and O(1) space complexity.

    Args:
        s: String containing only digits

    Returns:
        Number of ways to decode the string
    """
    if not s or s[0] == '0':
        return 0

    n = len(s)
    if n == 1:
        return 1

    # prev2 = number of ways to decode up to i-2
    # prev1 = number of ways to decode up to i-1
    prev2 = 1  # empty string has one way to decode
    prev1 = 1  # first character is valid (already checked)

    for i in range(1, n):
        current = 0

        # Check if single digit is valid (1-9)
        if s[i] >= '1' and s[i] <= '9':
            current += prev1

        # Check if two digits form a valid number (10-26)
        two_digit = int(s[i-1]) * 10 + int(s[i])
        if 10 <= two_digit <= 26:
            current += prev2

        prev2 = prev1
        prev1 = current

    return prev1
