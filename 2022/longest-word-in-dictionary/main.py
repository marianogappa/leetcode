# Intuition: put all words in a Trie (linear to words), then traverse Trie once
# in reverse lexicographical order greedily keeping a max length word, only
# traversing while current node is always a word stop. Should be linear to words too.
class Solution:
    # Time: O(w) where w is not len(words) but the whole content
    # Space: O(w)
    def longestWord(self, words: List[str]) -> str:
        t = Trie()
        for word in words:
            t.add(word)

        return ''.join(traverse_stop_words_reverse_find_max_len(t, [], []))


class Trie:
    def __init__(self):
        self.is_stop = False
        self.nodes = [None for _ in range(26)]

    def add(self, word: str) -> None:
        if not word:
            self.is_stop = True
            return

        idx = ord(word[0])-ord('a')

        if not self.nodes[idx]:
            self.nodes[idx] = Trie()

        self.nodes[idx].add(word[1:])


def traverse_stop_words_reverse_find_max_len(t: Trie, cur_word, max_word: list[str]) -> list[str]:
    if len(cur_word) >= len(max_word):
        max_word = cur_word[:]

    for i in range(len(t.nodes)-1, -1, -1):
        if not t.nodes[i] or not t.nodes[i].is_stop:
            continue

        cur_word.append(chr(i+ord('a')))
        max_word = traverse_stop_words_reverse_find_max_len(
            t.nodes[i], cur_word[:], max_word)
        cur_word = cur_word[:-1]

    return max_word
