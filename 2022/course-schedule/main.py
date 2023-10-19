# Kahn's Topological Sort will yield the order of courses satisfying prerequisites.
#
# If the resulting order has less items than `numCourses`, then there is a cycle!
#
# So just run Kahn and return `not(len(order) < numCourses)`.

class Solution:
    # Time: O(V*E) where V is numCourses and E is len(prerequisites)
    # Space: O(V + E)
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        return not kahn_has_cycle(numCourses, prerequisites)

def kahn_has_cycle(numCourses: int, prerequisites: List[List[int]]) -> bool:
    prerequisites_of: dict[int, list[int]] = defaultdict(list)
    for prereq in prerequisites:
        prerequisites_of[prereq[0]].append(prereq[1])

    indegree = {vertex: 0 for vertex in range(numCourses)}
    
    for vertex in prerequisites_of.keys():
        for prereq in prerequisites_of[vertex]:
            indegree[prereq] += 1
    
    queue = deque([vertex for vertex, count in indegree.items() if count == 0])

    order = []
    while queue:
        vertex = queue.popleft()
        order.append(vertex)

        for prereq in prerequisites_of[vertex]:
            indegree[prereq] -= 1
            if indegree[prereq] == 0:
                queue.append(prereq)
    
    return len(order) < len(indegree)
        

