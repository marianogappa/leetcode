class Solution:
    # Time: O(n)
    # Space: O(n)
    #
    # Intuition: 
    # - Store every number in the array as a key in a "subsequence_lens" dict.
    # - As value, in principle use: 1 (i.e.: at least a sequence of len 1 starts here).
    # - When we store it, check if the "num - difference" key exists.
    #   If so, its value + 1 is the len of the sequence that started somewhere before!
    def longestSubsequence(self, arr: List[int], difference: int) -> int:
        subsequence_lens: dict[int, int] = defaultdict(int)

        for num in arr:
            subsequence_lens[num] = subsequence_lens.get(num-difference, 0) + 1
        
        return max(subsequence_lens.values())
