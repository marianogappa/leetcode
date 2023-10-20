# The only trick is to store all values of a key as a list of tuples: `(timestamp, value)`, and then use binary search to find the right one.

class TimeMap:

    def __init__(self):
        self.data: dict[str, list[tuple(int, str)]] = defaultdict(list)
        
    def set(self, key: str, value: str, timestamp: int) -> None:
        self.data[key].append((timestamp, value))

    def get(self, key: str, timestamp: int) -> str:
        values = self.data[key]
        return binary_search(values, timestamp, 0, len(values)-1) if values else ""

def binary_search(values: list[tuple[int, str]], timestamp: int, left: int, right: int) -> str:
    if left == right:
        return values[left][1] if values[left][0] <= timestamp else ""

    # Because the `<=` case preserves `mid`, this is also an edge case:
    if left+1 == right:
        result = ""
        # Greedily take higher timestamp that satisfies
        if values[left][0] <= timestamp:
            result = values[left][1]
        if values[right][0] <= timestamp:
            result = values[right][1]
        return result

    mid = int((left + right) / 2)

    if values[mid][0] > timestamp:
        return binary_search(values, timestamp, left, mid-1)
    if values[mid][0] <= timestamp:
        return binary_search(values, timestamp, mid, right)



# Your TimeMap object will be instantiated and called as such:
# obj = TimeMap()
# obj.set(key,value,timestamp)
# param_2 = obj.get(key,timestamp)
