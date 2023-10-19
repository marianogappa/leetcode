# Two-pointer solution (sliding window):
#
# - Move `right` until a dupe character is found (use a `set` to find dupes in `O(1)`)
# - Keep track of the max window length every time window is enlarged!
# - When a dupe is found, advance left (evicting chars from window) until and including the first occurrence of the dupe char

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = 0
        window = set([])
        max_len = 0

        # For every char...
        for right in range(len(s)):
            # If not in window
            if s[right] not in window:
                # Add to window and keep track of max window len
                window.add(s[right])
                max_len = max(max_len, len(window))
            # If already in window
            else:
                # Advance left pointer until first occurrence of this char
                while s[left] != s[right]:
                    # Removing characters from window as left advances
                    window.remove(s[left])
                    left += 1
                # When exiting loop, left stands at first occurrence! Move once more.
                left += 1
        
        return max_len
