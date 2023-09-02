# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# Time: O(n) where n is the number of nodes in the tree
# Space: O(h) where h is the height of the tree (h = n in the worst case)
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        val_by_level = []
        dfs(root, 0, val_by_level)
        return val_by_level

def dfs(root: Optional[TreeNode], level: int, val_by_level: list[int]):
    if not root:
        return
    
    if level >= len(val_by_level):
        val_by_level.append(root.val)
    else:
        val_by_level[level] = root.val
    
    dfs(root.left, level+1, val_by_level)
    dfs(root.right, level+1, val_by_level)
