# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def isValidBST(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.doIsValidBST(root, -sys.maxsize-1, sys.maxsize)

    def doIsValidBST(self, root, min, max):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if root == None or (root.left == None and root.right == None):
            return True
        elif root.left != None and (root.left.val >= root.val or root.left.val <= min):
            return False
        elif root.right != None and (root.right.val <= root.val or root.right.val >= max):
            return False
        else:
            return self.doIsValidBST(root.left, min, root.val) and self.doIsValidBST(root.right, root.val, max)
