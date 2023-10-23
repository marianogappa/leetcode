# This is DP with memo. Note that it's not enough to memo only the target!
#
# Given a position in nums and a target, the result of the decision tree from there will be duplicated work, so memo that.


from collections import defaultdict
from typing import List

class Solution:
    # Time: O(n * s) where n is the number of items and s is the sum of all items
    # Space: O(n * s) for memoization
    def canPartition(self, nums: List[int]) -> bool:
        # Let's try to find a subset of nums that sums to half the total
        total = sum(nums)
        # If the total is odd, we can't partition it
        return total % 2 == 0 and do_can_partition(nums, 0, total // 2, defaultdict(dict))

def do_can_partition(nums: list[int], i: int, target: int, memo: dict[dict[int, int]]) -> bool:
    # If we exceeded half, this path is invalid
    if target < 0:
        return False

    # If we reached half, this path is valid (rest of numbers will be the other half)
    if target == 0:
        return True

    # If we reached the end of the list, path is invalid (or we would have returned True)
    if i >= len(nums):
        return False

    # Don't repeat already calculated work
    if memo[i].get(target) is not None:
        return memo[i][target]

    # Try both paths: using or not using the current number on this half
    memo[i][target] = (
        do_can_partition(nums, i+1, target - nums[i], memo) or
        do_can_partition(nums, i+1, target, memo)
    )
    return memo[i][target]
