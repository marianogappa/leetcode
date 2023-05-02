# Very straightforward: calculate path to both and find last node of common prefix.
# Just remember to backtrack the partial path; every append must be constant time.

from typing import Optional

class TreeNode:
    def __init__(self, x, left = None, right = None):
        self.val = x
        self.left = left
        self.right = right

# Time: O(n) unless perfectly balanced tree in which case O(h)
# Space: O(n) unless perfectly balanced tree in which case O(h)
class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        path_to_p = calculate_path(root, p, [root])
        path_to_q = calculate_path(root, q, [root])
        return calculate_prefix(path_to_p, path_to_q)


def calculate_path(root, target: 'TreeNode', partial: list['TreeNode']) -> Optional[list['TreeNode']]:
    if root is None:
        return None

    if root == target:
        return partial

    partial.append(root.left)
    left_path = calculate_path(root.left, target, partial)
    if left_path:
        return left_path

    partial.pop()
    partial.append(root.right)
    right_path = calculate_path(root.right, target, partial)
    if right_path:
        return right_path

    partial.pop()
    return None


def calculate_prefix(path_a, path_b: list['TreeNode']) -> 'TreeNode':
    for i in range(min(len(path_a), len(path_b))):
        if path_a[i] != path_b[i]:
            return path_a[i-1]

    return path_a[min(len(path_a), len(path_b))-1]

root = TreeNode(3,TreeNode(5, TreeNode(6), TreeNode(2, TreeNode(7), TreeNode(4))), TreeNode(1, TreeNode(0), TreeNode(8)))

print(Solution().lowestCommonAncestor(root, root.left, root.right).val)
