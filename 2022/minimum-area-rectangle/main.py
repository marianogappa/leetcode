# wip!
# Time: O(n^2) where n is len(points)
# Space: O(n) because we put each point in the set
class Solution:
    def minAreaRect(self, points: List[List[int]]) -> int:
        pointSet: set[tuple[int, int]] = {(point[0], point[1]) for point in points}

        min_area = float("inf")
        for i, p1 in enumerate(points):
            for j in range(i):
                p2 = points[j]
                
                if p1[0] != p2[0] and p1[1] != p2[1] and (p1[0], p2[1]) in pointSet and (p2[0], p1[1]) in pointSet:
                    x_gap = abs(p1[0] - p2[0])
                    y_gap = abs(p1[1] - p2[1])
                    min_area = min(min_area, x_gap * y_gap)

        return 0 if min_area == float("inf") else min_area
