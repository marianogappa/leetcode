# Build the paths to `p` & `q` via DFS, and return last matching value node between them.

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
class Solution:
    # Time: O(n) Since we might have to traverse the entire tree
    # Space: O(n) Worst case: both nodes are at the bottom of a skewed tree
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        # Find individual paths to nodes
        p_path = dfs(root, p, [root])
        q_path = dfs(root, q, [root])

        # Return last matching value node between paths
        return [
            p_path[i] for i in range(min(len(p_path), len(q_path)))
            if p_path[i].val == q_path[i].val
        ][-1]
    
def dfs(root: TreeNode, target: TreeNode, path: list[TreeNode]) -> TreeNode | None:
    if root == target:
        return path
    
    if root.left:
        path.append(root.left)
        if dfs(root.left, target, path):
            return path
        path.pop() # Don't forget to backtrack

    if root.right:
        path.append(root.right)
        if dfs(root.right, target, path):
            return path
        path.pop() # Don't forget to backtrack

    return None
