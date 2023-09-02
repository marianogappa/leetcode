# As long as all nodes are visited only once, just traverse the graph 
# in any way you want (BFS, DFS, iterative, recursive),
# and for each tuple of (cloned_node, original_node):
# - Clone the neighbors
# - Connect the cloned neighbors to the current cloned node
# - Recurse to the cloned neighbors.
#
# To visit only once, keep a visited set. But since you need a reference
# to the cloned nodes when connecting already cloned neighbors, keep a map
# from val to cloned node instead. 

# Time: O(v*e) where v is number of nodes and e is number of neighbors.
# Space: O(v) because we keep a map of cloned nodes.
class Solution:
    def cloneGraph(self, node: 'Node') -> 'Node':
        if not node:
            return None

        cloned_root = Node(val=node.val)

        # Populate neighbors recursively, via DFS.
        # Keep a map of cloned nodes to avoid creating duplicates.
        dfs(cloned_root, node, {node.val: cloned_root})

        return cloned_root
    
def dfs(cloned_node: 'Node', orig_node: 'Node', cloned_nodes: dict[int, 'Node']) -> None:
    for neighbor in orig_node.neighbors:
        if neighbor.val not in cloned_nodes:
            # We haven't cloned this node yet: clone it.
            cloned_nodes[neighbor.val] = Node(val=neighbor.val)

            # Connect it.
            cloned_node.neighbors.append(cloned_nodes[neighbor.val])

            # Clone the neighbors of this neighbor recursively (DFS!)
            dfs(cloned_nodes[neighbor.val], neighbor, cloned_nodes)
        else:
            # We mustn't clone the node again (we have it). Connect it!
            cloned_node.neighbors.append(cloned_nodes[neighbor.val])
