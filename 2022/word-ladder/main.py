# If the wordlist was a graph including beginWord, finding the shortest transformation sequence would be running Dijkstra!
#
# Creating the adjacency list is just calculating, for all tuples of words, which have 1 different letter. This step is `O(n^2)*w`, which dominates the complexity.
# The space will also be quadratic, but only `O(n^2)`, since the adjacency matrix could connect everything.

class Solution:
    def ladderLength(self, beginWord: str, endWord: str, wordList: List[str]) -> int:
        # Find the index of endWord in wordList
        end_word_idxs = [i for i in range(len(wordList)) if wordList[i] == endWord]
        if not end_word_idxs:
            return 0
        end_i = end_word_idxs[0]

        # This ensured beginWord is in wordList, and gets the index of it (to run Dijkstra)
        begin_word_idxs = [i for i in range(len(wordList)) if wordList[i] == beginWord]
        begin_i = begin_word_idxs[0] if begin_word_idxs else len(wordList)
        if begin_i == len(wordList):
            wordList.append(beginWord)

        # Run Dijkstra
        distance = dijkstra(create_adjacency(wordList), len(wordList), begin_i)[end_i]
        
        # Remember that Dijkstra returns min count of edges, so we must +1.
        return distance + 1 if distance != float('inf') else 0

def are_adjacent(w1: str, w2: str) -> bool:
    differ_count = 0
    for i in range(len(w1)):
        if w1[i] != w2[i]:
            differ_count += 1
        if differ_count > 1:
            return False
    return True

def create_adjacency(words: list[str]) -> dict[int, list[int, int]]:
    adj: dict[int, list[int, int]] = defaultdict(list)
    for i in range(len(words)):
        for j in range(i+1, len(words)):
            if are_adjacent(words[i], words[j]):
                adj[i].append([j, 1])
                adj[j].append([i, 1])
    return adj

# Time: O((V+E)*logV)
# Space: O(V)
def dijkstra(
    adjacency: dict[int, list[list[int]]],
    num_vertices: int, # != len(adjacency) if there are orphan vertices!!
    start_vertex: int,
) -> list[int]:
    # We start with default assumption that distance from start_vertex to every vertex
    # is infinite, except for start_vertex.
    distances = [float('inf')] * (num_vertices+1) # +1 optional; depends on zero-indexed
    distances[start_vertex] = 0
    
    # We only want to visit each vertex once.
    visited: set[int] = set()

    # Use a heap to iteratively pop the shortest distance vertex, 
    # until all vertices are visited
    h = []
    heapq.heappush(h, (0, start_vertex))

    # O(v): While there are still unvisited vertices...
    while len(h) > 0:
        # O(log v): Pop the next shortest distance vertex
        cur_dist, vertex = heapq.heappop(h)
        
        # O(1): Get the vertex's edges and mark the vertex as visited
        edges = adjacency[vertex]
        visited.add(vertex)
        
        # O(e): For every edge...
        for edge in edges:
            [dest_vertex, distance] = edge

            # If the vertex is visited, or we already found a shortest distance, 
            # ignore the edge
            if dest_vertex in visited or distances[dest_vertex] <= cur_dist + distance:
                continue

            # O(log v): Store the shortest distance and add it to the heap, 
            # so that we visit it later
            distances[dest_vertex] = cur_dist + distance
            heapq.heappush(h, (distances[dest_vertex], dest_vertex))

    return distances

# Note that BFS is more efficient, because we could find the solution without traversing all levels.
#
# But it's less intuitive to me, because we must keep a visited `set`, but isn't the set different depending on the direction we took?
#
# According to this solution it's not, but I don't think I can reliably remember why: https://leetcode.com/problems/word-ladder/solutions/1764371/a-very-highly-detailed-explanation/
