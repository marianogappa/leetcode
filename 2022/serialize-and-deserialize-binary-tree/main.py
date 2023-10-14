# Pre-order DFS is the most intuitive solution, but the tree needn't be balanced,
# so a sparse tree must still use linear space.
#
# To solve this, for every node, always emit an entry for their two children, and use
# a special character if they are `None`. Don't recurse `None` children.
#
# The result of serialisation is a list of either node values or `None` characters, which
# can be joined by a separator.
#
# Deserialising is just doing another pre-order traversal, this time creating the nodes
# using the serialised list of values. When a `None` character is popped, leave the child
# empty and don't recurse over it.

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None
STR_NONE = "N"
NUM_NONE = float("inf")
SEPARATOR = "|"

# Time: O(n) both functions are linear since they go through every node once.
# Space: O(n) both functions use some constant space per node, so linear.
class Codec:

    def serialize(self, root):
        """Encodes a tree to a single string.
        
        :type root: TreeNode
        :rtype: str
        """
        if not root:
            return ''

        serialized = []
        _serialize(root, serialized)
        return SEPARATOR.join(serialized)


    def deserialize(self, data):
        """Decodes your encoded data to tree.
        
        :type data: str
        :rtype: TreeNode
        """
        if not data:
            return None
        
        vals = [int(elem) if elem != STR_NONE else NUM_NONE for elem in data.split(SEPARATOR)]
        
        root = TreeNode(vals[0])
        _deserialize(root, vals, 1)
        return root

# Pre-order DFS appending every number, and a STR_NONE for each None.
def _serialize(root: Optional[TreeNode], serialized: list[int]):
    if not root:
        serialized.append(STR_NONE)
        return
    
    serialized.append(str(root.val))
    _serialize(root.left, serialized)
    _serialize(root.right, serialized)

# Pre-order DFS insertion, but don't continue when a None is found.
def _deserialize(root: TreeNode, vals: list[int], i: int) -> int:
    if i < len(vals):
        i += 1
        if vals[i-1] != NUM_NONE:
            root.left = TreeNode(vals[i-1])
            i = _deserialize(root.left, vals, i)
    
    if i < len(vals):
        i += 1
        if vals[i-1] != NUM_NONE:
            root.right = TreeNode(vals[i-1])
            i = _deserialize(root.right, vals, i)

    return i
