# There's a linear solution but it's too tricky to figure out the formula.
#
# This solution just goes through and runs the entire sequence, counting time spent.
#
# The only tricky part is to make sense of the "cycle length" idea:
# - Example: `["A","A","A","B","B","B"]` with `n = 2` `=>` `(A,B,idle),(A,B,idle),A,B`.
# - The cycle length is thus `n+1`, and we fill it with tasks with the longest freq first.
# - Do this for every cycle and keep track of "consumed" tasks.

from heapq import heappush, heappop

class Solution:
    # Time: O(t*logt) where t is len(tasks)
    # Space: O(k) where k is the number of distinct tasks
    def leastInterval(self, tasks: List[str], n: int) -> int:
        h = [-freq for freq in Counter(tasks).values()]
        heapify(h)
        
        total_time = 0
        while h:
            remain = []
            # The "cycles" will have n+1 length. Example:
            # ["A","A","A","B","B","B"] with n = 2 => (A,B,idle),(A,B,idle),A,B
            # Since there must be 2 non-A between As.
            cycle_len = n + 1

            # In the current cycle
            while cycle_len and h:
                # Run tasks sorted by freq
                max_freq = -heappop(h)
                # We're gonna push the task back to the heap with freq-1
                if max_freq > 1:
                    remain.append(max_freq-1)
                
                # Since we ran a task, increment total time and decrement cycle length
                total_time += 1
                cycle_len -= 1
            
            # Push the decremented tasks back to the heap
            for count in remain:
                heappush(h, -count)
            
            # If heap is empty, this was the last cycle.
            # Thus, we can end here: idle space in this cycle doesn't count.
            if not h:
                break
            
            # The current cycle has some idle space left.
            # Since the heap is not empty, we have another cycle.
            # Therefore, add idle space to total time.
            total_time += cycle_len

        return total_time
