# Time: O(nlogn) where n is len(weights)
# Space: O(log n) for the binary search stack. Iterative binary search would improve to O(1).
#
# Brute force solution: if capacity == sum(weights) => days == 1. So do a for-loop that decrements
# capacity, until days is higher than the specified value. The candidate capacity right before the
# first invalid one is the right one.
#
# How can we improve on the brute force solution? Instead of looping over all candidates, binary
# search over them! That's it.
class Solution:
    def shipWithinDays(self, weights: List[int], days: int) -> int:
        return binary_search(max(weights), sum(weights), weights, days)
    
def binary_search(min_capacity, max_capacity: int, weights: List[int], ideal_days) -> int:
    # If the range of search is 1 capacity, then we found the ideal solution
    if min_capacity == max_capacity:
        return min_capacity
    
    candidate_capacity = (max_capacity + min_capacity) // 2
    candidate_days = days_it_takes(candidate_capacity, weights)
    
    # This candidate capacity doesn't work: it takes more days than ideal
    if candidate_days > ideal_days: 
        return binary_search(candidate_capacity + 1, max_capacity, weights, ideal_days)
    # This candidate capacity may work, but there might be a valid smaller one
    elif candidate_days <= ideal_days:
        # Note that candidate capacity is INCLUDED in the next binary search, cause it's possible!
        return binary_search(min_capacity, candidate_capacity, weights, ideal_days)

# How many days it takes to ship all packages with this capacity
def days_it_takes(candidate_capacity, weights) -> int:
    # Minimum number of days is 1 (assuming there's at least 1 package)
    days_it_took = 1
    
    # Try to fit all packages...
    remaining_capacity = candidate_capacity
    for weight in weights:
        # Fill up the ship until no capacity left
        remaining_capacity -= weight

        # Since we exceeded capacity, this package will go on the next day's ship
        if remaining_capacity < 0:
            days_it_took += 1
            remaining_capacity = candidate_capacity - weight
        
    return days_it_took
