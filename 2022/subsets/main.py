from typing import List

# Time: O(n*2^n) read below why
# Space: O(n) needed to construct each solution
#
# It's very tricky to reason about time complexity, so here's an easy intuition:
# There are no wasted loops in producing the subsets, so one option is counting how many subsets there are:
# For each number, it either is or isn't on the subset:
# For [1, 2, 3] => (with 1/without 1) * (with 2/without 2) * (with 3/without 3) => 2^n subsets
# But to construct each subset we loop through all numbers, so: Time: O(n*2^n).
#
# How to construct subsets:
#
# You should probably study backtracking, but this solution is more intuitive for me:
#
# 1) Construct subsets of each length, from 0 to len(nums) inclusive.
# 2) For size == 0 => only the empty subset
# 3) For size == len(nums) => only a subset with all numbers
# 4) For other sizes, loop through nums, and construct all subsets of size-1 starting with current number, and
#    only using numbers to the right of current number.
#
# Please note the strategy of having a cursor over nums, and using numbers to the right of it. This is a clever way
# of reusing the nums array without resorting to creating extra space to supply "remaining nums" to the recursive
# iteration.
#
# If you want to study backtracking, use this: 
# https://leetcode.com/problems/subsets/discuss/27281/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)

class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        result = []
        for size in range(len(nums)+1):
            result += make_subsets_of_size(size, nums)
        
        return result

def make_subsets_of_size(size: int, nums: list[int], left: int = 0) -> list[list[int]]:
    if size == 0:
        return [[]]

    if size == len(nums)-left:
        return [nums[left:]]
    
    result = []
    for i in range(left, len(nums)):
        result += [[nums[i]] + subset for subset in make_subsets_of_size(size-1, nums, i+1)]
    
    return result

print(Solution().subsets([1,2,3]))
print(Solution().subsets([1,2]))
print(Solution().subsets([1,2, 3, 4]))
print(Solution().subsets([]))
print(Solution().subsets([1]))
