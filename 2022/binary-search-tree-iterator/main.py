# Very trivial if you know the trick. Use a stack and do this:
#
# To get the "next" node, consume to the left until you can't.
#
# When you return that node, add node.right to the stack
# if it exists, and then consume to the left again.
#
# Repeat until you run out of nodes in the stack.

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class BSTIterator:
    # Time: O(h)
    # Space: O(h)
    def __init__(self, root: Optional[TreeNode]):
        self.stack = [root]
        self._consume_left()
    
    # Time: O(h)
    # Space: O(h)
    def _consume_left(self):
        while self.stack and self.stack[-1].left:
            self.stack.append(self.stack[-1].left)
        
    # Time: O(h)
    # Space: O(h)
    def next(self) -> int:
        cur_node = self.stack.pop()
        if cur_node.right:
            self.stack.append(cur_node.right)
            self._consume_left()
        return cur_node.val

    # Time: O(1)
    def hasNext(self) -> bool:
        return len(self.stack) > 0
