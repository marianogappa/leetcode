# The only trick here is to remember that each child must remember the min and max numbers seen before: a left child whose parent was a right child must still be greater than all of the left subtree of its grandparent.

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    # Time: O(n)
    # Space: O(n) in a very skewed BST due to stack space
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        # In the beginning, all values are allowed
        return dfs(root, float("-inf"), float("inf"))

def dfs(root: Optional[TreeNode], min_val: int, max_val: int) -> bool:
    if not root:
        return True
    return (
        # Basic condition the left & right children must satisfy in a BST
        (not root.left or min_val < root.left.val < root.val) and
        (not root.right or max_val > root.right.val > root.val) and
        
        # The whole left subtree cannot have values greater than this node
        dfs(root.left, min_val, root.val) and

        # The whole right subtree cannot have values lower than this node
        dfs(root.right, root.val, max_val)
    )
