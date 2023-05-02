# This may not be the most succinct strategy but it's the most intuitive to me.
#
# Optimise time by making path a deque rather than a list.

# Time: O(n) two dfs'. This assumes using a deque for building path though.
# Space: O(n) technically O(h) but h could be n.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        # First, find path to target. This will inform the distance between
        # root and target.
        path = find_path_to_target(root, target, [])
        
        # Now that we know distance from target, and path to target, all
        # movements towards target decrease distance, and all
        # movements away increase distance. Collect every node at k or -k distance.
        return collect_nodes_at_k(root, path, len(path), k)

def collect_nodes_at_k(node: TreeNode, path: list[int], cur_distance: int, k: int) -> list[int]:
    # Stop recursion at empty nodes
    if not node:
        return []
    
    # If we're off the right path, distance only increases. If > k, we're done.
    if path is None and cur_distance > k:
        return []
    
    # If we were on the right path, but passed it, distance only decreases. If < -k, we're done.
    if path == [] and cur_distance < -k:
        return []
    
    # Start collecting results
    result = []
    
    # If this node is at the right distance, include it.
    if abs(cur_distance) == k:
        result.append(node.val)
    
    # We're gonna move left...
    if path is None or (path and path[0] == 1):
        # We're on the wrong path. Distance increases.
        new_cur_distance = cur_distance + 1
        new_path = None
    else:
        # We're on the right path. Distance decreases.
        new_cur_distance = cur_distance - 1
        new_path = path[1:] if path else []
        
    result += collect_nodes_at_k(node.left, new_path, new_cur_distance, k)

    # We're gonna move right...
    if path is None or (path and path[0] == -1):
        # We're on the wrong path. Distance increases.
        new_cur_distance = cur_distance + 1
        new_path = None
    else:
        # We're on the right path. Distance decreases.
        new_cur_distance = cur_distance - 1
        new_path = path[1:] if path else []
    
    result += collect_nodes_at_k(node.right, new_path, new_cur_distance, k)
    
    # Done. Return all collected node values.
    return result
        
def find_path_to_target(root: TreeNode, target: int, partial: list[int]) -> list[int]:
    if not root:
        return None
    
    # Found target using list of movement decisions in partial!
    if root.val == target.val:
        return partial
    
    # Try left path
    partial.append(-1)
    result = find_path_to_target(root.left, target, partial)
    if result:
        return partial
    partial.pop()

    # Try right path
    partial.append(1)
    result = find_path_to_target(root.right, target, partial)
    if result:
        return partial
    partial.pop()

    # Give up search from this node
    return None
