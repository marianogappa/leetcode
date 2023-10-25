# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        head = current = ListNode()
        carry = 0
        while l1 or l2 or carry:
            val_1 = l1.val if l1 else 0
            val_2 = l2.val if l2 else 0
            raw_sum = val_1 + val_2 + carry
            sum = raw_sum % 10
            carry = raw_sum // 10
            current.next = ListNode(sum)
            current = current.next

            if l1:
                l1 = l1.next
            if l2:
                l2 = l2.next

        return head.next
