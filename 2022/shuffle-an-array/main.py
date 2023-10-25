from random import randint

class Solution:
    def __init__(self, nums: List[int]):
        self.original = nums

    def reset(self) -> List[int]:
        return self.original.copy()

    # This is called Fisher-Yates apparently
    def shuffle(self) -> List[int]:
        arr = self.original.copy()
        for i, el in enumerate(arr):
            j = randint(0, len(arr)-1)
            arr[i], arr[j] = arr[j], arr[i]
        return arr

        


# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.reset()
# param_2 = obj.shuffle()
