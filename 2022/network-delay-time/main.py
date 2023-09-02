# wip!
# This exercise is just running Dijkstra and finding the max of the minimum paths. Only trick is that if any of the
# nodes stays as "inf" after running Dijkstra, the node is unreachable, so we must return -1.

from collections import defaultdict
import heapq
from typing import List

# Time: O(v*log v + e)
# Space: O(v + e)
class Solution:
    def networkDelayTime(self, times: List[List[int]], n: int, k: int) -> int:
        # We're gonna need a vertex's edges in constant time for Dikstra, so compute it first
        adjacency: dict[int, list[list[int]]] = defaultdict(list[list[int]])
        for time in times:
            adjacency[time[0]].append(time)
        
        # Run Dijkstra
        min_paths = dijkstra(adjacency, n, k)
        
        # Readability hack: nodes are 1-indexed so I created a 1-indexed min-path, so the 0-th element must be removed.
        min_paths = min_paths[1:]
        
        # Compute the max of the min_paths. But if any path was unvisited (i.e. still infinite), then return -1.
        return -1 if float('inf') in min_paths else max(min_paths)
    
def dijkstra(adjacency: dict[int, list[list[int]]], n: int, k: int) -> list[int]:
    # We start with default assumption that distance from k to every vertex is infinite, except for k.
    distances = [float('inf')] * (n+1)
    distances[k] = 0
    
    # We only want to visit each vertex once.
    visited: set[int] = set()

    # Use a heap to iteratively pop the shortest distance vertex, until all vertices are visited
    h = []
    heapq.heappush(h, (0, k))

    # O(v): While there are still unvisited vertices...
    while len(h) > 0:
        # O(log v): Pop the next shortest distance vertex
        cur_dist, vertex = heapq.heappop(h)
        
        # O(1): Get the vertex's edges and mark the vertex as visited
        edges = adjacency[vertex]
        visited.add(vertex)
        
        # O(e): For every edge...
        for edge in edges:
            [_, dest_vertex, distance] = edge

            # If the vertex is visited, or we already found a shortest distance, ignore the edge
            if dest_vertex in visited or distances[dest_vertex] <= cur_dist + distance:
                continue

            # Store the shortest distance and add it to the heap, so that we visit it later
            distances[dest_vertex] = cur_dist + distance
            heapq.heappush(h, (distances[dest_vertex], dest_vertex))
        
    return distances


