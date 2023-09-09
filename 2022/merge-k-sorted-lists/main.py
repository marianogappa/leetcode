# ## Intuition
#
# Merging two sorted lists should be trivial, so the question is how to reuse to merge k:
#
# - Merge pairs of lists (with merge_2), 0 with 1, 2 with 3, etc,...
# - The result should be a list of merged pairs with half the length of original list.
# - Keep merging resulting pairs until there's only 1 list left. Return it 💥

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    # Time: O(l * log l)
    # Space: O(l) or O(1) if reusing lists doesn't count as space
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        while len(lists) > 1:
            lists = [
                merge_2_sorted_lists(
                    lists[i],
                    lists[i+1] if i+1 < len(lists) else None
                )
                for i in range(0, len(lists), 2)
            ]

        return lists[0] if lists else None

# Time: O(l1 + l2)
# Space: O(1)
def merge_2_sorted_lists(l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
    sentinel = cur = ListNode()

    while l1 and l2:
        if l1.val <= l2.val:
            cur.next = l1
            l1 = l1.next
        else:
            cur.next = l2
            l2 = l2.next
        cur = cur.next

    cur.next = l1 if l1 else l2

    return sentinel.next

# ## Time complexity analysis:
#
# - Each time we merge pairs we go through all nodes in all lists once. This is `O(l)`.
# - Multiply this `l` by how many times we merge pairs. How many? 🤔
# - Each time we merge pairs, we halve the number of lists. Thus, we do it `log(l)` times.
#
# ## Space complexity analysis (controversy as if this is constant or linear; discuss!):
#
# - Merging 2 is thought of `O(1)` because we just reassign the `.next`.
# - On each merging step we recreate & garbage collect `lists`, but this is still `O(l)`.
# - A more confusing algorithm is to reassign `lists` rather than recreate it. 🤷‍♂️
