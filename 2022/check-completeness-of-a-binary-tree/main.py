# The most straightforward solution is to use BFS in this case.
#
# Strategy to satisfy all rules:
# 1) on each level, if a node with right but no left nodes => FALSE!
# 2) on each level, index the node slots, and keep a "max_found" and "min_missing". If min_missing < max_found => FALSE!
# 3) on each level, if a missing node is found (note it'd be missing on next level), flag it. 2 levels deeper, if there
#    are nodes => FALSE!

from typing import Optional
from collections import deque

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Time: O(n)
# Space: O(n)
class Solution:
    def isCompleteTree(self, root: Optional[TreeNode]) -> bool:
        if not root:
            return True
        
        # Use a deque for standard DFS
        queue = deque()
        queue.append(root)

        level = 0
        found_incomplete_node_on_level = float("inf")

        while queue: # There are more nodes...
            level += 1
            min_missing = float("inf")
            max_found = float("-inf")

            # There are nodes one level deeper than a level where an incomplete node was found!
            if level > found_incomplete_node_on_level:
                return False

            for i in range(len(queue)):
                node = queue.popleft()

                # A node was found which has a right node, but no left node!
                if node.right and not node.left:
                    return False

                if node.left:
                    queue.append(node.left)
                    max_found = max(max_found, 2*i)
                else:
                    min_missing = min(min_missing, 2*i)
                    found_incomplete_node_on_level = min(found_incomplete_node_on_level, level+1)
                
                if node.right:
                    queue.append(node.right)
                    max_found = max(max_found, 2*i+1)
                else:
                    min_missing = min(min_missing, 2*i+1)
                    found_incomplete_node_on_level = min(found_incomplete_node_on_level, level+1)

            # Not all nodes on this level are as far left as possible!
            if min_missing < max_found:
                return False

        return True

print(Solution().isCompleteTree(TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(5)), TreeNode(3, TreeNode(6)))))
print(Solution().isCompleteTree(None))
print(Solution().isCompleteTree(TreeNode(1)))
print(Solution().isCompleteTree(TreeNode(1, TreeNode(2))))
print(Solution().isCompleteTree(TreeNode(1, None, TreeNode(2))))
print(Solution().isCompleteTree(TreeNode(1, TreeNode(1, TreeNode(1)), TreeNode(1, TreeNode(1), TreeNode(1)))))
print(Solution().isCompleteTree(TreeNode(1, TreeNode(1, TreeNode(1)), TreeNode(1, TreeNode(1, TreeNode(1)), TreeNode(1)))))
print(Solution().isCompleteTree(TreeNode(1, TreeNode(2, TreeNode(5)), TreeNode(3, TreeNode(7), TreeNode(8)))))
