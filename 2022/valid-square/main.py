# wip!
# Please note that the trickiness arises from the fact that the
# square needn't be aligned with the cartesian plane.
#
# At that point, a square has four sides whose length must be
# equal and >0, and two diagonals whose length must be
# equal and >0.
#
# It's just that the points aren't sorted so we don't know how to
# pair them as sides or diagonals. Easiest solution to that
# problem is to try all permutations and that's it.
#
# A clever optimisation is figuring out that many of those 24
# permutations result in the same comparisons, so there are
# duplicate calculations. But it won't change the complexity.
from itertools import permutations

# Time: O(1)
# Space: O(1)
def euclid(p1: List[int], p2: List[int]) -> int:
    return (p2[0]-p1[0])**2 + (p2[1]-p1[1])**2


def is_square(p1, p2, p3, p4: list[int]) -> bool:
    return euclid(p1, p2) == euclid(p2, p3) == euclid(p3, p4) == euclid(p4, p1) > 0 and euclid(p1, p3) == euclid(p2, p4) > 0


class Solution:
    def validSquare(self, p1: List[int], p2: List[int], p3: List[int], p4: List[int]) -> bool:
        return any([is_square(*points) for points in permutations([p1, p2, p3, p4], 4)])
