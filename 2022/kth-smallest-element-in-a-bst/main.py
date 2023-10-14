# Since this is a BST, one can just in-order dfs and build a sorted
# list, and then return the kth-1 element. But it'd be O(n) space.
#
# Instead of a list, pass a counter and a value (BY REFERENCE!)
#
# When counter reaches 0, set the value and STOP TRAVERSING!
#
# - Trick to pass by reference: pass a list with one element.
# - Trick to stop traversing: return when counter is 0.

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# Time: O(n) but only traverses until landing on kth element
# Space: O(1)
class Solution:
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        _v = [0]
        in_order(root, [k], _v)
        return _v[0]

def in_order(root: Optional[TreeNode], _k: list[int], _v: list[int]):
    if not root or _k[0] <= 0:
        return
    
    in_order(root.left, _k, _v)
    
    _k[0] -= 1
    if _k[0] == 0:
        _v[0] = root.val
        return

    in_order(root.right, _k, _v)
