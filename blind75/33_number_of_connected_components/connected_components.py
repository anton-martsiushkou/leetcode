from typing import List
from collections import defaultdict, deque


def count_components(n: int, edges: List[List[int]]) -> int:
    """
    Counts connected components using DFS.
    Time: O(V + E), Space: O(V + E)

    Args:
        n: Number of nodes
        edges: List of edges [a, b]

    Returns:
        Number of connected components
    """
    # Build adjacency list
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    visited = set()
    count = 0

    def dfs(node: int):
        visited.add(node)
        for neighbor in graph[node]:
            if neighbor not in visited:
                dfs(neighbor)

    # Count components
    for i in range(n):
        if i not in visited:
            count += 1
            dfs(i)

    return count


def count_components_bfs(n: int, edges: List[List[int]]) -> int:
    """
    Counts connected components using BFS.
    Time: O(V + E), Space: O(V + E)

    Args:
        n: Number of nodes
        edges: List of edges [a, b]

    Returns:
        Number of connected components
    """
    # Build adjacency list
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)
        graph[b].append(a)

    visited = set()
    count = 0

    def bfs(start: int):
        queue = deque([start])
        visited.add(start)

        while queue:
            node = queue.popleft()
            for neighbor in graph[node]:
                if neighbor not in visited:
                    visited.add(neighbor)
                    queue.append(neighbor)

    # Count components
    for i in range(n):
        if i not in visited:
            count += 1
            bfs(i)

    return count


class UnionFind:
    """Union Find data structure with path compression and union by rank."""

    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n
        self.components = n

    def find(self, x: int) -> int:
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])  # Path compression
        return self.parent[x]

    def union(self, x: int, y: int):
        root_x = self.find(x)
        root_y = self.find(y)

        if root_x == root_y:
            return  # Already in same component

        # Union by rank
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1

        self.components -= 1


def count_components_union_find(n: int, edges: List[List[int]]) -> int:
    """
    Counts connected components using Union Find.
    Time: O(V + E × α(V)), Space: O(V)

    Args:
        n: Number of nodes
        edges: List of edges [a, b]

    Returns:
        Number of connected components
    """
    uf = UnionFind(n)

    for a, b in edges:
        uf.union(a, b)

    return uf.components
