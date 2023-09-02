# wip!
# Strategy: make a dbl-linked list (with head + tail) and a hashmap, and keep a size
# - Hashmap goes from key to Node for O(1) get+update.
# - With len(hashmap) we can check if size <= capacity, to know if put must evict.
# - Get must move Node to head of dbl-linked list
# - Put must add node on head, and possibly evict node from tail
#
# Just watch the million edge cases and separate into smaller sub-operations.

class Node:
    def __init__(self, key, val: int, prev: Optional['Node'] = None, next: Optional['Node'] = None):
        self.key = key
        self.val = val
        self.prev = prev
        self.next = next

class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.head = None
        self.tail = None
        self.key_to_node: dict[int, Node] = {}

    def get(self, key: int) -> int:
        if not self.key_to_node.get(key):
            return -1
        
        node = self.key_to_node[key]
        self._move_node_to_top(node)
        return node.val

    def put(self, key: int, value: int) -> None:
        # Node exists (update it)
        if self.key_to_node.get(key):
            node = self.key_to_node[key]
            node.val = value
            self._move_node_to_top(node)
            return
        
        # Node doesn't exist (create it)
        node = self._new_head_node(key, value)
        self.key_to_node[key] = node
        
        if len(self.key_to_node) > self.capacity:
            self.key_to_node.pop(self._evict_tail())
    
    def _move_node_to_top(self, node: Node) -> None:
        # At this point list cannot be empty, but could have 1 entry
        if self.head == node:
            return

        # Extract node from current position
        node.prev.next = node.next
        if self.tail != node:
            node.next.prev = node.prev
        else:
            self.tail = self.tail.prev
        
        # Insert node into head
        node.prev = None
        node.next = self.head
        self.head.prev = node
        self.head = node
    
    def _new_head_node(self, key, value: int) -> Node:
        # At this point list could be empty
        node = Node(key, value)
        
        if not self.head:
            self.head = node
            self.tail = node
            return node
        
        node.next = self.head
        self.head.prev = node
        self.head = node
        return node

    def _evict_tail(self) -> int:
        # At this point list cannot be empty and can't have one entry
        key = self.tail.key
        self.tail = self.tail.prev
        self.tail.next = None
        return key
