# wip!
# In principle it would be trivial to just keep track of max, min and latest, but the issue is that updates can
# override invalid values, which requires keeping track of the "previous max, min & latest" recursively. For this reason,
# one must explore the most efficient way of keeping a live sorted list of numbers: heaps! But heaps have one flaw:
# how do you delete? Simplest way to deal with this problem is to not delete them, but to know if they are outdated or
# not by keeping a hashmap of latest value by timestamp.
#
# The other solution to this is to use sorted maps.
import heapq

# Space: O(3n) = O(n)
#
# Update:
#     Time: O(log n) where n is the amount of times update was called
#
# Current:
#     Time: O(1)
#
# Maximum:
#     Time: O(1 amortized) because in a very unfortunate case that heap is filled with outdated values it could be linear
#
# Minimum:
#     Time: O(1 amortized) because in a very unfortunate case that heap is filled with outdated values it could be linear
class StockPrice:

    def __init__(self):
        self.min_heap = []
        self.max_heap = []
        self.latest_timestamp = 0
        self.timestamp_to_price = {}

    def update(self, timestamp: int, price: int) -> None:
        heapq.heappush(self.min_heap, (price, timestamp))
        heapq.heappush(self.max_heap, (-price, timestamp))
        self.latest_timestamp = max(self.latest_timestamp, timestamp)
        self.timestamp_to_price[timestamp] = price

    def current(self) -> int:
        return self.timestamp_to_price[self.latest_timestamp]

    def maximum(self) -> int:
        while True:
            price, timestamp = self.max_heap[0]
            price = -price
            if self.timestamp_to_price[timestamp] == price:
                return price
            heapq.heappop(self.max_heap)

    def minimum(self) -> int:
        while True:
            price, timestamp = self.min_heap[0]
            if self.timestamp_to_price[timestamp] == price:
                return price
            heapq.heappop(self.min_heap)

        # Your StockPrice object will be instantiated and called as such:
        # obj = StockPrice()
        # obj.update(timestamp,price)
        # param_2 = obj.current()
        # param_3 = obj.maximum()
        # param_4 = obj.minimum()
