# wip!
# The intuitions are pretty clear: 
#  - For insert+remove in O(1), a set is all that's needed.
#  - With a set, getRandom cannot be done in O(1).
#  - To do getRandom in O(1) an array is needed, so let's keep both.
#  - The only issue is that remove in array is O(n). How to circumvent this issue?
# => Remove is O(n) only if keeping order is required. Since it's not, just patch the whole with
#    last item in array!


from random import choice


class RandomizedSet:

    def __init__(self):
        self.arr: list[int] = []
        self.dict: dict[int, int] = {}

    def insert(self, val: int) -> bool:
        # Val exists in set => no-op
        if val in self.dict:
            return False

        # Val doesn't exist: append to array and add to dict, with array key as value
        self.arr.append(val)
        self.dict[val] = len(self.arr)-1

        return True

    def remove(self, val: int) -> bool:
        # Val doesn't exist => no-op
        if val not in self.dict:
            return False

        # Remove from set, but get the idx in array first
        idx = self.dict[val]
        self.dict.pop(val)

        # Move the tail element to fill the gap left by removed element
        self.arr[idx] = self.arr[-1]
        self.arr.pop()

        # If element was already on the tail, we're done.
        # Otherwise, we need to update the dict's value for the moved element        
        if idx < len(self.arr):
            self.dict[self.arr[idx]] = idx

        return True

    def getRandom(self) -> int:
        return choice(self.arr)


# Your RandomizedSet object will be instantiated and called as such:
obj = RandomizedSet()
print(obj.insert(1))
print(obj.remove(2))
print(obj.insert(2))
print(obj.getRandom())
print(obj.remove(1))
print(obj.insert(2))
print(obj.getRandom())

