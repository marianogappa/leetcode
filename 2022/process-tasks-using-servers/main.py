# wip!
# Available servers seem to be sorted by weight and index, with reclaimed servers
# being added online, not to bottom nor top. This suggests available servers should be a min-heap.
#
# When picking up a task, servers are popped from the heap and become busy until a known time which
# can be earlier, later or same to other busy servers. To reclaim busy servers, we need to "peek" the
# earliest reclaim server time efficiently. That sounds like another min-heap.

# Time: O(t*log(s))
# Space: O(s)
class Solution:
    def assignTasks(self, servers: List[int], tasks: List[int]) -> List[int]:
        busy_servers = []

        free_servers = []
        for idx, weight in enumerate(servers):
            heapq.heappush(free_servers, ((weight, idx), None))
        
        allocated_server_idxs = []
        for time, task_len in enumerate(tasks):
            # Reclaim busy servers whose business has elapsed
            while len(busy_servers) and busy_servers[0][0] <= time:
                _, (weight, idx) = heapq.heappop(busy_servers)
                heapq.heappush(free_servers, ((weight, idx), None))
            
            # Hopefully we can just pop the next free server. If there aren't any, then we need to wait until
            # the earliest reclaimable busy server's time.
            busy_until = time
            if not len(free_servers):
                busy_until, (weight, idx) = heapq.heappop(busy_servers)
                heapq.heappush(free_servers, ((weight, idx), None))

            # Pop the next free server, note down their index and make it busy with the current task
            (weight, idx), _ = heapq.heappop(free_servers)
            allocated_server_idxs.append(idx)
            heapq.heappush(busy_servers, (busy_until + task_len, (weight, idx)))
        
        return allocated_server_idxs
