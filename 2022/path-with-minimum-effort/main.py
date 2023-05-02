# The exercise is asking to calculate the "minimum effort" in a graph, so pretty much Dijkstra.
# The only "differences", if any:
# 1) Normally we have an adjacency list, and here we have a matrix, so up to 4 edges per vertex.
# 2) Normally edges distance and we sum distances: here we calculate max effort instead.

from typing import List
from queue import PriorityQueue

# Time: O(e*logv) but e < 4*v so => O(v*logv)
# Space: O(e + v) but e < 4*v so => O(v)
class Solution:
    def minimumEffortPath(self, heights: List[List[int]]) -> int:
        return dijkstra(heights, (0, 0), (len(heights[0])-1, len(heights)-1))

def is_in_bounds(v: tuple[(int, int)], graph: list[list[int]]) -> bool:
    return v[0] >= 0 and v[1] >= 0 and v[1] < len(graph) and v[0] < len(graph[0])

def unvisited_neighbors(v: tuple[(int, int)], visited: set[tuple[(int, int)]], graph: list[list[int]]) -> list[tuple[int, int]]:
    options = [(v[0]-1, v[1]), (v[0]+1, v[1]), (v[0], v[1]-1), (v[0], v[1]+1)]
    return [option for option in options if is_in_bounds(option, graph) and option not in visited]

def dijkstra(graph: list[list[int]], startVertex, endVertex: tuple[int, int]) -> int:
    dists = [[float("inf") for _ in range(len(graph[0]))] for _ in range(len(graph))]
    dists[startVertex[1]][startVertex[0]] = 0
    
    pq = PriorityQueue()
    pq.put((0, (startVertex[0], startVertex[1])))

    visited: set[tuple[(int, int)]] = set()

    while not pq.empty():
        (cur_dist, cur_vertex) = pq.get()

        visited.add(cur_vertex)

        for neighbor in unvisited_neighbors(cur_vertex, visited, graph):
            new_distance = max(cur_dist, abs(graph[cur_vertex[1]][cur_vertex[0]] - graph[neighbor[1]][neighbor[0]]))
            if new_distance < dists[neighbor[1]][neighbor[0]]:
                dists[neighbor[1]][neighbor[0]] = new_distance
                pq.put((new_distance, neighbor))
    
    return dists[endVertex[1]][endVertex[0]]

print(Solution().minimumEffortPath([[1,2,2],[3,8,2],[5,3,5]]), "== 2")
print(Solution().minimumEffortPath([[1,2,3],[3,8,4],[5,3,5]]), "== 1")
print(Solution().minimumEffortPath(
    [
        [1,2,1,1,1],
        [1,2,1,2,1],
        [1,2,1,2,1],
        [1,2,1,2,1],
        [1,1,1,2,1]
    ]
), "== 0")
