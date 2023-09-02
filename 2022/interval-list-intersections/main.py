# wip!
# It's trivial but watch the edge cases!
#
# Carefully write functions for checking for overlap and calculating intersection.
#
# For looping over, note te following:
# - The same pair of intervals cannot participate on two intersections; at least one "cursor" always moves to right.
# - EDGE CASE! If the pair of intervals end at the same number, then advance both cursors.
# - Otherwise, advance the cursor with the smallest interval end. The later ending interval may participate in next.
# - Intersections don't need to be merged because they cannot overlap. This is because intervals are disjoint.

from typing import List

# Time: O(n)
# Space: O(1) unless solution space counts in which case O(n)
class Solution:
    def intervalIntersection(
        self, firstList: List[List[int]], secondList: List[List[int]]
    ) -> List[List[int]]:
        fi = 0
        si = 0
        intersections = []
        while fi < len(firstList) and si < len(secondList):
            if is_overlap(firstList[fi], secondList[si]):
                intersections.append(intersection(firstList[fi], secondList[si]))

            si_inc = 1 if firstList[fi][1] >= secondList[si][1] else 0
            fi_inc = 1 if firstList[fi][1] <= secondList[si][1] else 0
            si += si_inc
            fi += fi_inc

        return intersections


def is_overlap(a: list[int], b: list[int]) -> bool:
    return min_start(a, b)[1] >= max_start(a, b)[0]


def min_start(a: list[int], b: list[int]) -> bool:
    return a if a[0] <= b[0] else b


def max_start(a: list[int], b: list[int]) -> bool:
    return a if a[0] > b[0] else b


def intersection(a: list[int], b: list[int]) -> list[int]:
    return [max(a[0], b[0]), min(a[1], b[1])]


print(
    Solution().intervalIntersection(
        [[0, 2], [5, 10], [13, 23], [24, 25]],
        [[1, 5], [8, 12], [15, 24], [25, 26]],
    )
)

print(Solution().intervalIntersection([[1, 3], [5, 9]], []))
