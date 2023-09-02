# wip!
# Each recursive node in the BST should be at the middle of the list.
# So, recursively split the list in half (use slow/fast pointer), use
# middle node (or beginning of right-side list) as root.
#
# Splitting each list is O(n/2), but each time n => n/2. All iterations
# summed up should yield n*logn (similarly to quicksort).

# Time: O(n*logn)
# Space: O(n) or O(1) if solution space doesn't count
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def sortedListToBST(self, head: Optional[ListNode]) -> Optional[TreeNode]:
        if head is None:
            return None
        
        if head.next is None:
            return TreeNode(head.val)
        
        # At least there are 2 nodes, so we can split them
        left, right = split_list_in_half_tilt_left(head)
        
        return TreeNode(right.val, self.sortedListToBST(left), self.sortedListToBST(right.next))
        

# This splits into two lists, not exactly in half but makes a shorter left list.
# The right list should be a little larger since its first node will be used as
# root node.
#
# For even: 1, 2 => 1 | 2
# For even: 1, 2, 3, 4 => 1 | 2 3 4
# For odd: 1, 2, 3 => 1 | 2 3
# For odd: 1, 2, 3, 4, 5 => 1 2 | 3 4 5
def split_list_in_half_tilt_left(head: Optional[ListNode]) -> list[ListNode]:
    slow = head
    fast = head
    while fast and fast.next and fast.next.next and fast.next.next.next:
        slow = slow.next
        fast = fast.next.next

    left = head
    right = slow.next
    slow.next = None
    
    return left, right
