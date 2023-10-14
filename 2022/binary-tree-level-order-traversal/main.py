# Literally just BFS over the nodes and store their values per-level.

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# Time: O(n) because it traverses each node once
# Space: O(n) because it keeps each level of nodes in memory
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root:
            return []
        
        levels = []
        nodes = deque([root])

        while nodes:
            level = []
            for _ in range(len(nodes)):
                node = nodes.popleft()
                level.append(node.val)
                
                if node.left:
                    nodes.append(node.left)
                if node.right:
                    nodes.append(node.right)

            levels.append(level)
        
        return levels
