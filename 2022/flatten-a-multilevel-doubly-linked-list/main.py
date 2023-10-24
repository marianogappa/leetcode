#wip!

"""
# Definition for a Node.
class Node:
    def __init__(self, val, prev, next, child):
        self.val = val
        self.prev = prev
        self.next = next
        self.child = child
"""

class Solution:
    def flatten(self, head: 'Optional[Node]') -> 'Optional[Node]':
        if not head:
            return None

        cur = head
        while cur:
            if cur.child:
                fix(cur)
            cur = cur.next
        
        return head

def fix(node: 'Node'):
    # Must save parent next, so we can attach it later to the end of the child list
    saved_next = node.next

    # We must fix 3 pointers:
    node.next = node.child # Flatten beginning of child list
    node.next.prev = node  # Prev of beginning of child list is now this node
    node.child = None      # Flattened list will have no children
    
    # Recursively fix child list items
    cur = node.next
    while cur:
        if cur.child:
            fix(cur)
        
        # Once we reach the end of the child list, attach the saved parent next
        if not cur.next:
            cur.next = saved_next
            if saved_next:
                saved_next.prev = cur
            break

        cur = cur.next
