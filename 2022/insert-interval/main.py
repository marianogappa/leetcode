# There are three subproblems here:
#
# 1. Merging two intervals:
#    Given two sorted intervals, merge them if they overlap.
#
# 2. Merging a list of intervals: 
#    Put the first interval in the result list, and for each subsequent interval, run (`1.`) against the last interval
#    in the result list. If they overlap, merge them, otherwise, append the interval to the result list.
#
# 3. Inserting an interval to list:
#    This could be creating a new list copying the initial plus adding the new interval where it belongs, but it's `O(n)`.
#    The alternative is to use the original list and always check for the right time to consider the new interval.
#
# The trade-off of not duplicating the list is that the code gets quite messy, because the new interval could be:
# - The first element of the list, so the hardcoded beginning of the result list is conditional on it.
# - In the middle of the list, so we need a condition in the loop.
# - The last element of the list, so we need a condition after the loop.


class Solution:
    # Time: O(n)
    # Space: O(n)
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        if not intervals:
            return [newInterval]

        # Once we insert the interval, we don't need to process it again, so we set this flag.
        consumed = False

        # The inserted interval could be the first element of the list.
        if newInterval[0] <= intervals[0][0]:
            new_intervals = [newInterval]
            consumed = True
            rest = intervals
        else:
            new_intervals = [intervals[0]]
            rest = intervals[1:]

        for interval in rest:
            # The inserted interval could be in the middle of the list.
            if not consumed and newInterval[0] <= interval[0]:
                consumed = True
                process_interval(new_intervals, newInterval)

            process_interval(new_intervals, interval)
        
        # The inserted interval could be the last element of the list.
        if not consumed:
            process_interval(new_intervals, newInterval)

        return new_intervals

def process_interval(intervals: list[list[int]], interval: list[int]) -> None:
    if is_overlap(intervals[-1], interval):
        intervals[-1] = merge(intervals[-1], interval)
    else:
        intervals.append(interval)

# This assumes that i1[0] <= i2[0], that is, the intervals are sorted by start time
def is_overlap(i1: list[int], i2: list[int]) -> bool:
    return i1[1] >= i2[0]

def merge(i1: list[int], i2: list[int]) -> list[int]:
    return [min(i1[0], i2[0]), max(i1[1], i2[1])]
