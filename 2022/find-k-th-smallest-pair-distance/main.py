# An `O(n^2)` solution would be to generate all pairs, heapify them by distance & pop k times.
#
# Instead:
# - Code an `O(n)` function to count the number of pairs with distance `<= d` (using sliding window).
# - Binary search for the smallest distance `d` that has at least `k` pairs.

class Solution(object):
    def smallestDistancePair(self, nums, k):
        nums = sorted(nums)
        
        def count_pairs_with_distance_up_to(max_distance):
            # Slide a window where distance is <= max_distance
            count = left = 0
            for right in range(len(nums)):
                # If right - left > max_distance, window is invalid
                # Shrink from left until valid
                while nums[right] - nums[left] > max_distance:
                    left += 1

                # The count of pairs with distance <= max_distance is
                # right-left: imagine you pivot on right, so from left to right-1
                # there are right-left pairs.
                count += right - left

            return count

        # Binary search for the smallest distance that has at least k pairs
        lo = 0
        hi = nums[-1] - nums[0] # largest possible distance
        while lo < hi:
            mi = (lo + hi) // 2
            if count_pairs_with_distance_up_to(mi) >= k:
                hi = mi
            else:
                lo = mi + 1

        return lo
