# wip!
# Straightforward reasoning from exercise:
# - Remember to sort tasks by enqueue time.
# - Need to keep pulling shortest (processing_time, label), so use a min heap!
# - The only trick is that heap may empty if next enqueue task start time is later than last
#   task's final time, so in that case just advance cur_time and push tasks again (not just 1!).
import heapq

# Time: O(n*logn) both for sorting and for pushing and popping
# Space: O(n) for the extra space of keeping sorted tasks and the heap
class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        # We need to sort tasks by enqueue time, so first save the task label.
        for i in range(len(tasks)):
            tasks[i].append(i)
        
        # Sort tasks by enqueue time. Now: [enqueue_time, processing_time, label]
        tasks = sorted(tasks)
        
        
        cur_time = tasks[0][0] # We're guaranteed there's at least one task
        order = []             # The resulting order
        options = []           # We'll keep a heap of options
        i = 0                  # To index over the sorted tasks array
        
        while len(order) < len(tasks):
            # Push available options into heap
            while i < len(tasks) and cur_time >= tasks[i][0]:
                heapq.heappush(options, ((tasks[i][1], tasks[i][2]), tasks[i][0]))
                i += 1
            
            # It's possible that the heap is empty:
            if not options:
                # This happens when the next task's enqueue time is later than when
                # the last task finished processing. In this case, advance cur_time
                # to the next enqueue_time, and push available options again.
                cur_time = tasks[i][0]
                while i < len(tasks) and cur_time >= tasks[i][0]:
                    heapq.heappush(options, ((tasks[i][1], tasks[i][2]), tasks[i][0]))
                    i += 1
            
            (processing_time, label), _ = heapq.heappop(options)
            
            order.append(label)
            cur_time += processing_time
                        
        return order
