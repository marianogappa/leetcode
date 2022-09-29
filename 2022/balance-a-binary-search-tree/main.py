# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# Time: O(n)
# Space: O(n)
#
# Very trivial if O(n) space is allowed: just get the sorted list of vals by running an in-order
# dfs, and construct the BST by continuosly making the root the middle of the array and assigning
# the left subarray to the left child, and the right subarray to the right child.
class Solution:
    def balanceBST(self, root: TreeNode) -> TreeNode:
        return sorted_list_to_bst(in_order_dfs(root, []))

def in_order_dfs(root: TreeNode, sorted_vals: list[int]) -> list[int]:
    if not root:
        return
    
    in_order_dfs(root.left, sorted_vals)
    sorted_vals.append(root.val)
    in_order_dfs(root.right, sorted_vals)
    
    return sorted_vals
    
def sorted_list_to_bst(nums: list[int]) -> Optional[TreeNode]:
    if not len(nums):
        return None

    mid = len(nums) // 2
    return TreeNode(nums[mid], sorted_list_to_bst(nums[:mid]), sorted_list_to_bst(nums[mid+1:]))
