from typing import List
from collections import defaultdict


def valid_tree(n: int, edges: List[List[int]]) -> bool:
    """
    Checks if graph is a valid tree using DFS.
    Time: O(V + E), Space: O(V + E)

    Args:
        n: Number of nodes
        edges: List of edges [a, b]

    Returns:
        True if graph is a valid tree, False otherwise
    """
    # A tree with n nodes must have exactly n-1 edges
    if len(edges) != n - 1:
        return False

    # Build adjacency list
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    visited = set()

    def dfs(node: int, parent: int) -> bool:
        visited.add(node)

        for neighbor in graph[node]:
            if neighbor == parent:
                # Skip the edge we came from
                continue
            if neighbor in visited:
                # Found a cycle
                return False
            if not dfs(neighbor, node):
                return False

        return True

    # Start DFS from node 0
    if not dfs(0, -1):
        return False

    # Check if all nodes were visited (graph is connected)
    return len(visited) == n


class UnionFind:
    """Union Find data structure with path compression and union by rank."""

    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # Path compression
        return self.parent[x]

    def union(self, x: int, y: int) -> bool:
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return False  # Already in same set - cycle detected

        # Union by rank
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1

        return True


def valid_tree_union_find(n: int, edges: List[List[int]]) -> bool:
    """
    Checks if graph is a valid tree using Union Find.
    Time: O(V + E × α(V)), Space: O(V)

    Args:
        n: Number of nodes
        edges: List of edges [a, b]

    Returns:
        True if graph is a valid tree, False otherwise
    """
    # A tree with n nodes must have exactly n-1 edges
    if len(edges) != n - 1:
        return False

    uf = UnionFind(n)

    for a, b in edges:
        if not uf.union(a, b):
            # Cycle detected
            return False

    # With n-1 edges and no cycles, the graph must be connected
    return True
