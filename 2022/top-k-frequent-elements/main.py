# Straightforward exercise with no tricks. "Top k" in the title is screaming Heaps, but if
# not clear, it says in the end "solve sorting faster than n*logn!" which also means Heaps.
#
# The heap needs a value to max-heap over, and that is the frequency of each number, so
# we need a hashmap to calculate the frequency. Note that if it was letters in the alphabet
# instead of numbers, then an array[26] would be more efficient. In any other case, use a
# hashmap!
class Solution:
    # Time: O(k*logn) due to k heap pops
    # Space: O(n) to store all nums, first in the hashmap, then in the tuple array, then in top_k
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # Time: O(n)
        # Make a frequency hashmap of nums
        freqs: dict[int, int] = dict()
        for num in nums:
            if not freqs.get(num):
                freqs[num] = 0
            freqs[num] += 1

        # Time: O(n)
        # Convert the hashmap into an array of tuples, freq first and num second.
        # The freq needs to be negated so that, when made a heap, it's a max-heap.
        tuples: list[tuple[int, int]] = []
        for num, freq in freqs.items():
            tuples.append((-freq, num))

        # Time: O(n)
        heapq.heapify(tuples)

        # Time: O(k*logn)
        # Pop the k max-elements, each one taking logn
        top_k: list[int] = []
        for _ in range(k):
            top_k.append(heapq.heappop(tuples)[1])

        return top_k
