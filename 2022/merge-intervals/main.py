# Time: O(n*logn)
# Space: O(n)
#
# Straightforward. Intervals can be greedily merged left to right
# as long as they are sorted by start ascending.
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        sorted_intervals = sorted(intervals)

        result = [sorted_intervals[0]]
        for i in range(1, len(sorted_intervals)):
            if is_overlap(result[-1], sorted_intervals[i]):
                result[-1] = merge_intervals(result[-1], sorted_intervals[i])
            else:
                result.append(sorted_intervals[i])

        return result


def merge_intervals(i1, i2: List[int]) -> List[int]:
    return [min(i1[0], i2[0]), max(i1[1], i2[1])]


def is_overlap(i1, i2: List[int]) -> bool:
    return min_interval(i1, i2)[1] >= max_interval(i1, i2)[0]


def min_interval(i1, i2: List[int]) -> List[int]:
    return i1 if i1[0] <= i2[0] else i2


def max_interval(i1, i2: List[int]) -> List[int]:
    return i1 if i1[0] > i2[0] else i2
