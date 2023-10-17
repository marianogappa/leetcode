# Checking if a string is an anagram of another is checking their frequency map equality.
#
# So:
# - Build the frequency map of `p`
# - Build the frequency map of each window of `len(p)` in `s`.
# - If the maps are equal, an anagram is found!

class Solution:
    # Time: O(s), because we check a O(1) equality for every char in s.
    # Time: O(1) because we keep p_freq & s_freq, both maps with 26 chars max.
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(s) < len(p):
            return []

        anagrams = []
        p_freq = defaultdict(int)
        s_freq = defaultdict(int)

        # Build frequency map of p, and of the first window.
        for i in range(len(p)):
            p_freq[p[i]] += 1
            s_freq[s[i]] += 1
        
        if p_freq == s_freq:
            anagrams.append(0)

        for i in range(1, len(s)-len(p)+1):
            # Slide window by removing left and adding right.
            s_freq[s[i-1]] -= 1
            s_freq[s[i-1+len(p)]] += 1

            # Unfortunately, a key with a zero is not the same as
            # not having a key, so we must remove it manually.
            if s_freq[s[i-1]] == 0:
                del(s_freq[s[i-1]])

            if p_freq == s_freq:
                anagrams.append(i)
        
        return anagrams

# 🧠 Note that there is a solution that is up to 26x more efficient (same time complexity), but a lot trickier:
# - Checking the map equality is redundant. Only two characters change!
# - Instead, also keep a counter of characters that CONTRIBUTE (*) to `p_freq`.
# - If a removed character was contributing to the window, decrement counter.
# - If a new character contributes to the window, increment counter.
# - If `counter == total` after a window move, an anagram is found!
# - (*) Note that a new char that exists in `p` might not contribute if we already have it!
