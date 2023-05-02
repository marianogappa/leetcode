# Because the best path could potentially go up and left, there's no linear solution using memo.
# The zeroes in the grid are vertices, and their connected vertices form edges. Moving from one
# vertex to another has a cost of 1. This is straight up Dikstra!
#
# Just remember to add 1, because the number of moves == number of vertices - 1.

from typing import List
import heapq

# Time: O(n^2*log(n^2)) remember that n is the side length, not the number of cells.
# Space: O(n^2)
class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        # Edge case: if either the start or the end are not zeroes, there's no clear path
        if grid[0][0] == 1 or grid[len(grid)-1][len(grid)-1] == 1:
            return -1

        min_distances = dijkstra(grid, (0, 0))
        distance = min_distances[(len(grid[0])-1, len(grid)-1)]
        return distance + 1 if distance != float("inf") else -1

def dijkstra(grid: List[List[int]], start) -> dict[tuple[int, int], int]:
    # Time: O(n^2)
    distances: dict[tuple[int, int], int] = {(x, y): float("inf") for y in range(len(grid)) for x in range(len(grid))}
    distances[(0, 0)] = 0

    h: list[tuple[int, tuple[int, int]]] = []
    heapq.heappush(h, (0, start))

    visited: set[tuple[int, int]] = set()

    # n^2 iterations
    while len(h):
        # log(n^2)
        dist, vertex = heapq.heappop(h)

        # amortized constant
        visited.add(vertex)

        # up to 8 neighbors
        for neighbor in calculate_neighbors(vertex, grid):
            # amortized constant
            if neighbor not in visited and distances[neighbor] > dist + 1:
                distances[neighbor] = dist + 1
                # log(n^2)
                heapq.heappush(h, (distances[neighbor], neighbor))

    return distances

def calculate_neighbors(vertex: tuple[int, int], grid: List[List[int]]) -> tuple[int, int]:
    deltas = [(-1, 0), (0, -1), (1, 0), (0, 1), (-1, -1), (1, 1), (1, -1), (-1, 1)]

    neighbors = []
    for delta in deltas:
        candidate = (vertex[0] + delta[0], vertex[1] + delta[1])
        if is_in_bounds(candidate, grid) and grid[candidate[1]][candidate[0]] == 0:
            neighbors.append(candidate)
    
    return neighbors

def is_in_bounds(pos: tuple[int, int], grid: List[List[int]]) -> bool:
    return pos[0] >= 0 and pos[1] >= 0 and pos[0] < len(grid) and pos[1] < len(grid)

print(Solution().shortestPathBinaryMatrix([[0,1],[1,0]]), " == 2")
print(Solution().shortestPathBinaryMatrix([[0,0,0],[1,1,0],[1,1,0]]), " == 4")
print(Solution().shortestPathBinaryMatrix([[1,0,0],[1,1,0],[1,1,0]]), " == -1")
