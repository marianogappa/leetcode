# The intuitions of using a Counter and a running sum come naturally, but the logic to apply is
# incredibly tricky!
#
# Note that the third rule can be ignored because it's covered by the broader second rule.
#
# Invert the two first conditions, so that you get the "AND" rules for requesting friendship.

from typing import List
from collections import Counter

# Time: O(n)
# Space: O(n)
class Solution:
    def numFriendRequests(self, ages: List[int]) -> int:
        age_to_count = Counter(ages)

        # gt_count is the count of people greater than a certain age
        gt_count = [0] * 120
        running_sum = 0
        for age in range(120, 0, -1):
            running_sum += age_to_count.get(age, 0)
            gt_count[age-1] = running_sum

        # To calculate total requests, go through every x and apply rules to calculate how
        # many friend requests x would send:
        total_requests = 0
        for x in ages:
            """
                age[y] > 0.5 * age[x] + 7
                age[y] <= age[x]

                x = 1 => y > 7 && y <= 1 # NEVER BEFRIEND
                x = 2 => y > 8 && y <= 2 # NEVER BEFRIEND
                x = 4 => y > 9 && y <= 4 # NEVER BEFRIEND
                x = 6 => y > 10 && y <= 6 # NEVER BEFRIEND
                x = 10 => y > 12 && y <= 10 # NEVER BEFRIEND
                x = 12 => y > 13 && y <= 12 # NEVER BEFRIEND
                x = 14 => y > 14 && y <= 14 # NEVER BEFRIEND
                x = 15 => y > 14 && y <= 15 # FRIEND if y == 15: we need to start checking
            """
            if x <= 14:
                continue

            """ So at this point, for x to befriend y:
                y must be greater than x // 2 + 7
                y must be less than or equal to x
            """
            min_y_age_not_inclusive = x // 2 + 7
            max_y_age_inclusive = x - 1  # Ignore special case for equal age

            total_requests += gt_count[min_y_age_not_inclusive] - gt_count[max_y_age_inclusive]

            # Special case for equal age: befriend everyone except oneself
            total_requests += age_to_count[x] - 1

        return total_requests


print(Solution().numFriendRequests([16, 16]))
print(Solution().numFriendRequests([16,17,18]))
print(Solution().numFriendRequests([20,30,100,110,120]))
