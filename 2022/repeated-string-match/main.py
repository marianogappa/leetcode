import math

# Time: O(a+b)
# Space: O(a+b)
#
# Note: if the "in" Python keyword is allowed here, using KMP is redundant.
#
# - Obviously, as long as "a" is smaller than "b", "b" cannot be a substring.
# - NOT obvious: once "a" >= "b", either "b" is already a substring, or you need one more repeat, but no more.
class Solution:
    def repeatedStringMatch(self, a: str, b: str) -> int:
        minimum_repeat_times = int(math.ceil(float(len(b))/float(len(a))))
        lps = compute_lps(b)

        if kmp(lps, a*minimum_repeat_times, b) != -1:
            return minimum_repeat_times
        
        if kmp(lps, a*(minimum_repeat_times+1), b) != -1:
            return minimum_repeat_times + 1
        
        return -1

def compute_lps(pattern: str) -> list[int]:
    lps = [0] * len(pattern)
    for i in range(1, len(pattern)):
        ps_len = lps[i - 1]

        while ps_len > 0 and pattern[i] != pattern[ps_len]:
            ps_len = lps[ps_len - 1]

        if pattern[i] == pattern[ps_len]:
            ps_len += 1

        lps[i] = ps_len

    return lps


def kmp(lps, s, pattern: str) -> int:
    si = 0
    pi = 0

    while si < len(s) and pi < len(pattern):
        if s[si] == pattern[pi]:
            si += 1
            pi += 1
        elif s[si] != pattern[pi]:
            if pi == 0:
                si += 1
            else:
                pi = lps[pi - 1]

    return si - pi if pi >= len(pattern) else -1

print(Solution().repeatedStringMatch("abcd", "cdabcdab"), "== 3")
print(Solution().repeatedStringMatch("a", "aa"), "== 2")
print(Solution().repeatedStringMatch("aa", "a"), "== 1")
print(Solution().repeatedStringMatch("aa", "z"), "== -1")
print(Solution().repeatedStringMatch("baaaaa", "ab"), "== 2")
