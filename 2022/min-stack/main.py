# The confusing bit is that it seems it's asking to sort in `O(1)`.
#
# But it's not asking that. Every entry can know what the min is at the time it was pushed, by asking the previous top what their min was and comparing it with the new value. This won't change retroactively, so it is indeed `O(1)`.

class MinStack:

    def __init__(self):
        # The first elem is the value, and the second is the "current min".
        self.stack: list[tuple[int, int]] = []

    def push(self, val: int) -> None:
        self.stack.append((val, min(val, self.stack[-1][1] if self.stack else val)))

    def pop(self) -> None:
        self.stack.pop()

    def top(self) -> int:
        return self.stack[-1][0]

    def getMin(self) -> int:
        return self.stack[-1][1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
