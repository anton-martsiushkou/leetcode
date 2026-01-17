import pytest
from clone_graph import Node, clone_graph, clone_graph_bfs


def build_graph(adj_list: list[list[int]]) -> Node | None:
    """Build a graph from an adjacency list."""
    if not adj_list:
        return None

    nodes = [Node(i + 1) for i in range(len(adj_list))]

    for i, neighbors in enumerate(adj_list):
        for neighbor in neighbors:
            nodes[i].neighbors.append(nodes[neighbor - 1])

    return nodes[0]


def graph_to_adj_list(node: Node | None) -> list[list[int]]:
    """Convert a graph to an adjacency list."""
    if not node:
        return []

    visited = set()
    adj_list = []

    def dfs(n: Node):
        if n.val in visited:
            return
        visited.add(n.val)

        # Ensure adj_list has enough space
        while len(adj_list) < n.val:
            adj_list.append([])

        neighbors = [neighbor.val for neighbor in n.neighbors]
        adj_list[n.val - 1] = neighbors

        for neighbor in n.neighbors:
            if neighbor.val not in visited:
                dfs(neighbor)

    dfs(node)
    return adj_list


def test_example_1():
    """Test case from example 1 - four nodes."""
    adj_list = [[2, 4], [1, 3], [2, 4], [1, 3]]
    original = build_graph(adj_list)
    cloned = clone_graph(original)

    assert graph_to_adj_list(cloned) == adj_list
    assert original is not cloned


def test_example_2():
    """Test case from example 2 - single node."""
    adj_list = [[]]
    original = build_graph(adj_list)
    cloned = clone_graph(original)

    assert graph_to_adj_list(cloned) == adj_list
    assert original is not cloned


def test_example_3():
    """Test case from example 3 - empty graph."""
    adj_list = []
    original = build_graph(adj_list)
    cloned = clone_graph(original)

    assert cloned is None


def test_two_connected_nodes():
    """Test with two connected nodes."""
    adj_list = [[2], [1]]
    original = build_graph(adj_list)
    cloned = clone_graph(original)

    assert graph_to_adj_list(cloned) == adj_list
    assert original is not cloned


def test_bfs_example_1():
    """Test BFS implementation with example 1."""
    adj_list = [[2, 4], [1, 3], [2, 4], [1, 3]]
    original = build_graph(adj_list)
    cloned = clone_graph_bfs(original)

    assert graph_to_adj_list(cloned) == adj_list
    assert original is not cloned


def test_bfs_empty_graph():
    """Test BFS with empty graph."""
    adj_list = []
    original = build_graph(adj_list)
    cloned = clone_graph_bfs(original)

    assert cloned is None


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
