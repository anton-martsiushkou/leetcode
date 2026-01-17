from typing import List


class TrieNode:
    """Represents a node in the trie."""

    def __init__(self):
        self.children = {}
        self.word = None  # Stores the complete word at leaf nodes


class Solution:
    """Solution for Word Search II using Trie + DFS backtracking."""

    def find_words(self, board: List[List[str]], words: List[str]) -> List[str]:
        """
        Returns all words from the list that exist on the board.
        Uses Trie + DFS backtracking for O(M×N×4^L) time complexity.

        Args:
            board: 2D grid of characters
            words: List of words to search for

        Returns:
            List of words found on the board
        """
        if not board or not board[0]:
            return []

        # Build Trie from words
        root = self._build_trie(words)

        result = []
        m, n = len(board), len(board[0])

        # Try starting DFS from each cell
        for i in range(m):
            for j in range(n):
                self._dfs(board, i, j, root, result)

        return result

    def _build_trie(self, words: List[str]) -> TrieNode:
        """
        Constructs a Trie from the list of words.

        Args:
            words: List of words to insert into Trie

        Returns:
            Root node of the constructed Trie
        """
        root = TrieNode()

        for word in words:
            current = root
            for ch in word:
                if ch not in current.children:
                    current.children[ch] = TrieNode()
                current = current.children[ch]
            current.word = word  # Mark end of word

        return root

    def _dfs(
        self,
        board: List[List[str]],
        i: int,
        j: int,
        node: TrieNode,
        result: List[str],
    ) -> None:
        """
        Performs depth-first search with backtracking.

        Args:
            board: The board to search
            i: Row index
            j: Column index
            node: Current Trie node
            result: List to store found words
        """
        # Boundary check
        if i < 0 or i >= len(board) or j < 0 or j >= len(board[0]):
            return

        ch = board[i][j]

        # Check if already visited or character not in Trie
        if ch == "#" or ch not in node.children:
            return

        # Move to next Trie node
        node = node.children[ch]

        # Found a word
        if node.word is not None:
            result.append(node.word)
            node.word = None  # Avoid duplicate results

        # Mark cell as visited
        board[i][j] = "#"

        # Explore all 4 directions
        self._dfs(board, i - 1, j, node, result)  # up
        self._dfs(board, i + 1, j, node, result)  # down
        self._dfs(board, i, j - 1, node, result)  # left
        self._dfs(board, i, j + 1, node, result)  # right

        # Backtrack: restore cell
        board[i][j] = ch


def find_words(board: List[List[str]], words: List[str]) -> List[str]:
    """
    Convenience function for finding words on a board.

    Args:
        board: 2D grid of characters
        words: List of words to search for

    Returns:
        List of words found on the board
    """
    solution = Solution()
    return solution.find_words(board, words)
