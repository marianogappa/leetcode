# This is just knowing how to implement Trie. Ops are all linear time & space.

from dataclasses import dataclass

@dataclass
class Node:
    is_word: bool
    nodes: dict[str, 'Node']

class Trie:
    def __init__(self):
        self.root = Node(is_word=False, nodes={})

    def insert(self, word: str) -> None:
        cursor = self.root
        for char in word:
            if not cursor.nodes.get(char):
                cursor.nodes[char] = Node(is_word=False, nodes={})
            cursor = cursor.nodes[char]
        cursor.is_word = True

    def _search(self, word: str) -> Node | None:
        cursor = self.root
        for char in word:
            if not cursor.nodes.get(char):
                return None
            cursor = cursor.nodes[char]
        return cursor

    def search(self, word: str) -> bool:
        node = self._search(word)
        return node.is_word if node else False

    def startsWith(self, prefix: str) -> bool:
        node = self._search(prefix)
        return bool(node)

# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)
