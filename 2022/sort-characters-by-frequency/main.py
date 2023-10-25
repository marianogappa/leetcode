class Solution:
    def frequencySort(self, s: str) -> str:
        sorted_tpls = sorted([(-count, char) for char, count in dict(Counter(s)).items()])
        return ''.join(
            [char*(-count) for count, char in sorted_tpls]
        )
