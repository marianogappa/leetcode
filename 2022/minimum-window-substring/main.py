from collections import defaultdict

class Solution:
    def minWindow(self, s: str, t: str) -> str:
        left = right = 0
        min_range = (float('-inf'), float('inf'))

        t_freq = defaultdict(int)
        count = 0
        for char in t:
            t_freq[char] += 1
            count += 1

        s_freq = defaultdict(int)
        for right in range(len(s)):
            char = s[right]
            # print(char, count, left, right)
            s_freq[char] += 1
            if t_freq.get(char) and s_freq[char] <= t_freq[char]:
                count -= 1
            
            # if count <= 0:
                # print(char, count, left, right)

            while count <= 0:
                if right-left+1 < min_range[1]-min_range[0]+1:
                    min_range = (left, right)

                char = s[left]
                s_freq[char] -= 1
                if t_freq.get(char) and s_freq[char] < t_freq[char]:
                    count += 1
                left += 1
        
        return s[min_range[0]:min_range[1]+1] if min_range != (float('-inf'), float('inf')) else ""


print(Solution().minWindow("ADOBECODEBANC", "ABC")) # "BANC"

# A[DOBECODEBA]NC
