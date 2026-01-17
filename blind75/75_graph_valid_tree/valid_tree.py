from typing import List


def valid_tree(n: int, edges: List[List[int]]) -> bool:
    """
    Checks if the edges form a valid tree using DFS.
    Time: O(E + n), Space: O(n)

    Args:
        n: Number of nodes
        edges: List of edges [u, v]

    Returns:
        True if edges form a valid tree, False otherwise
    """
    # A tree with n nodes must have exactly n-1 edges
    if len(edges) != n - 1:
        return False

    # Build adjacency list
    graph = {i: [] for i in range(n)}
    for u, v in edges:
        graph[u].append(v)
        graph[v].append(u)

    # Check if all nodes are connected using DFS
    visited = set()

    def dfs(node):
        visited.add(node)
        for neighbor in graph[node]:
            if neighbor not in visited:
                dfs(neighbor)

    dfs(0)
    return len(visited) == n


class UnionFind:
    """Union-Find (Disjoint Set Union) data structure."""

    def __init__(self, n: int):
        """Initialize Union-Find with n disjoint sets."""
        self.parent = list(range(n))
        self.rank = [1] * n

    def find(self, x: int) -> int:
        """Find the root of the set containing x with path compression."""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        """
        Union the sets containing x and y.
        Returns False if they're already in the same set (cycle detected).
        """
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return False  # Already in same set - cycle detected

        # Union by rank
        if self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        elif self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1

        return True


def valid_tree_union_find(n: int, edges: List[List[int]]) -> bool:
    """
    Checks if the edges form a valid tree using Union-Find.
    Time: O(E * α(n)) ≈ O(E), Space: O(n)

    Args:
        n: Number of nodes
        edges: List of edges [u, v]

    Returns:
        True if edges form a valid tree, False otherwise
    """
    # A tree with n nodes must have exactly n-1 edges
    if len(edges) != n - 1:
        return False

    uf = UnionFind(n)

    for u, v in edges:
        # If union returns False, we found a cycle
        if not uf.union(u, v):
            return False

    # If we have n-1 edges and no cycles, it's a valid tree
    return True