times = [[34,32,29],[30,3,66],[49,44,96],[10,48,34],[30,13,63],[13,2,29],[36,48,13],[2,6,44],[29,14,73],[31,45,8],[14,24,39],[17,43,89],[12,39,48],[11,23,77],[25,38,91],[27,17,29],[24,33,1],[47,36,8],[41,14,71],[16,47,68],[26,10,41],[15,28,33],[4,41,21],[4,1,95],[33,26,33],[25,4,51],[6,29,83],[7,25,70],[35,38,83],[31,6,58],[42,50,72],[20,33,27],[10,7,30],[14,38,86],[40,43,38],[45,21,73],[3,46,44],[19,14,28],[8,48,10],[16,6,1],[15,31,70],[49,8,82],[7,41,94],[9,6,33],[47,45,25],[35,41,97],[9,45,44],[3,13,8],[33,22,55],[43,37,12],[44,3,94],[17,2,52],[1,8,44],[47,18,99],[30,15,45],[7,1,87],[6,9,31],[1,13,95],[11,47,68],[18,39,5],[30,35,7],[25,3,8],[29,5,18],[25,1,85],[26,32,1],[23,40,52],[12,45,62],[45,47,77],[40,7,23],[8,26,64],[15,8,0],[23,32,2],[33,12,8],[3,38,8],[41,15,91],[42,37,30],[46,48,26],[11,42,73],[40,33,81],[13,42,7],[29,7,14],[29,31,97],[33,15,60],[41,20,47],[49,10,61],[40,3,99],[20,1,96],[47,11,39],[22,50,71],[3,14,64],[10,43,72],[26,21,2],[23,4,19],[33,44,88],[32,45,99],[45,36,1],[44,8,3],[42,6,15],[31,27,14],[4,31,80],[19,41,24],[41,18,94],[46,41,49],[17,25,45],[40,6,50],[27,9,46],[34,13,37],[39,28,68],[15,45,21],[46,27,91],[38,1,34],[5,47,31],[48,22,43],[4,30,84],[11,32,60],[19,28,94],[33,16,32],[13,34,49],[9,7,94],[35,1,33],[40,31,15],[37,29,62],[10,22,17],[39,8,79],[22,19,49],[50,13,32],[15,47,46],[50,46,54],[8,40,40],[8,25,24],[42,15,56],[7,3,87],[21,12,60],[40,2,18],[19,23,36],[19,22,0],[47,49,38],[43,5,96],[34,45,3],[9,21,95],[25,29,53],[27,16,26],[40,27,15],[41,47,36],[15,33,54],[39,35,8],[26,31,8],[15,23,62],[13,7,56],[16,17,26],[13,41,56],[17,19,22],[43,8,0],[10,50,96],[46,14,80],[2,4,64],[12,32,44],[20,48,19],[32,10,27],[40,25,93],[18,38,35],[25,6,3],[47,50,12],[49,9,61],[45,7,26],[27,8,30],[43,20,63],[37,41,60],[8,23,69],[35,12,45],[30,32,84],[15,6,39],[20,4,88],[32,2,99],[45,48,55],[22,24,62],[4,18,33],[17,6,24],[27,14,48],[48,50,61],[45,30,42],[30,46,47],[21,16,50],[10,49,3],[9,44,50],[14,15,75],[28,20,94],[14,36,0],[40,14,63],[38,13,62],[49,7,43],[42,38,89],[41,9,54],[44,18,2],[49,28,47],[2,43,11],[29,22,80],[24,35,17],[33,37,62],[45,1,61],[4,45,54],[41,36,46],[49,42,41],[18,46,40],[47,44,41],[39,5,69],[11,38,92],[47,7,67],[4,35,85],[7,48,38],[1,4,79],[36,3,29],[27,44,49],[22,36,95],[15,27,86],[22,49,48],[11,43,94],[17,11,18],[6,16,53],[42,22,52],[32,24,58],[13,11,10],[6,42,10],[42,39,26],[20,25,99],[18,23,9],[13,20,26],[34,40,42],[46,9,86],[6,21,99],[16,15,37],[23,8,47],[14,47,26],[34,23,25],[38,17,64],[14,41,53],[49,33,96],[26,48,43],[35,16,22],[23,38,31],[22,4,27],[48,39,0],[26,41,97],[13,10,13],[28,2,66],[26,34,14],[3,31,57],[40,17,57],[48,3,24],[22,25,14],[3,47,67],[33,46,55],[35,23,97],[32,35,6],[47,12,58],[13,40,93],[30,44,26],[37,34,42],[47,33,96],[23,6,71],[37,10,69],[28,40,10],[11,24,68],[7,38,78],[22,29,72],[1,22,98],[39,7,92],[37,39,88],[10,46,11],[48,6,80],[4,49,97],[16,27,36],[26,37,11],[45,46,98],[23,17,76],[47,19,0],[7,39,93],[48,30,52],[43,41,49],[50,22,4],[49,21,20],[27,7,45],[15,20,91],[10,19,16],[28,47,12],[46,49,59],[3,6,49],[48,5,28],[50,49,19],[1,49,42],[8,49,62],[28,25,70],[25,39,89],[17,49,59],[7,47,71],[5,8,75],[12,8,72],[24,17,56],[26,17,98],[39,33,34],[27,25,2],[25,34,86],[3,42,34],[26,49,42],[1,10,51],[23,28,88],[48,46,45],[23,12,95],[9,17,63],[43,3,91],[17,7,68],[1,43,70],[44,28,47],[48,2,22],[22,31,62],[21,28,59],[43,39,55],[50,16,15],[19,31,9],[23,22,76],[30,40,78],[24,12,21],[3,9,63],[24,8,16],[48,31,48],[11,1,75],[20,46,15],[27,47,32],[44,15,12],[4,40,35],[30,19,75],[10,42,88],[31,47,49],[46,20,8],[14,20,91],[44,39,52],[39,6,29],[25,2,49],[28,43,85],[9,11,12],[18,5,52],[4,46,21],[13,12,42],[6,23,73],[49,14,78],[40,19,59],[34,30,7],[12,23,44],[9,38,88],[41,17,84],[13,31,45],[5,36,70],[31,37,33],[9,26,53],[18,2,11],[1,31,19],[12,3,15],[34,38,24],[21,26,89],[10,17,55],[46,31,88],[44,37,73],[14,17,9],[49,27,84],[4,44,37],[19,37,9],[43,10,79],[13,38,68],[24,6,24],[13,15,71],[38,30,40],[50,30,16],[36,33,12],[24,50,58],[44,38,67],[16,42,60],[25,13,89],[1,20,65],[22,5,29],[36,28,18],[17,20,17],[46,39,53],[30,1,78],[30,11,21],[17,30,36],[14,31,16],[26,22,68],[45,39,91],[1,18,26],[31,36,44],[9,19,91],[39,50,51],[32,46,27],[10,41,86],[33,11,33],[21,2,85],[40,30,67],[26,24,19],[31,13,8],[41,34,0],[29,35,3],[35,37,82],[44,2,86],[27,13,19],[44,25,63],[17,46,45],[32,4,85],[16,49,59],[5,4,45],[5,39,78],[9,25,6],[21,46,55],[34,4,93],[29,30,8],[46,25,5],[41,43,35],[8,43,46],[12,46,92],[4,33,98],[44,14,56],[29,15,76],[15,21,21],[43,50,81],[37,18,56],[29,12,43],[36,7,90],[49,2,82],[3,37,32],[46,50,41],[18,21,69],[32,50,33],[4,23,76],[4,5,67],[32,41,35],[38,20,48],[31,7,54],[21,6,58],[45,40,4],[22,16,99],[6,17,33],[11,40,61],[14,44,28],[16,34,6],[14,21,86],[19,29,43],[28,23,32],[34,20,62],[8,38,13],[15,4,46],[25,43,59],[41,38,40],[37,26,55],[21,32,71],[37,45,42],[25,15,70],[18,44,45],[50,28,16],[12,35,53],[26,1,57],[24,25,1],[36,15,91],[26,2,16],[31,24,99],[3,10,44],[47,34,60],[19,2,54],[16,3,33],[33,17,69],[36,42,69],[45,17,93],[9,40,26],[10,35,6],[3,12,78],[14,30,81],[35,3,29],[1,24,66],[44,43,61],[10,36,61],[14,46,0],[42,7,59],[3,41,65],[16,36,20],[34,9,95],[5,6,81],[9,37,31],[32,3,74],[24,16,86],[11,27,1],[7,29,42],[46,6,5],[13,43,77],[42,9,60],[2,5,16],[9,36,74],[45,20,55],[30,41,42],[24,26,11],[13,23,18],[20,37,60],[47,30,69],[9,31,25],[13,1,32],[32,6,31],[28,36,97],[9,27,39],[42,25,62],[17,36,77],[41,4,96],[46,7,85],[4,39,62],[31,28,24],[31,12,56],[39,45,35],[26,42,43],[2,32,34],[29,46,2],[14,23,92],[20,22,71],[36,17,77],[18,35,51],[45,11,84],[23,3,53],[48,45,97],[2,28,31],[30,25,86],[45,50,78],[50,1,38],[29,25,11],[14,43,21],[40,22,55],[44,19,51],[41,22,30],[24,1,6],[6,41,10],[28,42,63],[19,1,44],[13,24,43],[16,13,16],[50,9,55],[24,9,13],[35,6,19],[25,19,98],[17,34,68],[39,25,97],[44,23,99],[39,40,14],[5,14,21],[41,39,71],[49,46,46],[31,8,86],[29,26,69],[47,16,96],[46,26,72],[41,48,72],[29,40,86],[10,26,41],[31,26,88],[48,12,3],[30,21,89],[9,48,11],[7,13,7],[5,40,79],[4,26,27],[40,10,80],[50,35,73],[32,11,61],[13,47,11],[41,46,91],[7,12,61],[46,43,7],[34,47,82],[43,44,38],[29,19,22],[41,32,34],[38,45,7],[39,36,2],[32,38,96],[36,47,52],[50,44,67],[33,42,1],[44,41,69],[25,11,43],[15,43,77],[4,37,36],[45,24,38],[26,12,86],[14,8,38],[38,9,86],[38,22,80],[34,12,1],[47,22,66],[24,28,52],[29,50,98],[40,36,13],[49,26,77],[18,15,85],[6,38,87],[44,9,19],[29,23,72],[19,9,62],[42,36,66],[45,2,57],[34,37,70],[35,14,99],[38,10,14],[24,46,80],[48,18,17],[44,36,75],[14,22,30],[22,17,83],[29,2,21],[2,36,34],[3,40,11],[40,1,77],[3,11,84],[15,29,49],[45,42,71],[29,17,49],[20,30,90],[20,8,56],[19,10,61],[49,13,55],[6,10,84],[29,36,49],[18,48,44],[10,30,93],[40,47,94],[9,28,36],[30,37,92],[5,49,61],[16,12,9],[28,19,23],[21,24,84],[7,2,70],[8,21,19],[2,31,48],[47,20,96],[43,31,74],[30,12,10],[11,35,31],[22,46,11],[32,16,55],[38,18,77],[45,38,96],[48,20,97],[47,40,93],[39,20,67],[18,33,30],[42,1,40],[28,35,64],[5,45,78],[20,2,17],[39,30,13],[13,48,22],[19,48,57],[11,22,30],[50,8,68],[32,19,16],[45,16,90],[9,2,13],[47,21,55],[16,30,98],[29,49,81],[37,50,33],[50,42,2],[44,1,14],[18,6,81],[30,43,53],[13,39,23],[36,26,75],[21,40,22],[32,28,1],[16,45,8],[47,42,34],[5,38,88],[11,17,80],[17,32,21],[49,15,47],[28,38,86],[40,29,87],[3,50,36],[44,34,93],[7,20,31],[28,12,98],[15,34,6],[48,43,4],[5,25,37],[40,18,19],[49,6,51],[26,33,48],[32,48,83],[39,24,9],[8,17,43],[50,21,99],[15,39,72],[9,42,70],[43,30,46],[26,18,21],[37,44,60],[46,13,3],[43,25,52],[15,25,42],[7,43,46],[21,33,18],[29,24,41],[44,16,38],[18,40,34],[37,11,71],[33,23,16],[45,28,60],[19,32,25],[22,41,81],[46,32,74],[17,23,91],[8,9,80],[1,7,70],[4,36,19],[28,17,8],[50,20,79],[41,1,26],[24,27,41],[14,40,2],[24,20,72],[6,32,67],[6,40,89],[43,2,36],[23,20,44],[22,40,76],[19,44,36],[14,9,75],[16,33,46],[41,40,29],[10,45,18],[49,12,36],[43,15,72],[15,35,69],[39,26,45],[2,17,99],[41,35,43],[30,48,53],[7,16,35],[4,8,24],[35,17,20],[11,49,15],[10,32,16],[27,32,9],[41,11,29],[47,28,17],[2,35,93],[37,3,66],[42,19,98],[35,29,98],[50,6,31],[37,23,90],[20,40,68],[32,29,40],[42,11,38],[41,5,70],[37,20,76],[37,6,6],[15,19,42],[28,13,45],[1,28,4],[44,50,61],[32,7,19],[19,13,33],[37,8,20],[50,31,87],[33,49,47],[40,9,70],[24,11,99],[29,39,36],[24,40,49],[7,46,68],[26,9,54],[14,18,74],[24,45,77],[13,4,33],[2,14,55],[26,25,6],[6,15,42],[31,5,91],[37,2,12],[17,26,78],[43,40,68],[16,35,46],[21,23,8],[10,39,13],[37,13,20],[42,30,62],[22,34,6],[49,35,98],[39,48,93],[25,44,65],[49,4,90],[32,21,99],[7,17,59],[13,27,46],[33,48,94],[19,47,72],[47,37,19],[43,13,64],[8,11,17],[31,34,39],[32,26,55],[38,42,14],[3,20,12],[16,32,1],[4,3,45],[49,17,38],[40,50,18],[23,26,73],[13,19,40],[3,1,7],[25,45,11],[14,12,75],[17,28,38],[1,12,40],[46,19,58],[7,27,82],[1,40,11],[8,13,59],[36,13,77],[12,29,34],[29,47,61],[17,40,13],[7,37,52],[4,6,63],[10,34,97],[20,19,73],[29,6,87],[7,19,68],[43,48,14],[5,19,84],[6,27,72],[41,13,22],[34,16,6],[35,10,28],[47,35,26],[11,29,71],[39,1,33],[43,32,14],[30,17,29],[38,23,21],[24,34,1],[17,13,97],[11,45,83],[27,5,61],[14,34,12],[24,39,76],[40,20,94],[19,7,18],[2,44,97],[1,6,2],[32,49,51],[32,43,66],[17,8,60],[11,33,50],[39,22,69],[34,42,27],[2,33,58],[6,26,56],[23,1,18],[43,36,2],[2,48,65],[34,48,3],[39,46,97],[6,13,20],[26,19,48],[49,19,66],[4,19,35],[23,45,66],[16,7,90],[6,4,39],[48,34,95],[18,13,28],[43,7,94],[46,36,35],[4,28,42],[8,18,42],[39,32,20],[6,49,47],[22,15,95],[3,35,10],[39,14,5],[35,28,77],[33,14,64],[36,41,6],[49,36,22],[17,44,65],[32,39,19],[18,50,41],[5,24,74],[33,1,56],[48,15,93],[40,23,86],[20,5,8],[13,22,91],[2,42,60],[1,33,26],[40,39,1],[16,23,53],[24,22,69],[18,12,44],[16,29,37],[26,20,6],[32,30,68],[5,21,71],[22,30,9],[14,19,63],[31,43,16],[5,15,38],[33,31,89],[50,11,26],[33,10,61],[32,25,9],[41,23,10],[42,4,3],[3,34,76],[46,12,56],[48,28,88],[22,37,2],[22,27,26],[14,25,90],[20,24,59],[13,46,3],[39,19,11],[50,24,59],[16,44,13],[6,18,1],[9,39,78],[37,22,17],[10,20,71],[19,34,17],[22,26,66],[12,47,30],[11,20,89],[7,4,21],[11,3,58],[35,20,8],[30,6,98],[16,25,19],[43,42,32],[45,8,64],[3,16,20],[33,29,10],[5,42,63],[16,26,84],[15,14,22],[18,4,7],[23,16,24],[4,27,30],[46,22,25],[22,33,42],[9,1,39],[46,17,71],[29,27,32],[12,44,67],[21,4,51],[17,47,47],[38,24,20],[1,21,24],[16,8,83],[13,30,47],[18,31,65],[2,27,55],[2,47,31],[4,7,52],[25,40,37],[13,5,99],[9,22,79],[36,35,8],[35,2,41],[21,5,4],[28,41,29],[36,43,53],[27,30,17],[6,3,39],[27,33,76],[11,15,83],[36,29,47],[47,4,27],[13,33,39],[42,12,8],[14,27,36],[43,9,12],[5,2,25],[21,35,35],[2,41,51],[1,34,20],[11,10,64],[27,37,52],[18,25,47],[35,30,50],[25,31,89],[12,1,29],[5,7,31],[36,46,88],[34,15,14],[12,18,85],[19,4,38],[8,12,32],[7,18,64],[35,43,60],[23,27,74],[32,1,76],[19,15,19],[24,30,24],[44,46,82],[6,31,24],[3,43,43],[12,26,46],[40,46,12],[10,33,32],[33,35,23],[42,41,74],[18,34,19],[2,29,60],[7,31,54],[2,45,31],[45,12,80],[42,47,34],[20,35,67],[13,37,45],[46,21,57],[46,4,85],[45,4,9],[8,31,77],[1,29,64],[16,43,35],[48,41,7],[34,50,83],[43,21,16],[47,6,82],[49,20,54],[3,32,54],[27,26,15],[33,36,96],[46,30,85],[31,29,36],[35,13,41],[42,32,38],[46,15,92],[42,8,17],[37,14,11],[32,5,99],[27,2,63],[8,45,46],[48,42,12],[40,35,70],[48,23,35],[48,9,58],[29,13,57],[48,26,67],[12,15,11],[9,8,24],[1,42,94],[10,1,0],[13,49,24],[12,48,92],[20,28,4],[25,36,14],[39,49,81],[9,33,69],[31,46,32],[15,7,29],[7,26,56],[49,5,81],[47,38,86],[16,1,64],[11,30,97],[19,11,4],[17,1,50],[8,44,14],[2,23,60],[48,13,46],[22,11,16],[15,26,16],[50,41,6],[50,10,69],[8,39,86],[8,46,29],[24,49,4],[36,49,16],[27,50,33],[38,19,65],[7,24,20],[30,34,33],[39,31,14],[19,50,79],[35,33,50],[13,35,35],[19,3,98],[49,24,2],[38,43,67],[17,21,74],[2,9,7],[22,18,71],[36,40,2],[4,42,52],[5,18,7],[42,10,26],[38,29,66],[29,48,88],[3,49,99],[10,15,21],[25,50,85],[10,40,39],[43,22,99],[11,26,51],[42,44,29],[46,5,98],[12,31,49],[50,40,88],[24,31,29],[25,32,44],[36,8,95],[29,3,25],[17,35,11],[7,8,73],[4,14,71],[36,12,54],[36,24,11],[19,18,98],[26,13,26],[1,47,84],[20,32,79],[14,3,95],[49,16,50],[14,6,83],[12,34,46],[38,32,77],[41,30,13],[50,38,71],[3,44,72],[23,24,42],[30,33,40],[4,21,63],[43,34,53],[29,45,94],[2,8,53],[37,38,15],[13,18,49],[5,12,35],[31,40,51],[5,32,39],[10,37,16],[19,36,21],[1,38,69],[31,20,74],[7,50,65],[1,15,59],[42,24,95],[27,29,5],[47,14,26],[35,50,55],[28,8,63],[34,10,28],[44,35,30],[10,44,72],[33,41,34],[16,28,11],[50,27,67],[27,28,62],[37,35,46],[27,41,64],[1,26,4],[35,4,23],[49,38,88],[28,49,39],[32,42,50],[35,45,25],[17,31,25],[7,45,77],[24,13,74],[19,33,75],[11,6,8],[39,13,46],[16,46,46],[28,11,16],[5,22,51],[17,27,11],[20,34,30],[48,29,76],[28,14,98],[36,10,56],[26,43,85],[36,11,90],[7,14,58],[35,26,89],[1,9,83],[22,43,65],[14,33,49],[20,36,74],[1,32,5],[46,29,36],[3,21,43],[44,10,52],[43,27,73],[49,43,70],[2,19,50],[46,18,87],[26,7,17],[30,29,61],[34,11,30],[36,9,72],[17,10,42],[26,14,4],[12,41,68],[40,4,17],[25,10,26],[42,48,32],[7,32,15],[49,1,22],[19,21,29],[16,18,45],[27,46,7],[38,8,28],[20,29,54],[45,43,53],[35,11,73],[18,19,36],[49,50,93],[11,46,5],[36,1,3],[42,45,47],[45,31,58],[19,45,52],[16,11,41],[23,44,60],[12,2,5],[41,37,82],[16,48,51],[12,43,11],[36,38,71],[44,33,83],[43,6,24],[40,48,6],[23,13,55],[8,34,93],[24,5,7],[37,24,84],[27,18,41],[15,36,65],[23,10,29],[20,47,60],[29,41,98],[35,32,86],[22,7,22],[50,25,64],[2,21,60],[18,37,47],[9,13,89],[43,18,26],[30,20,99],[13,8,72],[20,31,71],[46,47,72],[18,7,41],[7,30,20],[25,17,9],[42,16,81],[6,25,55],[27,22,1],[1,39,59],[43,45,16],[30,8,62],[39,17,93],[9,41,59],[21,39,42],[12,5,75],[32,44,86],[40,11,55],[48,33,21],[41,27,65],[49,48,16],[5,41,8],[44,12,18],[39,4,97],[7,35,17],[38,49,27],[44,40,88],[37,42,55],[28,3,50],[25,35,54],[27,11,6],[7,40,93],[46,28,82],[38,34,65],[9,15,61],[9,5,16],[10,12,28],[1,3,6],[12,11,18],[11,7,14],[38,39,53],[23,7,90],[32,15,4],[45,25,8],[14,29,65],[22,6,90],[34,26,3],[10,16,23],[48,7,10],[32,23,66],[26,47,26],[14,45,47],[27,43,22],[4,11,34],[12,24,67],[12,10,77],[10,23,92],[32,17,24],[48,36,65],[14,11,16],[44,20,98],[36,22,26],[8,6,89],[38,31,99],[18,45,8],[36,25,92],[32,22,9],[24,47,40],[14,42,86],[25,33,61],[30,45,73],[22,10,31],[45,19,32],[45,15,19],[18,26,94],[15,18,15],[9,32,6],[48,8,88],[39,38,4],[37,15,43],[10,24,74],[4,43,58],[50,7,29],[18,32,23],[33,50,38],[29,37,54],[45,32,42],[25,12,6],[31,41,23],[15,17,58],[39,41,27],[22,8,87],[37,48,35],[50,4,18],[33,6,86],[39,47,21],[18,3,13],[7,9,26],[4,47,75],[15,40,46],[23,29,95],[37,17,93],[10,47,31],[1,41,2],[20,15,60],[34,21,70],[31,33,39],[31,32,80],[45,6,46],[3,4,93],[24,14,93],[40,24,64],[45,26,0],[46,11,22],[27,15,33],[50,3,55],[37,49,47],[14,37,41],[9,18,39],[33,9,40],[16,37,66],[42,14,69],[43,11,82],[31,4,72],[45,44,44],[3,36,93],[25,28,66],[41,10,50],[35,24,19],[5,9,91],[6,19,86],[36,31,30],[31,25,84],[7,11,76],[50,23,78],[1,14,32],[50,29,0],[39,16,80],[21,49,57],[30,38,98],[11,48,4],[6,46,4],[27,3,92],[27,31,93],[42,13,13],[29,11,26],[28,15,49],[47,8,12],[24,43,70],[16,21,3],[11,37,1],[48,21,16],[38,14,79],[39,23,3],[33,43,94],[47,43,16],[46,16,85],[14,13,14],[27,40,82],[22,1,57],[4,24,60],[26,6,51],[24,48,89],[21,13,15],[3,27,83],[15,10,71],[39,9,33],[44,11,8],[27,21,75],[12,28,84],[8,36,50],[18,41,75],[20,14,81],[8,47,99],[33,3,36],[5,3,56],[4,48,46],[50,37,86],[12,6,70],[25,42,53],[46,42,96],[26,23,79],[15,9,93],[11,34,5],[16,4,81],[33,32,65],[20,50,52],[44,31,97],[22,12,31],[15,11,15],[29,28,69],[42,33,56],[27,38,46],[48,16,45],[10,13,14],[46,2,98],[28,9,64],[33,2,49],[34,44,87],[2,40,41],[4,15,86],[15,32,58],[8,37,20],[18,14,18],[45,22,52],[45,29,21],[4,16,42],[48,32,74],[6,30,8],[14,16,28],[37,1,97],[50,15,67],[50,2,57],[39,11,1],[49,37,31],[33,21,57],[36,20,9],[15,49,95],[13,16,14],[23,5,60],[11,18,13],[12,4,5],[3,17,46],[38,37,55],[20,45,12],[35,7,75],[33,8,75],[28,7,27],[20,13,41],[27,10,57],[17,14,41],[21,30,20],[23,46,21],[27,4,71],[23,19,66],[22,38,81],[20,7,66],[19,35,19],[40,49,4],[48,40,79],[19,27,61],[45,35,64],[41,49,63],[32,31,91],[3,22,91],[33,39,85],[43,4,79],[5,1,16],[36,19,44],[47,13,78],[12,14,42],[4,12,57],[16,20,39],[44,17,4],[46,1,34],[21,17,62],[36,5,92],[28,4,81],[20,27,13],[5,23,54],[12,38,14],[41,29,70],[15,48,60],[28,31,21],[34,43,39],[12,13,58],[34,8,98],[31,3,49],[23,31,1],[16,38,46],[42,31,17],[22,28,91],[9,20,11],[41,3,51],[16,22,78],[21,25,57],[29,4,99],[33,5,21],[32,47,30],[24,32,86],[26,3,70],[22,23,44],[3,24,66],[36,23,96],[47,23,90],[50,47,22],[15,22,76],[35,44,50],[3,29,6],[11,12,48],[46,34,96],[43,14,42],[34,6,23],[8,35,12],[24,10,15],[21,45,12],[34,28,18],[37,32,8],[34,39,34],[4,17,48],[17,38,52],[36,32,85],[12,33,73],[26,28,20],[7,22,58],[43,17,79],[37,4,36],[41,44,32],[7,49,90],[35,25,8],[45,33,30],[35,34,86],[49,45,89],[20,17,20],[4,32,37],[33,13,39],[17,48,77],[2,24,85],[2,11,8],[42,27,8],[31,16,77],[21,27,15],[48,14,23],[36,39,42],[22,3,64],[33,7,78],[46,37,41],[26,11,31],[10,9,50],[3,18,36],[12,20,54],[50,43,96],[18,10,43],[46,38,63],[27,49,31],[34,41,24],[24,42,3],[21,41,97],[19,49,86],[35,8,55],[35,49,61],[26,4,72],[2,18,95],[11,16,93],[15,38,66],[23,43,52],[48,49,76],[10,14,7],[19,30,4],[25,48,5],[31,44,3],[7,21,17],[46,40,39],[40,44,39],[22,13,78],[48,10,47],[20,10,2],[49,30,35],[31,10,27],[14,7,34],[46,3,62],[25,24,64],[14,4,82],[42,21,66],[19,39,43],[32,34,74],[16,14,74],[33,18,4],[1,16,98],[26,29,66],[48,44,20],[21,31,43],[41,6,83],[19,25,52],[9,3,8],[5,13,94],[5,48,65],[47,5,0],[49,31,60],[45,49,34],[47,1,27],[33,25,89],[16,41,80],[37,33,84],[42,49,53],[16,39,91],[18,17,25],[11,41,21],[34,31,99],[24,23,37],[40,38,45],[15,44,8],[19,26,83],[5,20,31],[28,44,20],[39,12,61],[18,43,29],[28,26,7],[21,44,94],[8,3,67],[3,15,22],[1,46,79],[6,24,85],[50,45,11],[11,8,59],[31,48,81],[8,10,81],[28,10,3],[38,11,95],[50,26,48],[13,32,81],[49,3,52],[39,34,63],[39,37,8],[43,24,25],[9,10,45],[29,38,41],[20,26,38],[6,8,69],[29,43,82],[14,28,30],[15,41,53],[21,36,8],[4,13,82],[1,50,56],[38,21,68],[46,35,73],[41,16,57],[28,27,60],[25,20,8],[24,29,44],[38,35,91],[27,24,12],[39,2,78],[40,42,31],[8,20,60],[14,48,38],[11,36,0],[20,6,94],[48,1,7],[13,21,82],[2,15,65],[30,9,61],[20,49,36],[23,39,46],[35,40,10],[47,29,68],[10,27,67],[7,23,19],[17,24,69],[42,34,29],[49,29,38],[21,48,84],[32,37,69],[2,1,7],[27,6,48],[31,2,79],[48,25,40],[21,38,72],[30,22,32],[26,45,6],[6,47,85],[23,14,99],[10,25,45],[50,18,51],[15,13,33],[21,9,64],[21,22,75],[35,36,44],[27,12,79],[27,39,31],[28,6,30],[10,3,55],[48,11,34],[16,5,82],[41,26,74],[24,4,5],[38,28,36],[44,4,8],[15,46,83],[7,15,6],[10,5,48],[15,50,57],[47,17,82],[39,43,36],[3,45,10],[2,3,14],[50,33,34],[38,44,62],[19,12,52],[9,43,48],[44,29,84],[5,37,94],[7,28,36],[19,16,29],[34,49,35],[3,26,54],[31,11,60],[15,1,42],[35,18,23],[8,19,3],[12,30,42],[44,45,22],[13,6,4],[22,35,20],[41,28,44],[1,25,10],[32,36,73],[4,9,56],[1,11,54],[38,48,99],[17,12,51],[48,19,35],[42,23,35],[18,11,68],[42,20,15],[42,28,93],[11,21,67],[2,39,23],[11,39,84],[45,18,63],[10,31,90],[27,19,61],[10,38,88],[6,28,87],[4,29,9],[3,23,87],[40,32,74],[29,33,84],[34,3,13],[8,33,73],[34,19,63],[18,30,2],[36,30,28],[30,50,17],[46,45,50],[6,34,23],[30,26,92],[1,17,24],[45,10,9],[30,49,24],[11,50,51],[22,2,77],[17,5,98],[37,27,94],[47,39,24],[18,20,24],[44,47,82],[23,49,20],[25,49,24],[8,16,64],[10,28,27],[19,20,42],[20,43,18],[34,18,43],[5,16,11],[31,21,61],[45,5,76],[28,48,41],[8,7,47],[33,27,94],[37,25,58],[48,17,36],[48,47,40],[26,36,61],[6,48,89],[41,24,77],[24,44,64],[27,1,98],[13,44,79],[49,11,42],[21,3,77],[40,21,32],[40,45,9],[9,14,1],[45,37,52],[31,18,5],[2,20,42],[35,47,38],[26,5,85],[14,49,12],[23,25,32],[33,38,18],[25,27,5],[34,14,0],[11,28,23],[27,20,60],[34,29,94],[5,43,87],[11,9,46],[11,13,18],[28,29,83],[32,13,69],[18,24,23],[6,44,57],[3,2,54],[9,4,20],[24,18,57],[21,11,43],[32,18,72],[4,50,76],[45,23,35],[21,37,36],[16,31,33],[11,5,68],[30,39,55],[47,48,6],[35,22,69],[50,14,91],[28,5,9],[31,19,84],[32,40,98],[37,21,53],[27,23,67],[24,3,77],[31,17,68],[45,41,17],[25,18,25],[36,21,17],[26,46,90],[41,12,4],[9,30,79],[27,35,6],[23,15,53],[22,45,34],[20,12,38],[45,9,81],[8,4,25],[30,18,78],[45,3,18],[19,42,77],[14,35,9],[23,33,55],[44,30,11],[8,14,20],[19,17,22],[46,8,28],[15,37,26],[36,2,66],[5,17,15],[4,25,28],[35,48,1],[42,17,16],[42,18,80],[27,36,75],[42,29,10],[41,7,47],[25,7,89],[22,39,91],[43,26,47],[35,39,97],[42,46,33],[33,45,41],[49,40,97],[38,3,60],[29,34,94],[25,21,23],[23,36,80],[35,19,44],[33,28,37],[28,21,58],[42,2,28],[37,9,54],[28,37,96],[28,33,24],[41,19,34],[1,23,48],[2,26,55],[32,27,15],[45,14,99],[6,14,87],[3,8,66],[9,12,56],[38,50,79],[43,35,12],[16,40,97],[25,14,65],[25,5,3],[28,46,43],[23,2,2],[13,3,82],[29,16,74],[37,31,5],[34,1,86],[30,4,53],[23,50,15],[48,24,26],[24,21,50],[9,29,27],[37,47,1],[14,5,14],[9,35,18],[28,39,80],[47,15,73],[6,36,33],[33,24,14],[30,2,94],[31,22,17],[13,29,65],[44,32,11],[21,7,48],[50,36,11],[22,42,15],[8,2,64],[40,12,67],[18,27,23],[6,45,16],[11,2,17],[28,34,81],[49,18,50],[19,24,59],[43,19,17],[43,49,84],[2,37,77],[20,41,39],[11,44,63],[2,13,4],[43,12,89],[37,36,24],[44,48,19],[47,26,74],[41,45,27],[43,29,11],[26,16,45],[38,26,20],[18,16,36],[25,46,11],[4,22,31],[28,50,31],[35,31,70],[6,2,39],[20,16,35],[5,34,71],[50,34,89],[38,15,26],[43,23,65],[44,6,7],[37,12,76],[47,32,54],[18,8,68],[1,36,38],[41,50,43],[12,16,83],[32,20,82],[38,47,19],[14,50,4],[19,43,32],[37,28,38],[44,24,28],[4,20,61],[43,46,98],[6,20,23],[6,35,42],[6,22,81],[18,36,79],[37,40,36],[31,50,59],[49,25,63],[21,50,61],[32,9,1],[1,27,77],[13,36,97],[19,8,58],[8,5,84],[36,4,97],[37,30,26],[38,4,51],[40,15,40],[5,30,20],[25,8,1],[36,6,16],[48,27,5],[8,41,82],[21,18,63],[27,42,38],[18,42,27],[8,30,21],[42,26,7],[20,3,75],[5,33,34],[33,20,50],[29,10,0],[7,5,2],[25,47,48],[28,18,10],[8,50,19],[32,8,46],[18,29,48],[21,47,19],[15,3,31],[39,44,65],[39,27,38],[13,25,76],[17,33,91],[14,2,38],[18,9,10],[17,16,67],[25,23,16],[6,43,95],[34,2,18],[14,26,73],[11,25,39],[2,34,25],[2,38,88],[17,50,85],[30,14,38],[31,9,60],[38,40,8],[21,20,30],[24,15,89],[16,10,21],[31,1,71],[48,38,81],[46,44,34],[11,31,86],[33,34,39],[21,8,29],[28,30,47],[41,8,69],[30,42,78],[4,10,54],[21,29,0],[9,16,50],[38,33,75],[20,42,64],[30,7,67],[3,30,52],[6,12,97],[39,18,28],[22,20,53],[6,37,6],[12,9,64],[15,5,40],[26,8,40],[28,22,70],[45,27,44],[18,28,54],[15,42,13],[1,44,8],[8,27,2],[22,14,39],[40,16,16],[12,25,13],[22,9,15],[32,14,0],[31,35,74],[23,48,92],[2,46,13],[35,5,79],[41,33,95],[1,48,78],[9,24,23],[31,42,21],[40,41,56],[21,34,93],[40,37,14],[2,50,94],[33,4,25],[27,45,89],[23,37,96],[49,23,42],[49,34,97],[30,47,43],[30,24,10],[6,50,5],[34,36,68],[40,13,95],[31,49,12],[26,40,0],[17,45,44],[33,19,60],[44,13,7],[47,31,78],[29,44,20],[20,38,63],[2,16,13],[29,1,14],[47,46,39],[21,43,5],[18,49,7],[5,31,79],[30,31,71],[12,36,54],[42,40,56],[33,47,66],[49,32,70],[29,20,83],[8,1,0],[49,47,41],[4,2,15],[44,21,26],[8,42,47],[10,2,74],[5,10,24],[15,30,2],[13,28,76],[23,9,84],[34,27,47],[24,7,84],[3,19,24],[34,33,44],[21,1,99],[24,19,76],[25,9,45],[1,19,86],[29,18,7],[23,42,92],[15,24,43],[22,21,92],[42,3,35],[37,19,94],[46,24,56],[50,17,22],[44,22,83],[28,24,15],[23,18,93],[6,5,52],[47,10,77],[50,48,64],[37,7,21],[17,15,4],[34,35,56],[46,33,31],[39,3,74],[40,28,94],[28,1,20],[41,21,57],[1,35,57],[6,39,70],[13,45,83],[9,47,64],[12,19,12],[47,9,54],[9,46,41],[38,6,40],[28,16,39],[38,41,40],[5,44,12],[12,17,18],[21,14,46],[9,34,74],[31,23,58],[6,7,64],[37,16,50],[17,39,8],[17,37,81],[31,38,9],[8,32,73],[32,12,86],[16,9,63],[2,49,26],[25,26,47],[48,35,59],[42,43,62],[36,45,46],[25,37,7],[47,2,54],[1,2,59],[36,16,58],[18,47,91],[35,15,5],[38,46,42],[17,42,62],[11,4,7],[7,44,68],[44,49,74],[3,28,13],[35,21,15],[47,24,8],[50,32,23],[16,24,19],[47,25,53],[43,28,81],[46,10,32],[13,26,93],[36,44,34],[5,29,41],[40,34,99],[2,12,1],[1,45,61],[7,42,94],[8,28,26],[7,34,16],[1,37,20],[43,16,59],[4,38,32],[21,19,38],[22,44,29],[25,22,53],[50,39,32],[35,46,54],[33,30,43],[3,33,24],[2,7,20],[10,11,83],[36,50,12],[21,42,89],[20,23,66],[16,50,21],[46,23,22],[20,39,24],[32,33,33],[6,33,52],[26,15,52],[29,42,17],[10,6,30],[5,28,30],[1,30,67],[43,1,52],[14,1,64],[7,33,38],[30,36,86],[9,23,60],[21,15,75],[17,4,53],[38,16,39],[34,7,43],[17,9,31],[12,22,61],[24,41,7],[37,43,10],[18,1,83],[23,35,91],[38,12,80],[49,41,44],[29,32,49],[22,48,43],[30,5,15],[2,25,44],[1,5,56],[20,9,96],[41,31,89],[24,36,96],[12,50,24],[26,50,26],[15,16,49],[10,29,32],[31,39,49],[17,18,3],[45,13,62],[26,38,14],[12,42,31],[44,5,29],[2,30,79],[28,45,29],[44,27,16],[12,49,78],[38,5,55],[9,49,15],[12,7,19],[29,8,69],[22,47,57],[24,37,96],[33,40,90],[5,46,64],[23,21,46],[5,35,80],[37,5,32],[25,41,10],[43,38,11],[50,19,19],[34,24,14],[7,10,55],[35,9,7],[41,25,4],[30,10,57],[10,21,12],[20,44,95],[36,18,94],[13,50,12],[30,27,66],[13,14,55],[31,14,27],[23,41,70],[17,41,65],[24,2,72],[34,46,26],[50,12,52],[44,42,39],[5,50,14],[34,17,87],[26,30,63],[47,41,2],[36,14,64],[5,27,40],[22,32,11],[14,32,49],[14,39,51],[25,16,13],[20,11,67],[42,5,32],[27,48,45],[30,28,80],[40,8,33],[39,29,64],[42,35,24],[26,27,90],[20,18,79],[2,22,22],[20,21,36],[35,42,39],[3,5,12],[38,36,55],[16,19,32],[27,34,61],[43,47,47],[4,34,99],[7,36,59],[26,44,80],[10,4,11],[31,30,26],[23,34,80],[38,25,95],[8,15,37],[3,7,23],[11,14,25],[19,46,41],[17,29,5],[8,29,29],[36,37,85],[36,34,39],[44,7,32],[13,17,20],[15,2,67],[39,15,57],[5,26,66],[23,47,70],[19,40,30],[34,5,52],[14,10,76],[11,19,84],[13,9,17],[2,10,95],[24,38,93],[37,46,77],[19,5,14],[40,5,67],[12,21,97],[3,39,99],[48,4,54],[6,1,87],[7,6,74],[18,22,91],[31,15,21],[12,40,3],[16,2,52],[9,50,57],[30,23,26],[39,42,57],[30,16,32],[23,30,58],[41,2,48],[12,27,31],[8,24,53],[47,27,50],[35,27,29],[3,25,5],[38,7,43],[23,11,61],[15,12,14],[49,22,21],[29,9,19],[17,22,60],[34,22,24],[41,42,84],[28,32,39],[39,21,25],[10,18,78],[19,38,71],[25,30,97],[38,2,30],[5,11,98],[36,27,53],[29,21,8],[44,26,48],[12,37,73],[10,8,84],[50,5,44],[47,3,32],[8,22,50],[39,10,78],[43,33,71],[40,26,66],[6,11,96],[26,35,22],[49,39,54],[45,34,91],[19,6,59],[26,39,55],[34,25,20],[48,37,49],[21,10,28],[17,3,77],[3,48,80],[38,27,53]]
n = 50
k = 44

print(Solution().networkDelayTime(times, n, k))
