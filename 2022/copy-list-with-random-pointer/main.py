class Node:
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val = int(x)
        self.next = next
        self.random = random

    def __str__(self) -> str:
        if not self:
            return ""
        return f"({self.val}) -> {self.next}"


class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        if not head:
            return head
        intersperse_duplicate_nodes(head)
        set_duplicate_random_pointers(head)
        return isolate_duplicate_list(head)


def intersperse_duplicate_nodes(head: "Optional[Node]") -> None:
    cur = head
    # Cur will always point to an original node
    while cur:
        # Duplicate original node
        new_node = Node(cur.val, cur.next)
        # Advance to next original node, but before, set the original's next to the duplicate
        cur.next, cur = new_node, cur.next


def set_duplicate_random_pointers(head: "Optional[Node]") -> None:
    cur = head
    # Cur will always point to an original node
    while cur:
        # If original node has a random pointer, note that the random pointer will always precede a duplicate
        if cur.random:
            # Set the duplicate of this original node's random pointer to the duplicate of the original's random
            cur.next.random = cur.random.next
        # Advance to the next original node
        cur = cur.next.next


def isolate_duplicate_list(head: "Optional[Node]") -> None:
    cur = head.next
    # cur always points to a duplicate
    while cur:
        # It might be the last node
        if cur.next:
            # If it's not the last, point its next to the next duplicate (may be None)
            cur.next = cur.next.next

        # Advance to the next duplicate
        cur = cur.next

    # Head still points to the first original node. Return its next (head of duplicates).
    return head.next


head = Node(1, Node(2, Node(3, Node(4))))
print(Solution().copyRandomList(head))
