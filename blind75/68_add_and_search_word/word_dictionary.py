class TrieNode:
    """Represents a node in the trie."""

    def __init__(self):
        self.children = {}
        self.is_end_of_word = False


class WordDictionary:
    """Supports adding words and searching with wildcards."""

    def __init__(self):
        """Initializes the WordDictionary object."""
        self.root = TrieNode()

    def add_word(self, word: str) -> None:
        """
        Adds a word to the data structure.
        Time complexity: O(m) where m is the length of the word.

        Args:
            word: The word to add
        """
        current = self.root

        for ch in word:
            if ch not in current.children:
                current.children[ch] = TrieNode()
            current = current.children[ch]

        current.is_end_of_word = True

    def search(self, word: str) -> bool:
        """
        Returns true if the word is in the data structure.
        The word may contain dots '.' which can match any letter.
        Time complexity: O(m) for exact match, O(26^k × m) with k wildcards.

        Args:
            word: The word to search for (may contain '.' wildcards)

        Returns:
            True if the word matches any added word, False otherwise
        """
        return self._search_helper(word, 0, self.root)

    def _search_helper(self, word: str, index: int, node: TrieNode) -> bool:
        """
        Recursive helper for wildcard search.

        Args:
            word: The word being searched
            index: Current position in the word
            node: Current trie node

        Returns:
            True if a match is found, False otherwise
        """
        # Base case: reached end of word
        if index == len(word):
            return node.is_end_of_word

        ch = word[index]

        # Wildcard: try all possible children
        if ch == ".":
            for child in node.children.values():
                if self._search_helper(word, index + 1, child):
                    return True
            return False

        # Regular character: search specific child
        if ch in node.children:
            return self._search_helper(word, index + 1, node.children[ch])

        return False
