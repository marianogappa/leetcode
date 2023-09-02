# Straightforward. Intervals can be greedily merged left to right
# as long as they are sorted by start ascending.

# Time: O(n*logn) because we're sorting
# Space: O(n) because we're creating a new list
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        sorted_intervals = sorted(intervals) # by start ascending (first list element)

        # We know the first interval is the first, because of the sorting
        merged_intervals = [sorted_intervals[0]]

        # For all other intervals, if they overlap with last, merge them
        # Otherwise, append them! That's it!
        for interval in sorted_intervals[1:]:
            if is_overlapping(merged_intervals[-1], interval):
                merged_intervals[-1] = merge_intervals(merged_intervals[-1], interval)
            else:
                merged_intervals.append(interval)
        
        return merged_intervals

def is_overlapping(i1: list[int], i2: list[int]) -> bool:
    return i1[1] >= i2[0] # >= because equal number considered overlapping

# The min function isn't necessary because we know i1[0] <= i2[0] (because of sorting)
def merge_intervals(i1: list[int], i2: list[int]) -> list[int]:
    return [min(i1[0], i2[0]), max(i1[1], i2[1])]
