# No point describing how to build subsets; look at the code. It's DFS with backtracking.
#
# It's very tricky to reason about time complexity, so here's an easy intuition:
#
# There are no wasted loops in producing the subsets, so one option is counting how many subsets there are:
#
# For each number, it either is or isn't on the subset:
#
# For [1, 2, 3] => (with 1/without 1) * (with 2/without 2) * (with 3/without 3) => 2^n subsets
#
# But to construct each subset we loop through all numbers, so: Time: O(n*2^n).
#
# Please note the strategy of having a cursor over nums, and using numbers to the right of it. This is a clever way
# of reusing the nums array without resorting to creating extra space to supply "remaining nums" to the recursive
# iteration.
#
# If you want to study backtracking, use this: 
#
# https://leetcode.com/problems/subsets/discuss/27281/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)

class Solution:
    # Time: O(n*2^n) read above why
    # Space: O(n) needed to construct each solution, otherwise matches time complexity
    def subsets(self, nums: List[int]) -> List[List[int]]:
        results = []
        for length in range(len(nums)+1):
            subsets_of_len(nums, 0, length, [], results)
        return results

def subsets_of_len(nums: List[int], num_start: int, length: int, partial: list[int], results: list[list[int]]) -> None:
    if len(partial) == length:
        results.append(partial.copy()) # Don't forget to clone! We'll modify partial later.
        return
    
    for i in range(num_start, len(nums)):
        partial.append(nums[i])
        # Note the cursor i+1! Don't reuse earlier numbers.
        subsets_of_len(nums, i+1, length, partial, results)
        partial.pop() # Don't forget to backtrack!
