#wip

from typing import Optional
from dataclasses import dataclass
from collections import OrderedDict

@dataclass
class Node:
    freq: int
    keys: OrderedDict
    prev: Optional['Node'] = None
    next: Optional['Node'] = None

class LFUCache:

    def __init__(self, capacity: int):
        self.key_to_node = {}
        self.head = None
        self.capacity = capacity

    def get(self, key: int) -> int:
        if not self.key_to_node.get(key):
            return -1

        val = self.key_to_node[key].keys[key]

        # Bump freq
        self._bump_freq(key)

        return val

    def _bump_freq(self, key):
        # Get node with value
        old_node = self.key_to_node[key]
        val = old_node.keys[key]
        freq = old_node.freq

        # Now we have to add the key to the next freq.
        # Next node could not exist, have a greater freq or correct freq.
        if old_node.next is None:
            next_node = Node(freq=freq+1, prev=old_node, next=None, keys=OrderedDict())
            old_node.next = next_node
        elif old_node.next.freq != freq+1:
            next_node = Node(freq=freq+1, prev=old_node, next=old_node.next, keys=OrderedDict())
            old_node.next = next_node
        else:
            next_node = old_node.next
        
        next_node.keys[key] = val
        self.key_to_node[key] = next_node 

        del old_node.keys[key]
        # We might have emptied prev node by removing the last key        
        self._remove_node_if_empty(old_node)

    def _remove_node_if_empty(self, node):
        if len(node.keys) > 0:
            return

        if self.head == node:
            if node.next:
                node.next.prev = None
            self.head = node.next
            return
        
        node.prev.next = node.next
        if node.next:
            node.next.prev = node.prev


    def put(self, key: int, value: int) -> None:
        if self.get(key) != -1: # This will bump its freq
            # Since we bumped it already, just update its value
            self.key_to_node[key].keys[key] = value
            return

        # Since we added an item, evict if reached capacity
        self._evict()

        # It doesn't exist:

        # Ensure there is a head node with frequency 1
        if not self.head:
            # There's no head. Create it with frequency 1
            self.head = Node(freq=1, keys=OrderedDict())
        elif self.head.freq != 1:
            # There's a head but without frequency 1. Create it and hook it up.
            node = Node(freq=1, keys=OrderedDict())
            self.head.prev, node.next, node.prev, self.head = node, self.head, None, node
        
        self.head.keys[key] = value
        self.key_to_node[key] = self.head

    def _evict(self):
        if len(self.key_to_node) < self.capacity:
            return
        
        key_to_evict = next(iter(self.head.keys))
        del self.key_to_node[key_to_evict]
        del self.head.keys[key_to_evict]
        self._remove_node_if_empty(self.head)
