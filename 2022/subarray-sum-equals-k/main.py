# The linear solution is very unintuitive.
#
# If we construct a "sums" dict, we can check if (running_sum - k) exists on every step.
# There are zeroes and negative values, so the same sum could be reached many times.
#
# That part is not SO tricky. The unintuitive part is why it works for all subarrays.
# So let's look at an example:
#
# For [1, 2, 3] and k = 3, how do [1, 2] and [3] work?
# - [1, 2] is clear: because sums will be [1, 3, ...] so 3 is reached
# - [3] is less clear: sums will be [1, 3, 6], so running_sum[6-3] will exist in the dict
from typing import List
from collections import defaultdict

# Time: O(n)
# Space: O(n)
class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        total = 0
        running_sum = 0
        sums = defaultdict(int)
        sums[0] = 1
        for num in nums:
            running_sum += num
            total += sums[running_sum-k]
            sums[running_sum] += 1
        
        return total


print(Solution().subarraySum([1,1,1], 2), "== 2")
print(Solution().subarraySum([1,2,3], 3), "== 2")
print(Solution().subarraySum([-1,-1,1], 0), "== 1")
