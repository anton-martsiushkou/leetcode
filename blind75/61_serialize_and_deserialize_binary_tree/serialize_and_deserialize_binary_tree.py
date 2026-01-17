from typing import Optional


class TreeNode:
    """Definition for a binary tree node."""

    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Codec:
    """Codec to serialize and deserialize binary trees."""

    def serialize(self, root: Optional[TreeNode]) -> str:
        """
        Encodes a tree to a single string using preorder traversal.

        Args:
            root: Root node of the binary tree

        Returns:
            Serialized string representation of the tree
        """
        values = []

        def dfs(node: Optional[TreeNode]) -> None:
            if node is None:
                values.append("null")
                return

            values.append(str(node.val))
            dfs(node.left)
            dfs(node.right)

        dfs(root)
        return ",".join(values)

    def deserialize(self, data: str) -> Optional[TreeNode]:
        """
        Decodes your encoded data to tree.

        Args:
            data: Serialized string representation

        Returns:
            Root node of the reconstructed tree
        """
        values = data.split(",")
        self.index = 0

        def dfs() -> Optional[TreeNode]:
            if self.index >= len(values) or values[self.index] == "null":
                self.index += 1
                return None

            node = TreeNode(int(values[self.index]))
            self.index += 1

            node.left = dfs()
            node.right = dfs()

            return node

        return dfs()
