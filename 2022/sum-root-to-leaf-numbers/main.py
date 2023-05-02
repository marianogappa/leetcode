# Straightforward dfs. No tricks.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Time: O(n)
# Space: O(n) hopefully O(height) but there could be a really skewed tree
class Solution:
    def sumNumbers(self, root: Optional[TreeNode]) -> int:
        return dfs(root, [root])
        
def dfs(root, path: list[TreeNode]) -> int:
    if not root:
        return 0
    
    if not root.left and not root.right:
        return calculate_number(path)
    
    total = 0
    
    if root.left:
        path.append(root.left)
        total += dfs(root.left, path)
        path.pop()

    if root.right:
        path.append(root.right)
        total += dfs(root.right, path)
        path.pop()

    return total    
    
    
def calculate_number(path: list[TreeNode]) -> int:
    multiplier = 1
    num = 0
    for i in range(len(path)-1, -1, -1):
        num += path[i].val * multiplier
        multiplier *= 10
    
    return num
