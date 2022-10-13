# Definition for a binary tree node.
# class TreeNode:
#     def _init_(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
# Time: O(nlogn)
# Space: O(n)
from collections import deque
from typing import List, Optional


class Solution:
    def verticalOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return root
        # One pass to find out the horizontal dimensions of the tree
        min_col, max_col = find_width(0, 0, 0 ,root)
        # Construct the result array now that we know the dimensions
        result = [[] for _ in range((max_col - min_col + 1)) ]
        # BFS to traverse nodes in the correct vertical order, to populate properly
        populate_result(0 - (min_col), result, root)
        return result
    
def find_width(min_seen :int, max_seen :int, cur:int,root: Optional[TreeNode]) -> tuple[int, int]:
    
    if root is None:
        return (min_seen +1 , max_seen-1)
    left_min, left_max = find_width(min(min_seen,cur-1), max_seen, cur-1, root.left)
    right_min, right_max = find_width(min_seen, max(max_seen,cur+1), cur+1, root.right)
     
    return (min(left_min, right_min), max(left_max,right_max))

def populate_result(cur:int, result:List[List[int]], root)-> None:
    queue = deque()
    queue.append((root, cur))

    while queue:
        for _ in range(len(queue)):
            node, col = queue.popleft()
            result[col].append(node.val)
            if node.left:
                queue.append((node.left, col-1))
            if node.right:
                queue.append((node.right, col+1))
