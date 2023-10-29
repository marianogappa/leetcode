# Very simple DFS with memoization. Shouldn't be a hard exercise!

class Solution:
    def isInterleave(self, s1: str, s2: str, s3: str) -> bool:
        def dfs(i1, i2, i3, memo) -> bool:
            # If reached end of s3, true if other two strings are also empty
            if i3 == len(s3):
                return i1 == len(s1) and i2 == len(s2)
            
            # If reached end of s1, true if rest of s2 and s3 are equal
            if i1 == len(s1):
                return s2[i2:] == s3[i3:]

            # If reached end of s2, true if rest of s1 and s3 are equal
            if i2 == len(s2):
                return s1[i1:] == s3[i3:]
            
            # If next char in s3 isn't equal to either in s1 or s2, FALSE!
            if s1[i1] != s3[i3] and s2[i2] != s3[i3]:
                return False
            
            # Next letter is only from s1
            if s1[i1] == s3[i3] and s2[i2] != s3[i3]:
                return dfs(i1+1, i2, i3+1, memo)

            # Next letter is only from s2
            if s1[i1] != s3[i3] and s2[i2] == s3[i3]:
                return dfs(i1, i2+1, i3+1, memo)

            # If we've already seen this state, return it
            if memo.get((i1, i2)) is not None:
                return memo[(i1, i2)]

            # If next char in s3 is equal to both in s1 and s2, recurse on both and store answer
            memo[(i1, i2)] = (
                dfs(i1+1, i2, i3+1, memo) or
                dfs(i1, i2+1, i3+1, memo)
            )
            return memo[(i1, i2)]

        return dfs(0, 0, 0, {})
