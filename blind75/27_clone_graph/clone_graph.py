from typing import Optional, Dict


class Node:
    """A node in an undirected graph."""

    def __init__(self, val: int = 0, neighbors: Optional[list] = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


def clone_graph(node: Optional[Node]) -> Optional[Node]:
    """
    Creates a deep copy of a graph using DFS.
    Time: O(N + E), Space: O(N)

    Args:
        node: Reference to a node in the graph

    Returns:
        Reference to the cloned node
    """
    if not node:
        return None

    cloned: Dict[Node, Node] = {}

    def dfs(n: Node) -> Node:
        if n in cloned:
            return cloned[n]

        # Create a clone
        clone = Node(n.val)
        cloned[n] = clone

        # Clone all neighbors
        for neighbor in n.neighbors:
            clone.neighbors.append(dfs(neighbor))

        return clone

    return dfs(node)


def clone_graph_bfs(node: Optional[Node]) -> Optional[Node]:
    """
    Creates a deep copy of a graph using BFS.
    Time: O(N + E), Space: O(N)

    Args:
        node: Reference to a node in the graph

    Returns:
        Reference to the cloned node
    """
    if not node:
        return None

    cloned: Dict[Node, Node] = {}
    queue = [node]

    # Clone the starting node
    cloned[node] = Node(node.val)

    while queue:
        current = queue.pop(0)

        for neighbor in current.neighbors:
            if neighbor not in cloned:
                # Clone the neighbor
                cloned[neighbor] = Node(neighbor.val)
                queue.append(neighbor)

            # Add the cloned neighbor to current clone's neighbors
            cloned[current].neighbors.append(cloned[neighbor])

    return cloned[node]
