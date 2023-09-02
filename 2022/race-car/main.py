# wip!
# BFS looks natural for this, because a DFS would not finish in an infinite tree. Still, some 
# optimizations are necessary, since vanilla BFS times out.
#
# The typical optimisation of pruning branches seen before is good but solution still times out.
# To be clear, that means a "visited" set using a (position, speed) tuple.
#
# At this point, reducing branching factor makes sense.
#
# Reversing makes sense only in two cases, and shouldn't be used otherwise:
# - If position is moving away from target
# - If position leads to target, but will exceed it on the next step
#
# Accelerating is harder to reason about. It would seem that one wouldn't want to accelerate if
# moving away from target, but look at this example (target => 5):
#
#                     HERE
# 0 =>  1 => 3 => 3 => 2 => 2 => 3 => 5
#     ['A', 'A', 'R', 'A', 'R', 'A', 'A']
#
# In that case, accelerating away once yields a shorter path.
#
# A compromise optimisation is to not allow accelerating if abs(position-target) > 2*target.
# This is because the farthest that an acceleration can overshoot target is 2*target-1.
import collections

class Solution:
    """
    Time complexity analysis:
    
        An upper bound can be: iterations can cover positions in the range (-2*target, 2*target), 
        and 2*log(target) different speeds, since they grow quadratically.

        All combinations yield 4t*2logt => O(t*logt)
    
    Space: O(t*logt) is identical because visited iterations are saved
    """
    def racecar(self, target: int) -> int:
        queue = collections.deque()
        
        # Position & speed
        queue.append((0, 1))
        visited = {(0, 1)}
        len_of_instructions = 0
        
        while True:
            for _ in range(len(queue)):
                position, speed = queue.popleft()
                
                if position == target:
                    return len_of_instructions
                
                visited.add((position, speed))
                
                accelerate = (position + speed, speed * 2)
                if accelerate not in visited and abs(position-target) < 2*target:
                    queue.append(accelerate)

                reverse = (position, -1 if speed > 0 else 1)
                if reverse not in visited and ((speed > 0 and accelerate[0] > target) or (speed < 0 and accelerate[0] < target)):
                    queue.append(reverse)
            
            len_of_instructions += 1
            
            

print(Solution().racecar(4), "== 5")
print(Solution().racecar(5), "== 7")
