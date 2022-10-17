# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
#
# Time: O(nlogn) Most steps are linear, except the sorting step which is klogk where k is n/2 in the worst case
# Space: O(n)
class Solution:
    def verticalTraversal(self, root: Optional[TreeNode]) -> List[List[int]]:
        if root is None:
            return root
        # One pass to find out the horizontal dimensions of the tree
        min_col, max_col = find_width(0, 0, 0 ,root)
        # Construct the result array now that we know the dimensions
        result = [[] for _ in range((max_col - min_col + 1)) ]
        # BFS to traverse nodes in the correct vertical order, to populate properly
        populate_result(0 - (min_col), result, root)
        
        # Sort each column's values first by row and then by its value
        for i in range(len(result)):
            result[i] = sorted(result[i])

        # Now that it's sorted, leave only values on each column
        for i in range(len(result)):
            result[i] = [elem[1] for elem in result[i]]

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
    row = 0

    while queue:
        for _ in range(len(queue)):
            node, col = queue.popleft()
            result[col].append((row, node.val))
            if node.left:
                queue.append((node.left, col-1))
            if node.right:
                queue.append((node.right, col+1))
        row += 1
