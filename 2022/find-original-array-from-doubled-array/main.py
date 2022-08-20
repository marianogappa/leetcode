from collections import defaultdict

# Time: O(n*logn) for sorting
# Space: O(n) for freq map and sorted list
#
# Initial intuition of considering that a number is in the original array if it has a double doesn't work; here's
# the counter example:
#
# 1 -> 2 -> 4 -> 8
#
# 2 has a double, but since it's already double of 1, it cannot be original and double at the same time.
#
# This means it matters the order in which we check. It's fair to say that, if the smallest number in the array has
# no double, then definitely the result is []. So, this points to the fact that array must be sorted first.
#
# At that point, it's pretty trivial: go left to right and check for doubles.
#
# Only catch is that we have to "remove/mark" numbers, because if we used them as doubles, we must ignore them later.
# Solution is to create a frequency map, and decrement frequency when used. This takes care of duplicates too.


class Solution:
    def findOriginalArray(self, changed: List[int]) -> List[int]:
        # Optimisation: odd number list must have a number without double
        if len(changed) % 2 != 0:
            return []

        # Make frequency map. Optimisation! Count odd numbers. There cannot be more than half the list.
        freqs = defaultdict(int)
        odd_count = 0
        for changed_num in changed:
            freqs[changed_num] += 1
            if changed_num % 2 == 1:
                odd_count += 1

        # If odd_count > len(changed) // 2, there must be a number without a double.
        if odd_count > len(changed) // 2:
            return []

        # Timsort numbers, in order to go from smallest to largest.
        sorted_numbers = sorted(changed)

        # Going from smallest to largest...
        result = []
        for num in sorted_numbers:
            # Number could have been used up. If so, ignore it.
            if not freqs[num]:
                continue

            # Mark it as used. It's important to mark it here and not later, because 0*2 == 0.
            freqs[num] -= 1

            # If double of this number does not exist (or was used up), then this is not a doubled array.
            double_num = num*2
            if not freqs.get(double_num):
                return []

            # Num has double! Then append it to the final list, and use up the double.
            result.append(num)
            freqs[double_num] -= 1

        return result
