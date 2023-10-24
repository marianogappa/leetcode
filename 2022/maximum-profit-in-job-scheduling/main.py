# If there was no overlap, we could just sum all profits.
# 
# For overlapping job time ranges, there's no easy clever solutions than to try all options.
#
# To efficiently check for overlaps, it's necessary to sort by start times. Zip a tuple of all lists to sort easily.
#
# When a job is picked, the next job must start at at least the end time of picked job. Finding this index is `O(n)` with a for-loop, so use binary search to optimise.
#
# This algorithm is `2^n` time unless we use memoization. It's safe to memoize on `dict[start,profit]`.

class Solution:
    # Time: O(n*logn) n == len(jobs) since sorting is most expensive (dfs+memo is linear)
    # Space: O(n) for DP memo + recursion stack
    def jobScheduling(self, startTime: List[int], endTime: List[int], profit: List[int]) -> int:
        return dfs(sorted(zip(startTime, endTime, profit)), 0, {})

def dfs(jobs: list[tuple[int, int, int]], i: int, memo: dict[int, int]) -> int:
    if i >= len(jobs):
        return 0
    
    if memo.get(i):
        return memo[i]
    
    # If we're at the last job, this should should always be done
    if i == len(jobs)-1:
        return jobs[i][2]
    
    # If there's no overlap, this job should always be done
    if not is_overlap(jobs[i], jobs[i+1]):
        return jobs[i][2] + dfs(jobs, i+1, memo)

    # At this point there's an overlap and no clever way to know which option is better
    # so try both!
    # - If not do it, proceed to next
    # - If do it, proceed to immediate next after end time (use binary search to find next)
    memo[i] = max(
        dfs(jobs, i+1, memo), 
        jobs[i][2] + dfs(jobs, binsearch(jobs, jobs[i][1], i + 1), memo)
    )
    return memo[i]
    
def binsearch(jobs: list[tuple[int, int, int]], target: int, start: int) -> int:
    end = len(jobs)-1
    candidate = end + 1
    while start <= end:
        mid = (start + end) // 2
        if jobs[mid][0] < target:
            start = mid + 1
        else:
            candidate = mid
            end = mid - 1
    
    return candidate

# Note that i1 & i2 are sorted by start time, making overlap check trivial
def is_overlap(i1, i2) -> bool:
    return i1[1] > i2[0]
