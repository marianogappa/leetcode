class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

# Time: O(v+e)
# Space: O(v), or O(v+e) if solution space counts
#
# Can be solved with BFS, iterative DFS or in this case recursive DFS.
# As long as any node is not visited twice (i.e. keep a visited set), just
# traverse the graph and for each node create a new one if not created yet, 
# and connect their neighbors.
class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if not node:
            return None
        
        new_head = Node(node.val, None)

        if not node.neighbors:
            return new_head
        
        do_clone(node, {node.val: new_head}, {})
        return new_head

def do_clone(old_node: Node, new_nodes: dict[int, Node], visited: set[int]):
    visited.add(old_node.val)

    for old_neighbor in old_node.neighbors:
        if old_neighbor.val not in new_nodes:
            new_nodes[old_neighbor.val] = Node(old_neighbor.val)

        new_nodes[old_node.val].neighbors.append(new_nodes[old_neighbor.val])

        if old_neighbor.val not in visited:
            do_clone(old_neighbor, new_nodes, visited)
