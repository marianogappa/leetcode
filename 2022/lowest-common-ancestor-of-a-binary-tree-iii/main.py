# wip!
class Node:
    def _init_(self, val):
        self.val = val
        self.left = None
        self.right = None
        self.parent = None

class Solution:
    def lowestCommonAncestor(self, p: 'Node', q: 'Node') -> 'Node':
        visited = {p.val, q.val}
        while True:
            if p.parent:
                p = p.parent 
                if p.val in visited:
                    return p
                visited.add(p.val)
            if q.parent:
                q = q.parent
                if q.val in visited:
                    return q
                visited.add(q.val)
