# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        return do_build_tree(deque(preorder), inorder)

def do_build_tree(preorder, inorder):
    # If the subtree we sent is empty, we're done with this part
    if inorder:
        # First element of preorder is root of current subtree (or main tree)
        # Problem clarifies that values are unique so we can find the index 
        # in inorder by value.
        i = inorder.index(preorder.popleft())
        root = TreeNode(inorder[i])

        # In inorder traversal, the values to the left of x are the left subtree,
        # and vice versa. So we can just send the correct inorder subtree values.
        root.left = do_build_tree(preorder, inorder[0:i])
        root.right = do_build_tree(preorder, inorder[i+1:])
        return root
