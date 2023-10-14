# Intuitively, since we want to avoid a `O(n^2)` solution, a sliding
# window approach is attractive.
#
# We can leverage the idea of the window being "valid": it's valid if it
# contains all characters in `t`. We must return the minimal "valid" window.
#
# - While the window is *invalid*, move right pointer to *enlarge it*.
# - While the window is *valid*, move left pointer to *shrink it*.
# - Keep track of the smallest VALID window found.
#
# It's crucial to answer if the current window is valid in `O(1)` time.
#
# To do this, initialise a frequency table for `t`, but also a `count` that
# sums all the frequencies.
#
# The `count` is the key: we can intelligently update it when we add or
# remove characters in the window, as long as they participate in the
# validity of the window. This way, the window is valid if `<= 0`.

from collections import defaultdict, Counter

# Time: O(s + t) since we iterate over `t` once and `s` up to twice.
# Space: O(s + t) since we store `t` in a Counter and `s` in a defaultdict.
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        t_freq    = Counter(t)
        # To check window validity in O(1), `count` keeps track of
        # how many characters in `t` are missing in the current window.
        count     = sum(t_freq.values())
        min_range = (float('-inf'), float('inf'))

        # Implement sliding window:
        # While invalid window, move right (thus enlarging).
        # While valid window, move left (thus shrinking).
        left   = 0
        s_freq = defaultdict(int)
        for right in range(len(s)):
            # We enlarged the window by moving right. Add new char to s_freq.
            s_freq[s[right]] += 1

            # This current char contributes to window validity only if it is
            # in `t` and we have not yet found all occurrences of it.
            if t_freq.get(s[right]) and s_freq[s[right]] <= t_freq[s[right]]:
                count -= 1

            # If window is valid, move left to shrink it.
            while count <= 0:

                # Since at this point the window is valid, check if it's also
                # the smallest window found so far, and update `min_range` if so.
                if right-left+1 < min_range[1]-min_range[0]+1:
                    min_range = (left, right)

                # We shrank the window by moving left. Remove char from s_freq.
                s_freq[s[left]] -= 1

                # Removing a char can make the window invalid if it is in `t`,
                # but only if we now have less occurrences of it than in `t`.
                if t_freq.get(s[left]) and s_freq[s[left]] < t_freq[s[left]]:
                    count += 1

                left += 1

        return s[min_range[0]:min_range[1]+1] if min_range != (float('-inf'), float('inf')) else ""
