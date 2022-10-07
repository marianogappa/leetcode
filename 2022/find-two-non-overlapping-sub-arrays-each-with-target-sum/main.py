from typing import List

# Time: O(n)
# Space: O(n)
#
# Solvable with a sliding-window. Note that this is simple because arr has no negatives nor zeroes.
#
# The tricky part is finding non-overlapping subarrays. By keeping a "best subarray up to ith", we
# can consult it upon finding each subarray, checking best subarray up to "before start of current".
class Solution:
    def minSumOfLengths(self, arr: List[int], target: int) -> int:
        start = running_sum = 0
        min_len_up_to = [float("inf")] * len(arr)
        result = float("inf")
        for end in range(len(arr)):
            running_sum += arr[end]
            if running_sum > target:
                # By summing the next element we exceeded target. Move the start of the range
                # to shrink the range until the sum is smaller or equal to target.
                #
                # e.g. [1, 2, 3] with target = 5 would sum 7, and we would miss [2, 3] if we
                # didn't advance "start"
                while running_sum > target:
                    running_sum -= arr[start]
                    start += 1
                
                # Note that after shrinking the range we might have found a range sum that
                # equals "target"! Therefore the next condition cannot be an "elif"!

            # Because of the way we transition the sliding window (and because there are no zeroes)
            # we're guaranteed to find all subarrays summing target on this if clause.
            #
            # We can construct an array that contains the minimum len subarray at each index.
            # This is so that when we find the next subarray, we can know the best subarray up to
            # before the start of the current one, and therefore find solutions without overlap.
            if running_sum == target:
                sum_len = end - start + 1

                if start > 0 and min_len_up_to[start - 1] < float("inf"):
                    result = min(result, min_len_up_to[start - 1] + sum_len)

                min_len_up_to[end] = min(min_len_up_to[end-1], sum_len) if end > 0 else sum_len
            else:
                min_len_up_to[end] = min_len_up_to[end-1] if end > 0 else float("inf")
        
        return result if result != float("inf") else -1

print(Solution().minSumOfLengths([3,2,2,4,3], 3), "== 2")
print(Solution().minSumOfLengths([7,3,4,7], 7), "== 2")
print(Solution().minSumOfLengths([4,3,2,6,2,3,4], 6), "== -1")

