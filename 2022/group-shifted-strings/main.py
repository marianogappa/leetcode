# wip!
# To avoid the quadratic comparisons, a linear solution is to hashmap based on a tuple of the distances between letters.

from collections import defaultdict

# Time: O(n*k) where n is length of strings and k is the length of the largetst string
# Space: O(n*k) same as above
class Solution:
    def groupStrings(self, strings: List[str]) -> List[List[str]]:
        result = defaultdict(list[str])
        for s in strings:
            result[self.getDiffString(s)].append(s)
        return result.values()

    def getDiffString(self, s: str):
        return str(
            [
                ord(s[i]) - ord(s[i-1])
                if ord(s[i]) - ord(s[i-1]) >= 0
                else 
                    26 - (ord(s[i-1]) - ord(s[i]))
                for i in range(1, len(s))
            ]
        )
