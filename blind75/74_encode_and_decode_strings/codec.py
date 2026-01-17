from typing import List


class Codec:
    """
    Codec for encoding and decoding strings using length-prefix encoding.
    Time: O(n), Space: O(n)
    """

    def encode(self, strs: List[str]) -> str:
        """
        Encodes a list of strings to a single string.

        Args:
            strs: List of strings to encode

        Returns:
            Encoded string
        """
        result = []
        for s in strs:
            # Format: length + "#" + string
            result.append(str(len(s)) + "#" + s)
        return "".join(result)

    def decode(self, s: str) -> List[str]:
        """
        Decodes a single string to a list of strings.

        Args:
            s: Encoded string

        Returns:
            List of decoded strings
        """
        result = []
        i = 0

        while i < len(s):
            # Find the delimiter '#'
            j = i
            while s[j] != "#":
                j += 1

            # Parse the length
            length = int(s[i:j])

            # Extract the string using the length
            i = j + 1
            result.append(s[i : i + length])

            # Move pointer past the extracted string
            i += length

        return result
