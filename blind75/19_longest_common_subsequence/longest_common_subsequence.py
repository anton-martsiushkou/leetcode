def longest_common_subsequence(text1: str, text2: str) -> int:
    """
    Returns the length of the longest common subsequence.
    Uses dynamic programming with O(m*n) time and space complexity.

    Args:
        text1: First string
        text2: Second string

    Returns:
        Length of the longest common subsequence
    """
    m, n = len(text1), len(text2)

    # dp[i][j] = LCS length of text1[0...i-1] and text2[0...j-1]
    dp = [[0] * (n + 1) for _ in range(m + 1)]

    # Build DP table
    for i in range(1, m + 1):
        for j in range(1, n + 1):
            if text1[i - 1] == text2[j - 1]:
                # Characters match, extend previous LCS
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                # Take maximum of excluding either character
                dp[i][j] = max(dp[i - 1][j], dp[i][j - 1])

    return dp[m][n]
