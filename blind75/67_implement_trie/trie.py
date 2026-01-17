class TrieNode:
    """Represents a node in the trie."""

    def __init__(self):
        self.children = {}
        self.is_end_of_word = False


class Trie:
    """A prefix tree (trie) data structure for efficient string operations."""

    def __init__(self):
        """Initializes the trie object."""
        self.root = TrieNode()

    def insert(self, word: str) -> None:
        """
        Inserts a word into the trie.
        Time complexity: O(m) where m is the length of the word.

        Args:
            word: The word to insert
        """
        current = self.root

        for ch in word:
            if ch not in current.children:
                current.children[ch] = TrieNode()
            current = current.children[ch]

        current.is_end_of_word = True

    def search(self, word: str) -> bool:
        """
        Returns true if the word is in the trie.
        Time complexity: O(m) where m is the length of the word.

        Args:
            word: The word to search for

        Returns:
            True if the word exists in the trie, False otherwise
        """
        current = self.root

        for ch in word:
            if ch not in current.children:
                return False
            current = current.children[ch]

        return current.is_end_of_word

    def starts_with(self, prefix: str) -> bool:
        """
        Returns true if there is any word in the trie that starts with the given prefix.
        Time complexity: O(m) where m is the length of the prefix.

        Args:
            prefix: The prefix to search for

        Returns:
            True if any word has this prefix, False otherwise
        """
        current = self.root

        for ch in prefix:
            if ch not in current.children:
                return False
            current = current.children[ch]

        return True
